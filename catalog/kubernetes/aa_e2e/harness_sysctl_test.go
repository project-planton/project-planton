package aa_e2e

import (
	"reflect"
	"testing"
)

// The pod-certificates profile must carry upstream's kind node settings, and
// only the IPv6 one may be skipped on a kernel without it.
func TestPodCertificatesNodeSysctls(t *testing.T) {
	want := [][]string{
		{"exec", "node-1", "sysctl", "-w", "net.ipv4.conf.all.proxy_arp=1"},
		{"exec", "node-1", "sysctl", "-e", "-w", "net.ipv6.conf.all.proxy_ndp=1"},
	}
	var got [][]string
	for _, s := range podCertificatesNodeSysctls {
		got = append(got, s.dockerExecArgs("node-1"))
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("pod-certificates node sysctls = %q, want %q (upstream's hack/create-kind-cluster.sh sets proxy_arp, and proxy_ndp with sysctl -e)", got, want)
	}
	if h := NewPodCertificatesHarness("planton-e2e"); !reflect.DeepEqual(h.nodeSysctls, podCertificatesNodeSysctls) {
		t.Errorf("the pod-certificates harness carries node sysctls %v, want the profile's %v", h.nodeSysctls, podCertificatesNodeSysctls)
	}
}
