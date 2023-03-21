package k8s

import (
	"context"
	"fmt"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var K8sCli *K8sClient

type K8sClient struct {
	ClientSet *kubernetes.Clientset
}

type Config struct {
	Host  string
	Token string
	Port  int64
}

func (k *K8sClient) NewKubernetesClient(c *Config) error {
	kubeConf := &rest.Config{
		Host:        fmt.Sprintf("%s:%d", c.Host, c.Port),
		BearerToken: c.Token,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	clientSet, err := kubernetes.NewForConfig(kubeConf)
	if err != nil {
		return err
	}

	k.ClientSet = clientSet

	return nil
}

// GetAllNamespace get all namespace in cluster.
func (k *K8sClient) GetAllNamespace() []string {
	var namespaces []string
	namespaceList, err := k.ClientSet.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		return namespaces
	}

	for _, nsList := range namespaceList.Items {
		namespaces = append(namespaces, nsList.Name)
	}

	return namespaces
}

// GetAllPodByNs get all pod in cluster by namespace.
func (k *K8sClient) GetAllPodByNs(namespace string) []string {
	var pods []string
	podsList, err := k.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})

	if err != nil {
		return pods
	}

	for _, nsList := range podsList.Items {
		pods = append(pods, nsList.Name)
	}

	return pods
}
