/**
SETUP:
start dcoker
minikube start
// create nginx pod
kubectl apply -f https://k8s.io/examples/pods/simple-pod.yaml
*/

package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func main() {

	kubeconfig := flag.String("kubeconfig", "/Users/akaneri/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("The clientset cannot be created: %v\n", err)
		os.Exit(1)
	}

	pod, err := clientset.CoreV1().Pods("default").Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("The pod query failed: %v\n", err)
		os.Exit(1)
	}
	print(pod.Status.PodIP)
}
