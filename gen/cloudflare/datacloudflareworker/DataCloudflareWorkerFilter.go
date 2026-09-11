package datacloudflareworker


type DataCloudflareWorkerFilter struct {
	// Sort direction. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/worker#order DataCloudflareWorker#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Property to sort results by. Available values: "deployed_on", "updated_on", "created_on", "name".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/worker#order_by DataCloudflareWorker#order_by}
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
}

