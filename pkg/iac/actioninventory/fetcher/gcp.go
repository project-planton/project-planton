package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/plantonhq/planton/pkg/iac/actioninventory"
	"github.com/plantonhq/planton/pkg/iac/permissions"
)

// queryTestablePermissionsURL is IAM's inventory of every permission name
// GCP defines -- the same list `gcloud iam list-testable-permissions`
// reads. The query anchors on a resource, but permission NAMES are
// global: permissions of services not enabled on the anchor project are
// still listed (flagged apiDisabled), so any project answers the
// existence question.
const queryTestablePermissionsURL = "https://iam.googleapis.com/v1/permissions:queryTestablePermissions"

// The org and billing-account anchor environment variables. IAM's
// inventory is SCOPE-TYPED: a permission is listed only under the
// resource types it can be tested on, so permissions that authorize
// exclusively on organizations/folders (resourcemanager.projects.create)
// or billing accounts (billing.resourceAssociations.create) are
// structurally invisible to a project-anchored query -- the gate would
// reject real, live-verified permissions. The fetcher therefore queries
// all three scopes and unions the results. Like the project anchor, WHICH
// org or billing account does not matter (names are global within a
// scope); the anchors exist only to type the queries, and they are
// explicit configuration because a principal can often see several. The
// union keeps which scopes listed each permission: a permission no project
// lists is one a connection's project role can never hold, and the
// snapshot records it as an organization or billing-account permission so
// every manifest entry's reach is checked against Google's own answer.
const (
	gcpOrgEnvVar            = "PLANTON_GCP_ORG"
	gcpBillingAccountEnvVar = "PLANTON_GCP_BILLING_ACCOUNT"
)

