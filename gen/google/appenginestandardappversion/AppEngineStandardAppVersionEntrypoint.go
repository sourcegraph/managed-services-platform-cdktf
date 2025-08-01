package appenginestandardappversion


type AppEngineStandardAppVersionEntrypoint struct {
	// The format should be a shell command that can be fed to bash -c.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/app_engine_standard_app_version#shell AppEngineStandardAppVersion#shell}
	Shell *string `field:"required" json:"shell" yaml:"shell"`
}

