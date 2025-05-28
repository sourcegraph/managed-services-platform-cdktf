package googlecomputebackendservice


type GoogleComputeBackendServiceSecuritySettingsAwsV4Authentication struct {
	// The access key used for s3 bucket authentication.
	//
	// Required for updating or creating a backend that uses AWS v4 signature authentication, but will not be returned as part of the configuration when queried with a REST API GET request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_backend_service#access_key GoogleComputeBackendService#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// The identifier of an access key used for s3 bucket authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_backend_service#access_key_id GoogleComputeBackendService#access_key_id}
	AccessKeyId *string `field:"optional" json:"accessKeyId" yaml:"accessKeyId"`
	// The optional version identifier for the access key.
	//
	// You can use this to keep track of different iterations of your access key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_backend_service#access_key_version GoogleComputeBackendService#access_key_version}
	AccessKeyVersion *string `field:"optional" json:"accessKeyVersion" yaml:"accessKeyVersion"`
	// The name of the cloud region of your origin.
	//
	// This is a free-form field with the name of the region your cloud uses to host your origin.
	// For example, "us-east-1" for AWS or "us-ashburn-1" for OCI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_compute_backend_service#origin_region GoogleComputeBackendService#origin_region}
	OriginRegion *string `field:"optional" json:"originRegion" yaml:"originRegion"`
}

