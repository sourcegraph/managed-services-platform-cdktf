package workerversion


type WorkerVersionAssets struct {
	// Configuration for assets within a Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#config WorkerVersion#config}
	Config *WorkerVersionAssetsConfig `field:"optional" json:"config" yaml:"config"`
	// Path to the directory containing asset files to upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#directory WorkerVersion#directory}
	Directory *string `field:"optional" json:"directory" yaml:"directory"`
	// Token provided upon successful upload of all files from a registered manifest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#jwt WorkerVersion#jwt}
	Jwt *string `field:"optional" json:"jwt" yaml:"jwt"`
}

