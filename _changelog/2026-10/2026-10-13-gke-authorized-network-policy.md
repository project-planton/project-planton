# Explicit GKE Authorized-Network Policies

GKE control-plane IP filtering now has an explicit catalog choice. A configured
empty `masterAuthorizedNetworks: {}` restricts access without authorizing external
CIDR ranges; omitting the policy or choosing `unrestrictedNetworks: {}` configures
no authorized-networks restriction. Authentication and endpoint availability remain
independent.

The two policies are alternatives of `authorized_networks_policy`. The existing
`master_authorized_networks` field keeps number 28 and its JSON spelling, so
existing manifests retain their meaning. The unrestricted alternative is an empty
marker at field 77. The schema now tells descriptor-driven editors that an empty
restricted policy is an authored choice rather than form residue.

Pulumi reads the selected policy through the generated accessor. OpenTofu's
generator-owned variables expose both alternatives, while the module retains a
present empty restricted block. Tests cover JSON and binary round trips, conflicting
alternatives, actual Pulumi registration inputs and mocked OpenTofu plans for all
policy states. A preset demonstrates an empty authorized CIDR list with a separate
IAM-authenticated DNS access path.

Consumers adopting this schema must choose the policy explicitly: an untouched
form must not seed an empty restricted block. Changing from restricted to
unrestricted is an intentional policy change; editing an unrelated setting must
preserve the chosen policy.
