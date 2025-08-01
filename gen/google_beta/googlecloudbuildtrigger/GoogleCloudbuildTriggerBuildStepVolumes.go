package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildStepVolumes struct {
	// Name of the volume to mount.
	//
	// Volume names must be unique per build step and must be valid names for
	// Docker volumes. Each named volume must be used by at least two build steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#name GoogleCloudbuildTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Path at which to mount the volume.
	//
	// Paths must be absolute and cannot conflict with other volume paths on
	// the same build step or with certain reserved volume paths.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#path GoogleCloudbuildTrigger#path}
	Path *string `field:"required" json:"path" yaml:"path"`
}

