package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// NewForConfig creates a new Clientset for the given config.
	// https://godoc.org/k8s.io/client-go/kubernetes#NewForConfig
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// We need to create a serializer client to let us access API objects.
	// Type Clientset, from package kubernetes, provides access to generated serializer clients
	// to access versioned API objects
	// https://godoc.org/k8s.io/client-go/kubernetes#Clientset.CoreV1
	ctx := context.TODO()

	pods, err := clientset.CoreV1().Pods("kube-system").List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}

	// print pods
	// pods.Items: []v1.Pod
	for i, pod := range pods.Items {
		fmt.Printf("[%d] %s\n", i, pod.GetName())
	}
}
