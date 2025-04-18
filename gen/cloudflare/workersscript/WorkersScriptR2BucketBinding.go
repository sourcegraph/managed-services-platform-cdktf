package workersscript


type WorkersScriptR2BucketBinding struct {
	// The name of the Bucket to bind to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/workers_script#bucket_name WorkersScript#bucket_name}
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The global variable for the binding in your Worker code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/workers_script#name WorkersScript#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

