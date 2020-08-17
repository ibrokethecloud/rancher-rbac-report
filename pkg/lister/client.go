package lister

import (
	managementv3 "github.com/rancher/types/apis/management.cattle.io/v3"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateClientset() (managementv3.Interface, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	managementClient, err := managementv3.NewForConfig(*config)

	if err != nil {
		return nil, err
	}

	return managementClient, nil
}
