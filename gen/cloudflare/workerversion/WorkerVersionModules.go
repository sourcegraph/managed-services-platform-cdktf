package workerversion


type WorkerVersionModules struct {
	// The content type of the module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#content_type WorkerVersion#content_type}
	ContentType *string `field:"required" json:"contentType" yaml:"contentType"`
	// The name of the module.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#name WorkerVersion#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The base64-encoded module content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#content_base64 WorkerVersion#content_base64}
	ContentBase64 *string `field:"optional" json:"contentBase64" yaml:"contentBase64"`
	// The file path of the module content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#content_file WorkerVersion#content_file}
	ContentFile *string `field:"optional" json:"contentFile" yaml:"contentFile"`
}

