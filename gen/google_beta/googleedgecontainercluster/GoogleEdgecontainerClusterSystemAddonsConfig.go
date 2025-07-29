package googleedgecontainercluster


type GoogleEdgecontainerClusterSystemAddonsConfig struct {
	// ingress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_edgecontainer_cluster#ingress GoogleEdgecontainerCluster#ingress}
	Ingress *GoogleEdgecontainerClusterSystemAddonsConfigIngress `field:"optional" json:"ingress" yaml:"ingress"`
}

