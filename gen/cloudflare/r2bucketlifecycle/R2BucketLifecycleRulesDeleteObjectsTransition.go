package r2bucketlifecycle


type R2BucketLifecycleRulesDeleteObjectsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/r2_bucket_lifecycle#condition R2BucketLifecycle#condition}
	Condition *R2BucketLifecycleRulesDeleteObjectsTransitionCondition `field:"optional" json:"condition" yaml:"condition"`
}

