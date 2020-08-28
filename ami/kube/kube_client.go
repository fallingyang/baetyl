package kube

import (
	"github.com/baetyl/baetyl-go/v2/errors"
	"k8s.io/client-go/kubernetes"
	appv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientset "k8s.io/metrics/pkg/client/clientset/versioned"
	metricsv1beta1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1beta1"

	"github.com/baetyl/baetyl/config"
)

type client struct {
	core    corev1.CoreV1Interface
	app     appv1.AppsV1Interface
	metrics metricsv1beta1.MetricsV1beta1Interface
}

func newClient(cfg config.KubeConfig) (*client, error) {
	kubeConfig, err := func() (*rest.Config, error) {
		if !cfg.OutCluster {
			return rest.InClusterConfig()
		}
		return clientcmd.BuildConfigFromFlags("", cfg.ConfPath)
	}()
	if err != nil {
		return nil, errors.Trace(err)
	}
	kubeClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, errors.Trace(err)
	}

	metricsCli, err := clientset.NewForConfig(kubeConfig)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return &client{
		core:    kubeClient.CoreV1(),
		app:     kubeClient.AppsV1(),
		metrics: metricsCli.MetricsV1beta1(),
	}, nil
}
