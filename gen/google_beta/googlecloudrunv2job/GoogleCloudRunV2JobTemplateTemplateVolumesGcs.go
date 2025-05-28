package googlecloudrunv2job


type GoogleCloudRunV2JobTemplateTemplateVolumesGcs struct {
	// Name of the cloud storage bucket to back the volume.
	//
	// The resource service account must have permission to access the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloud_run_v2_job#bucket GoogleCloudRunV2Job#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// A list of flags to pass to the gcsfuse command for configuring this volume.
	//
	// Flags should be passed without leading dashes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloud_run_v2_job#mount_options GoogleCloudRunV2Job#mount_options}
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// If true, mount this volume as read-only in all mounts. If false, mount this volume as read-write.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloud_run_v2_job#read_only GoogleCloudRunV2Job#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
}

