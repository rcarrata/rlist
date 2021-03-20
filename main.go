package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	// Definition of the namespace flag
	var ns string
	flag.StringVar(&ns, "n", "", "the namespace to show")
	kube := flag.Bool("kubeconfig", false, "show the kubeconfig loaded")
	nodes := flag.Bool("nodes", false, "show the nodes")
	pods := flag.Bool("pods", false, "show the pods in the namespace / cluster")
	svcs := flag.Bool("svcs", false, "show the svcs in the namespace / cluster")
	flag.Parse()

	// filepath for select where is the kubeconfig file located
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	// Create the clientset of K8s
	clientset, err := setupkube(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// If its enabled, show the kubeconfig loaded
	if *kube {
		fmt.Println("Using kubeconfig default file: ", kubeconfig)
		fmt.Println("")
	}

	// If its enabled, list the nodes of the cluster
	if *nodes {
		nodelist(clientset)
	}

	if *pods {
		podnslist(clientset, ns)
	}

	if *svcs {
		svclist(clientset, ns)
	}

}
