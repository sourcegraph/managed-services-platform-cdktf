package snippet


type SnippetFiles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippet#content Snippet#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/snippet#name Snippet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