// refreshGcp rewrites the committed GCP snapshot
// (pkg/iac/actioninventory/gcp.yaml) from IAM's testable-permissions
// inventory, scoped to exactly the services (the first dotted segment,
// e.g. "container" in "container.clusters.create") the committed
// permissions manifests reference. The credential contract mirrors the
// GCP price fetcher's ADC arm: the operator's own application-default
// credentials mint the bearer token (network in make, never in CI). The
// queries anchor on three scopes -- the operator's active gcloud project,
// the organization named by PLANTON_GCP_ORG, and the billing account
// named by PLANTON_GCP_BILLING_ACCOUNT -- and union the results, because
// the inventory is scope-typed (see the anchor constants above). It also
// fetches the predefined roles manifests name in a custom role's place,
// keeping the permissions of theirs that manifests reference.
func refreshGcp(repoRoot string) error {
	services, referenced, predefinedRoles, err := referencedGcp(repoRoot)
	if err != nil {
		return err
	}
	if len(services) == 0 {
		fmt.Println("no GCP permissions in any permissions manifest -- gcp snapshot not written")
		return nil
	}

	token, err := gcpAccessToken()
	if err != nil {
		return err
	}
	project, err := gcloudProject()
	if err != nil {
		return err
	}
	org := strings.TrimSpace(os.Getenv(gcpOrgEnvVar))
	if org == "" {
		return fmt.Errorf("no organization anchor -- set %s=<organization-id> (any organization the credential can query; it only anchors the org-scoped inventory query, without which org-scoped permissions like resourcemanager.projects.create are invisible)", gcpOrgEnvVar)
	}
	billingAccount := strings.TrimSpace(os.Getenv(gcpBillingAccountEnvVar))
	if billingAccount == "" {
		return fmt.Errorf("no billing-account anchor -- set %s=<billing-account-id> (any billing account the credential can query; it only anchors the billing-scoped inventory query, without which billing-scoped permissions like billing.resourceAssociations.create are invisible)", gcpBillingAccountEnvVar)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	anchors := []struct {
		scope            string
		fullResourceName string
	}{
		{gcpScopeProject, "//cloudresourcemanager.googleapis.com/projects/" + project},
		{gcpScopeOrganization, "//cloudresourcemanager.googleapis.com/organizations/" + org},
		{gcpScopeBillingAccount, "//cloudbilling.googleapis.com/billingAccounts/" + billingAccount},
	}
	listed := map[string]*gcpListedPermission{}
	for _, anchor := range anchors {
		if err := mergeTestablePermissions(client, token, anchor.scope, anchor.fullResourceName, listed); err != nil {
			return err
		}
	}

	today := time.Now().UTC().Format("2006-01-02")
	inv := &actioninventory.Inventory{Provider: "gcp"}
	byService := map[string]*actioninventory.Service{}
	for _, service := range services {
		byService[service] = &actioninventory.Service{
			Prefix:      service,
			SourceURL:   queryTestablePermissionsURL,
			RetrievedOn: today,
		}
	}
	for permission, entry := range listed {
		service, name, _ := strings.Cut(permission, ".")
		svc := byService[service]
		if svc == nil {
			continue
		}
		svc.Actions = append(svc.Actions, name)
		// A permission's reach is the narrowest scope that lists it: a
		// project when one does, else a billing account, else the
		// organization.
		switch {
		case entry.scopes[gcpScopeProject]:
		case entry.scopes[gcpScopeBillingAccount]:
			svc.BillingAccountActions = append(svc.BillingAccountActions, name)
		default:
			svc.OrganizationActions = append(svc.OrganizationActions, name)
		}
		if entry.customRolesUnsupported {
			svc.CustomRoleUnsupportedActions = append(svc.CustomRoleUnsupportedActions, name)
		}
	}
	for _, service := range services {
		svc := byService[service]
		if len(svc.Actions) == 0 {
			return fmt.Errorf("service %q (used by a permissions manifest) defines no permissions in GCP's IAM inventory -- the service segment itself is wrong", service)
		}
		sort.Strings(svc.Actions)
		sort.Strings(svc.OrganizationActions)
		sort.Strings(svc.BillingAccountActions)
		sort.Strings(svc.CustomRoleUnsupportedActions)
		inv.Services = append(inv.Services, *svc)
	}
	for _, role := range predefinedRoles {
		included, err := predefinedRolePermissions(client, token, role)
		if err != nil {
			return err
		}
		var kept []string
		for _, permission := range included {
			if referenced[permission] {
				kept = append(kept, permission)
			}
		}
		if len(kept) == 0 {
			return fmt.Errorf("predefined role %q holds none of the permissions the manifests reference -- the role named in place of a custom role is wrong", role)
		}
		sort.Strings(kept)
		inv.PredefinedRoles = append(inv.PredefinedRoles, actioninventory.PredefinedRole{Name: role, Permissions: kept})
	}

	out := filepath.Join(repoRoot, "pkg", "iac", "actioninventory", actioninventory.GcpFileName)
	if err := os.WriteFile(out, []byte(actioninventory.Render(inv)), 0o644); err != nil {
		return err
	}
	fmt.Printf("wrote %d service permission list(s) to %s\n", len(inv.Services), out)
	return nil
}

// referencedGcp collects, from every committed permissions manifest, the
// distinct GCP service segments (sorted), the full set of permissions they
// name, and the predefined roles they name in a custom role's place
// (sorted).
func referencedGcp(repoRoot string) ([]string, map[string]bool, []string, error) {
	discovered, err := permissions.Discover(repoRoot)
	if err != nil {
		return nil, nil, nil, err
	}
	services := map[string]bool{}
	referenced := map[string]bool{}
	roles := map[string]bool{}
	for provider, kinds := range discovered {
		for _, kind := range kinds {
			manifest, err := permissions.Load(repoRoot, provider, kind)
			if err != nil {
				return nil, nil, nil, err
			}
			for _, group := range manifest.GetSpec().GetGcp().GetGroups() {
				if role := group.GetPredefinedRole(); role != "" {
					roles[role] = true
				}
				for _, permission := range group.GetPermissions() {
					service, _, found := strings.Cut(permission, ".")
					if !found {
						return nil, nil, nil, fmt.Errorf("%s/%s: gcp permission %q has no service segment", provider, kind, permission)
					}
					services[service] = true
					referenced[permission] = true
				}
			}
		}
	}
	return sortedKeys(services), referenced, sortedKeys(roles), nil
}

func sortedKeys(set map[string]bool) []string {
	keys := make([]string, 0, len(set))
	for key := range set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// gcpAccessToken mints a bearer token from the operator's own
// application-default credentials -- the fetcher's credential contract,
// shared with the GCP price fetcher. Missing credentials are a hard error
// naming the way in.
func gcpAccessToken() (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("gcloud", "auth", "application-default", "print-access-token")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("the IAM testable-permissions API needs a credential: run `gcloud auth application-default login` so the fetcher can mint an access token: %v: %s",
			err, strings.TrimSpace(stderr.String()))
	}
	token := strings.TrimSpace(stdout.String())
	if token == "" {
		return "", fmt.Errorf("`gcloud auth application-default print-access-token` returned an empty token")
	}
	return token, nil
}

// gcloudProject reads the operator's active gcloud project -- the query's
// resource anchor (any project answers; see queryTestablePermissionsURL).
func gcloudProject() (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("gcloud", "config", "get-value", "project")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("reading the active gcloud project (`gcloud config get-value project`): %v: %s",
			err, strings.TrimSpace(stderr.String()))
	}
	project := strings.TrimSpace(stdout.String())
	if project == "" || project == "(unset)" {
		return "", fmt.Errorf("no active gcloud project -- run `gcloud config set project <project-id>` (any project; it only anchors the query)")
	}
	return project, nil
}

