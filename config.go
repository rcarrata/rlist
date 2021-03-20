package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func setupkube(kubeconfig string) (cl *kubernetes.Clientset, err error) {
	// Build config from master url and kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig) // "" used because the cluster is local
	if err != nil {
		log.Fatal(err)
	}

	// Create a clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return clientset, err

}
