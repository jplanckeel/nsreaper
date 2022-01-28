package internal

import (
	"context"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func clientk8s() (*kubernetes.Clientset, context.Context, error) {

	pathOptions := clientcmd.NewDefaultPathOptions()
	pathOptions.LoadingRules.DoNotResolvePaths = false
	config, err := pathOptions.GetStartingConfig()
	if err != nil {
		return nil, nil, err
	}

	configOverrides := clientcmd.ConfigOverrides{}
	clientConfig := clientcmd.NewDefaultClientConfig(*config, &configOverrides)
	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, nil, err
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, nil, err
	}

	ctx := context.TODO()

	return clientset, ctx, nil
}
