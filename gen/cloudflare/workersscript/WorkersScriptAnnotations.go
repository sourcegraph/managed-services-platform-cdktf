package workersscript


type WorkersScriptAnnotations struct {
	// Human-readable message about the version. Truncated to 1000 bytes if longer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#workers_message WorkersScript#workers_message}
	WorkersMessage *string `field:"optional" json:"workersMessage" yaml:"workersMessage"`
	// User-provided identifier for the version. Maximum 100 bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#workers_tag WorkersScript#workers_tag}
	WorkersTag *string `field:"optional" json:"workersTag" yaml:"workersTag"`
}

