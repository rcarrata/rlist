package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

func nodelist(cl *kubernetes.Clientset) {

	// Retrieve the Corev1 Client via clientset and list all Nodes in the cluster
	pods, err := cl.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	log.Printf("There are %d nodes in the cluster\n", len(pods.Items))
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	// Definition of the namespace flag
	var ns string
	flag.StringVar(&ns, "namespace", "", "a namespace")
	flag.Parse()

	// filepath for select where is the kubeconfig file located
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	// Logging to stdout
	// https://stackoverflow.com/questions/19646889/why-should-i-use-log-println-instead-of-fmt-println
	fmt.Println("Using kubeconfig default file: ", kubeconfig) // why use log instead of fmt
	fmt.Println("")

	clientset, err := setupkube(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// List the nodes of the cluster
	nodelist(clientset)

}
