package containercluster


type ContainerClusterLoggingConfig struct {
	// GKE components exposing logs. Valid values include SYSTEM_COMPONENTS, APISERVER, CONTROLLER_MANAGER, KCP_CONNECTION, KCP_SSHD, KCP_HPA, SCHEDULER, and WORKLOADS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_cluster#enable_components ContainerCluster#enable_components}
	EnableComponents *[]*string `field:"required" json:"enableComponents" yaml:"enableComponents"`
}

