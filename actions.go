package main

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

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

func svclist(cl *kubernetes.Clientset, ns string) {

	svcs, err := cl.CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})

	if ns != "" {
		log.Printf("There are %d services in the %v namespace\n", len(svcs.Items), ns)
	} else {
		log.Printf("There are %d services in the cluster\n", len(svcs.Items))
	}

	if err != nil {
		log.Fatal(err)
	}

}
