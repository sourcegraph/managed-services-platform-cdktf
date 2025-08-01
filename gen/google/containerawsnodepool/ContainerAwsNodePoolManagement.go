package containerawsnodepool


type ContainerAwsNodePoolManagement struct {
	// Optional. Whether or not the nodes will be automatically repaired.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/container_aws_node_pool#auto_repair ContainerAwsNodePool#auto_repair}
	AutoRepair interface{} `field:"optional" json:"autoRepair" yaml:"autoRepair"`
}

