package pagesproject


type PagesProjectSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#config PagesProject#config}.
	Config *PagesProjectSourceConfig `field:"required" json:"config" yaml:"config"`
	// The source control management provider. Available values: "github", "gitlab".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#type PagesProject#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

