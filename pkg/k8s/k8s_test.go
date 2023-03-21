package k8s

import (
	"fmt"
	"testing"
)

func TestK8s(t *testing.T) {
	InitK8s()
	for _, namespace := range K8sCli.GetAllNamespace() {
		fmt.Println(namespace)
	}
	for _, pod := range K8sCli.GetAllPodByNs("default") {
		fmt.Println(pod)
	}
}
