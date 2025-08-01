package containerattachedcluster


type ContainerAttachedClusterProxyConfig struct {
	// kubernetes_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_attached_cluster#kubernetes_secret ContainerAttachedCluster#kubernetes_secret}
	KubernetesSecret *ContainerAttachedClusterProxyConfigKubernetesSecret `field:"optional" json:"kubernetesSecret" yaml:"kubernetesSecret"`
}

