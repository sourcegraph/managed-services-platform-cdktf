package pagesproject


type PagesProjectDeploymentConfigs struct {
	// Configs for preview deploys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/pages_project#preview PagesProject#preview}
	Preview *PagesProjectDeploymentConfigsPreview `field:"optional" json:"preview" yaml:"preview"`
	// Configs for production deploys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/pages_project#production PagesProject#production}
	Production *PagesProjectDeploymentConfigsProduction `field:"optional" json:"production" yaml:"production"`
}

