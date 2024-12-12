package pulumigkekubernetesprovider

const (
	GcpExecPluginPath = "/usr/local/bin/kube-client-go-gcp-exec-plugin"
)

// GcpExecPluginKubeConfigTemplate requires the following inputs for rendering a kubeconfig that works
// 1. cluster endpoint ip
// 2. cluster cert-authority data
// 3. base64 encoded google service account key
const GcpExecPluginKubeConfigTemplate = `apiVersion: v1
kind: Config
current-context: kube-context
contexts:
- name: kube-context
  context: {cluster: gke-cluster, user: kube-user}
clusters:
- name: gke-cluster
  cluster:
    server: https://%s
    certificate-authority-data: %s
users:
- name: kube-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1
      interactiveMode: Never
      command: %s
      args:
        - %s
`
