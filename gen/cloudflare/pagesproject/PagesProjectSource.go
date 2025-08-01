package pagesproject


type PagesProjectSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/pages_project#config PagesProject#config}.
	Config *PagesProjectSourceConfig `field:"optional" json:"config" yaml:"config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/pages_project#type PagesProject#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

