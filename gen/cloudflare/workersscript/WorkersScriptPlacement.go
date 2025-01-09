package workersscript


type WorkersScriptPlacement struct {
	// The placement mode for the Worker. Available values: `smart`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/workers_script#mode WorkersScript#mode}
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

