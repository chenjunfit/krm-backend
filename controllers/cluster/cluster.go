package cluster

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ClusterConfig struct {
	ClusterInfo
	KubeConfig string `json:"kubeConfig"`
}
type ClusterInfo struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"` //集群别名
	City        string `json:"city"`
	District    string `json:"district"`
}
type ClusterStatus struct {
	ClusterInfo
	ClusterVersion string `json:"clusterVersion"`
	Status         string `json:"status"`
}

func (cluster *ClusterConfig) GetClusterStatus() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = cluster.ClusterInfo
	var clientset *kubernetes.Clientset
	resetConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(cluster.KubeConfig))
	if err != nil {
		return clusterStatus, err
	}
	clientset, err = kubernetes.NewForConfig(resetConfig)
	if err != nil {
		return clusterStatus, err
	}
	version, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, err
	}
	clusterStatus.Status = "Active"
	clusterStatus.ClusterVersion = version.String()
	return clusterStatus, nil
}
