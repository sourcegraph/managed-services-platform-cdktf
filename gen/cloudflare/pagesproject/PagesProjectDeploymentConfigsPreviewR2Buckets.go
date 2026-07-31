package pagesproject


type PagesProjectDeploymentConfigsPreviewR2Buckets struct {
	// Name of the R2 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#name PagesProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Jurisdiction of the R2 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#jurisdiction PagesProject#jurisdiction}
	Jurisdiction *string `field:"optional" json:"jurisdiction" yaml:"jurisdiction"`
}

