package googlecloudrunv2service


type GoogleCloudRunV2ServiceBinaryAuthorization struct {
	// If present, indicates to use Breakglass using this justification.
	//
	// If useDefault is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_cloud_run_v2_service#breakglass_justification GoogleCloudRunV2Service#breakglass_justification}
	BreakglassJustification *string `field:"optional" json:"breakglassJustification" yaml:"breakglassJustification"`
	// If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_cloud_run_v2_service#use_default GoogleCloudRunV2Service#use_default}
	UseDefault interface{} `field:"optional" json:"useDefault" yaml:"useDefault"`
}

