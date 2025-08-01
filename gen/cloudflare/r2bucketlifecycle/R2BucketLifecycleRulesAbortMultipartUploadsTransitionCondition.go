package r2bucketlifecycle


type R2BucketLifecycleRulesAbortMultipartUploadsTransitionCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#max_age R2BucketLifecycle#max_age}.
	MaxAge *float64 `field:"required" json:"maxAge" yaml:"maxAge"`
	// Available values: "Age".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#type R2BucketLifecycle#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

