package pagesproject


type PagesProjectDeploymentConfigsPreviewServices struct {
	// The entrypoint to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/pages_project#entrypoint PagesProject#entrypoint}
	Entrypoint *string `field:"optional" json:"entrypoint" yaml:"entrypoint"`
	// The Service environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/pages_project#environment PagesProject#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The Service name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/pages_project#service PagesProject#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
}

