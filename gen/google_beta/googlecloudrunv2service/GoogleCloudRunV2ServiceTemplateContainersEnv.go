package googlecloudrunv2service


type GoogleCloudRunV2ServiceTemplateContainersEnv struct {
	// Name of the environment variable. Must be a C_IDENTIFIER, and mnay not exceed 32768 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_cloud_run_v2_service#name GoogleCloudRunV2Service#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables.
	//
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to "", and the maximum length is 32768 bytes
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_cloud_run_v2_service#value GoogleCloudRunV2Service#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
	// value_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_cloud_run_v2_service#value_source GoogleCloudRunV2Service#value_source}
	ValueSource *GoogleCloudRunV2ServiceTemplateContainersEnvValueSource `field:"optional" json:"valueSource" yaml:"valueSource"`
}

