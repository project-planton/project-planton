package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/vars"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	pulumiyaml "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// gatewayApis installs kubernetes gateway-api crds
func gatewayApis(ctx *pulumi.Context,
	kubernetesProvider *pulumikubernetes.Provider) error {
	//create gateway-api crd resources
	for _, crdFile := range vars.GatewayApis.CrdFiles {
		_, err := pulumiyaml.NewConfigFile(ctx,
			fmt.Sprintf("gateway-api-crd-%s", crdFile),
			&pulumiyaml.ConfigFileArgs{
				File: fmt.Sprintf("%s/%s", vars.GatewayApis.CrdDownloadBaseUrl, crdFile),
			}, pulumi.Provider(kubernetesProvider),
		)
		if err != nil {
			return errors.Wrapf(err, "failed to add %s gateway-api crd manifest", crdFile)
		}
	}
	return nil
}
