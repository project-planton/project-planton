package pulumiakskubernetesprovider

// AzureExecPluginKubeConfigTemplate requires the following inputs for rendering a kube-config that works
// 1. cluster endpoint ip
// 2. cluster cert-authority data
// 3. azure client id
// 4. azure client secret
// 5. azure tenant id
const AzureExecPluginKubeConfigTemplate = `
apiVersion: v1
kind: Config
current-context: kube-context
contexts:
  - name: kube-context
    context:
      cluster: kube-cluster
      user: kube-user
clusters:
  - name: kube-cluster
    cluster:
      server: %s
      certificate-authority-data: %s
users:
  - name: kube-user
    user:
      exec:
        apiVersion: client.authentication.k8s.io/v1
        interactiveMode: Never
        command: kube-client-go-azure-exec-plugin
        args:
          - %s
          - %s
          - %s
`
