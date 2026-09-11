package snippet


type SnippetMetadata struct {
	// Specify the name of the file that contains the main module of the snippet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/snippet#main_module Snippet#main_module}
	MainModule *string `field:"required" json:"mainModule" yaml:"mainModule"`
}

