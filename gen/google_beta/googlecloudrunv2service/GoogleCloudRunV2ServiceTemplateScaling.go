package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateScaling struct {
	// Maximum number of serving instances that this resource should have.
	//
	// Must not be less than minimum instance count. If absent, Cloud Run will calculate
	// a default value based on the project's available container instances quota in the region and specified instance size.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_cloud_run_v2_service#max_instance_count GoogleCloudRunV2Service#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Minimum number of serving instances that this resource should have.
	//
	// Defaults to 0. Must not be greater than maximum instance count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_cloud_run_v2_service#min_instance_count GoogleCloudRunV2Service#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
}

