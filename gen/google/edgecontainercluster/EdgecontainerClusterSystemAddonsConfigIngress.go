package edgecontainercluster


type EdgecontainerClusterSystemAddonsConfigIngress struct {
	// Whether Ingress is disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#disabled EdgecontainerCluster#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Ingress VIP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/edgecontainer_cluster#ipv4_vip EdgecontainerCluster#ipv4_vip}
	Ipv4Vip *string `field:"optional" json:"ipv4Vip" yaml:"ipv4Vip"`
}

