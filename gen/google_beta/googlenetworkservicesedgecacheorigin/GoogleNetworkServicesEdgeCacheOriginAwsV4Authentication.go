package googlenetworkservicesedgecacheorigin


type GoogleNetworkServicesEdgeCacheOriginAwsV4Authentication struct {
	// The access key ID your origin uses to identify the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_network_services_edge_cache_origin#access_key_id GoogleNetworkServicesEdgeCacheOrigin#access_key_id}
	AccessKeyId *string `field:"required" json:"accessKeyId" yaml:"accessKeyId"`
	// The name of the AWS region that your origin is in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_network_services_edge_cache_origin#origin_region GoogleNetworkServicesEdgeCacheOrigin#origin_region}
	OriginRegion *string `field:"required" json:"originRegion" yaml:"originRegion"`
	// The Secret Manager secret version of the secret access key used by your origin.
	//
	// This is the resource name of the secret version in the format 'projects/*\/secrets/*\/versions/*' where the '*' values are replaced by the project, secret, and version you require.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_network_services_edge_cache_origin#secret_access_key_version GoogleNetworkServicesEdgeCacheOrigin#secret_access_key_version}
	SecretAccessKeyVersion *string `field:"required" json:"secretAccessKeyVersion" yaml:"secretAccessKeyVersion"`
}

