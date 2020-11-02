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
	nodes, err := cl.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	log.Printf("There are %d nodes in the cluster\n", len(nodes.Items))
	if err != nil {
		log.Fatal(err)
	}

}

func podnslist(cl *kubernetes.Clientset, ns string) {

	// Retrieve the Corev1 Client via clientset and list all Nodes in the cluster
	pods, err := cl.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if ns != "" {
		log.Printf("There are %d pods in the %v namespace\n", len(pods.Items), ns)
	} else {
		log.Printf("There are %d pods in the cluster\n", len(pods.Items))
	}

	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	// Definition of the namespace flag
	var ns string
	flag.StringVar(&ns, "n", "", "the namespace to show")
	kube := flag.Bool("kubeconfig", false, "show the kubeconfig loaded")
	nodes := flag.Bool("nodes", false, "show the nodes")
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

	podnslist(clientset, ns)

}
