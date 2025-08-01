package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildOptions struct {
	// Requested disk size for the VM that runs the build.
	//
	// Note that this is NOT "disk free";
	// some of the space will be used by the operating system and build utilities.
	// Also note that this is the minimum disk size that will be allocated for the build --
	// the build may run with a larger disk than requested. At present, the maximum disk size
	// is 1000GB; builds that request more than the maximum are rejected with an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#disk_size_gb GoogleCloudbuildTrigger#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Option to specify whether or not to apply bash style string operations to the substitutions.
	//
	// NOTE this is always enabled for triggered builds and cannot be overridden in the build configuration file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#dynamic_substitutions GoogleCloudbuildTrigger#dynamic_substitutions}
	DynamicSubstitutions interface{} `field:"optional" json:"dynamicSubstitutions" yaml:"dynamicSubstitutions"`
	// A list of global environment variable definitions that will exist for all build steps in this build.
	//
	// If a variable is defined in both globally and in a build step,
	// the variable will use the build step value.
	//
	// The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#env GoogleCloudbuildTrigger#env}
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// Option to specify the logging mode, which determines if and where build logs are stored.
	//
	// Possible values: ["LOGGING_UNSPECIFIED", "LEGACY", "GCS_ONLY", "STACKDRIVER_ONLY", "CLOUD_LOGGING_ONLY", "NONE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#logging GoogleCloudbuildTrigger#logging}
	Logging *string `field:"optional" json:"logging" yaml:"logging"`
	// Option to define build log streaming behavior to Google Cloud Storage. Possible values: ["STREAM_DEFAULT", "STREAM_ON", "STREAM_OFF"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#log_streaming_option GoogleCloudbuildTrigger#log_streaming_option}
	LogStreamingOption *string `field:"optional" json:"logStreamingOption" yaml:"logStreamingOption"`
	// Compute Engine machine type on which to run the build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#machine_type GoogleCloudbuildTrigger#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// Requested verifiability options. Possible values: ["NOT_VERIFIED", "VERIFIED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#requested_verify_option GoogleCloudbuildTrigger#requested_verify_option}
	RequestedVerifyOption *string `field:"optional" json:"requestedVerifyOption" yaml:"requestedVerifyOption"`
	// A list of global environment variables, which are encrypted using a Cloud Key Management Service crypto key.
	//
	// These values must be specified in the build's Secret. These variables
	// will be available to all build steps in this build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#secret_env GoogleCloudbuildTrigger#secret_env}
	SecretEnv *[]*string `field:"optional" json:"secretEnv" yaml:"secretEnv"`
	// Requested hash for SourceProvenance. Possible values: ["NONE", "SHA256", "MD5"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#source_provenance_hash GoogleCloudbuildTrigger#source_provenance_hash}
	SourceProvenanceHash *[]*string `field:"optional" json:"sourceProvenanceHash" yaml:"sourceProvenanceHash"`
	// Option to specify behavior when there is an error in the substitution checks.
	//
	// NOTE this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden
	// in the build configuration file. Possible values: ["MUST_MATCH", "ALLOW_LOOSE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#substitution_option GoogleCloudbuildTrigger#substitution_option}
	SubstitutionOption *string `field:"optional" json:"substitutionOption" yaml:"substitutionOption"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#volumes GoogleCloudbuildTrigger#volumes}
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
	// Option to specify a WorkerPool for the build. Format projects/{project}/workerPools/{workerPool}.
	//
	// This field is experimental.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#worker_pool GoogleCloudbuildTrigger#worker_pool}
	WorkerPool *string `field:"optional" json:"workerPool" yaml:"workerPool"`
}

