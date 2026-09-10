package datacloudflareworkflow


type DataCloudflareWorkflowFilter struct {
	// Allows filtering workflows` name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/workflow#search DataCloudflareWorkflow#search}
	Search *string `field:"optional" json:"search" yaml:"search"`
}

