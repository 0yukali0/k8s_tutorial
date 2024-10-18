package controller

import (
	"flag"

	"k8s_monitor/pkg/common"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset  *kubernetes.Clientset
	controller *Controller
)

func init() {
	var kubeconfig string
	var master string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	flag.StringVar(&master, "master", "", "master url")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		common.GetLogger().Fatal(err.Error())
	}

	if clientset, err = kubernetes.NewForConfig(config); err != nil {
		common.GetLogger().Fatal(err.Error())
	}
}

func GetClientSet() *kubernetes.Clientset {
	return clientset
}
