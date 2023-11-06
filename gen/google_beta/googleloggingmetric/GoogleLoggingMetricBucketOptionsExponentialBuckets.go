package googleloggingmetric


type GoogleLoggingMetricBucketOptionsExponentialBuckets struct {
	// Must be greater than 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_logging_metric#growth_factor GoogleLoggingMetric#growth_factor}
	GrowthFactor *float64 `field:"optional" json:"growthFactor" yaml:"growthFactor"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_logging_metric#num_finite_buckets GoogleLoggingMetric#num_finite_buckets}
	NumFiniteBuckets *float64 `field:"optional" json:"numFiniteBuckets" yaml:"numFiniteBuckets"`
	// Must be greater than 0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_logging_metric#scale GoogleLoggingMetric#scale}
	Scale *float64 `field:"optional" json:"scale" yaml:"scale"`
}

