package snippets


type SnippetsMetadata struct {
	// Name of the file that contains the main module of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/snippets#main_module Snippets#main_module}
	MainModule *string `field:"required" json:"mainModule" yaml:"mainModule"`
}