// The scopes the inventory is queried at, as the snapshot names reaches.
const (
	gcpScopeProject        = "project"
	gcpScopeOrganization   = "organization"
	gcpScopeBillingAccount = "billing_account"
)

// gcpListedPermission is what the three scoped queries said about one
// permission: which scopes listed it, and whether Google publishes it as
// unsupported in custom roles.
type gcpListedPermission struct {
	scopes                 map[string]bool
	customRolesUnsupported bool
}

// mergeTestablePermissions pages through IAM's permission inventory for
// one scope anchor and records, per permission, that this scope lists it
// and whether custom roles can carry it. The three anchors' answers merge
// order-independently: a permission testable on several scopes is one
// entry holding every scope that listed it.
func mergeTestablePermissions(client *http.Client, token, scope, fullResourceName string, listed map[string]*gcpListedPermission) error {
	pageToken := ""
	for {
		body, err := json.Marshal(map[string]any{
			"fullResourceName": fullResourceName,
			"pageSize":         1000,
			"pageToken":        pageToken,
		})
		if err != nil {
			return err
		}
		request, err := http.NewRequest(http.MethodPost, queryTestablePermissionsURL, bytes.NewReader(body))
		if err != nil {
			return err
		}
		request.Header.Set("Authorization", "Bearer "+token)
		request.Header.Set("Content-Type", "application/json")
		response, err := client.Do(request)
		if err != nil {
			return err
		}
		var doc struct {
			Permissions []struct {
				Name string `json:"name"`
				// Absent means SUPPORTED: the API omits the enum's zero
				// value.
				CustomRolesSupportLevel string `json:"customRolesSupportLevel"`
			} `json:"permissions"`
			NextPageToken string `json:"nextPageToken"`
		}
		decodeErr := json.NewDecoder(response.Body).Decode(&doc)
		response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return fmt.Errorf("POST %s for %s: %s", queryTestablePermissionsURL, fullResourceName, response.Status)
		}
		if decodeErr != nil {
			return fmt.Errorf("POST %s for %s: decoding: %w", queryTestablePermissionsURL, fullResourceName, decodeErr)
		}
		for _, permission := range doc.Permissions {
			if _, name, found := strings.Cut(permission.Name, "."); !found || name == "" {
				continue
			}
			entry := listed[permission.Name]
			if entry == nil {
				entry = &gcpListedPermission{scopes: map[string]bool{}}
				listed[permission.Name] = entry
			}
			entry.scopes[scope] = true
			if permission.CustomRolesSupportLevel == "NOT_SUPPORTED" {
				entry.customRolesUnsupported = true
			}
		}
		if doc.NextPageToken == "" {
			return nil
		}
		pageToken = doc.NextPageToken
	}
}

// predefinedRoleURL is IAM's read of one role; a predefined role's name
// ("roles/billing.user") is its resource path.
const predefinedRoleURL = "https://iam.googleapis.com/v1/"

// predefinedRolePermissions reads the permissions a predefined role holds.
func predefinedRolePermissions(client *http.Client, token, role string) ([]string, error) {
	if !strings.HasPrefix(role, "roles/") {
		return nil, fmt.Errorf("predefined role %q is not a predefined role name (roles/<name>)", role)
	}
	request, err := http.NewRequest(http.MethodGet, predefinedRoleURL+role, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s%s: %s -- the role a manifest names does not exist", predefinedRoleURL, role, response.Status)
	}
	var doc struct {
		IncludedPermissions []string `json:"includedPermissions"`
	}
	if err := json.NewDecoder(response.Body).Decode(&doc); err != nil {
		return nil, fmt.Errorf("GET %s%s: decoding: %w", predefinedRoleURL, role, err)
	}
	return doc.IncludedPermissions, nil
}
