package workersscript


type WorkersScriptAssets struct {
	// Configuration for assets within a Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#config WorkersScript#config}
	Config *WorkersScriptAssetsConfig `field:"optional" json:"config" yaml:"config"`
	// Path to the directory containing asset files to upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#directory WorkersScript#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Token provided upon successful upload of all files from a registered manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_script#jwt WorkersScript#jwt}
	Jwt *string `field:"optional" json:"jwt" yaml:"jwt"`
}

