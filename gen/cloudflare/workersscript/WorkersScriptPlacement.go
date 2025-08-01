package workersscript


type WorkersScriptPlacement struct {
	// Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Available values: "smart".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#mode WorkersScript#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

