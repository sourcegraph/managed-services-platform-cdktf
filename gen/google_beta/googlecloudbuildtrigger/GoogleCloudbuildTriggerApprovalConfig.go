package googlecloudbuildtrigger


type GoogleCloudbuildTriggerApprovalConfig struct {
	// Whether or not approval is needed.
	//
	// If this is set on a build, it will become pending when run,
	// and will need to be explicitly approved to start.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#approval_required GoogleCloudbuildTrigger#approval_required}
	ApprovalRequired interface{} `field:"optional" json:"approvalRequired" yaml:"approvalRequired"`
}

