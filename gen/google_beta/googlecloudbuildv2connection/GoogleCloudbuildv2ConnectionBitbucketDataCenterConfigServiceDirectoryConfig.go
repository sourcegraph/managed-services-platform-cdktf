package googlecloudbuildv2connection


type GoogleCloudbuildv2ConnectionBitbucketDataCenterConfigServiceDirectoryConfig struct {
	// Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_connection#service GoogleCloudbuildv2Connection#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

