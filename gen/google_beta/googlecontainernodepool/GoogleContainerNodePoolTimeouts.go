package googlecontainernodepool


type GoogleContainerNodePoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_container_node_pool#create GoogleContainerNodePool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_container_node_pool#delete GoogleContainerNodePool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_container_node_pool#update GoogleContainerNodePool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

