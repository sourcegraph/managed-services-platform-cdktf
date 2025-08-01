package identityplatformconfig


type IdentityPlatformConfigMonitoring struct {
	// request_logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/identity_platform_config#request_logging IdentityPlatformConfig#request_logging}
	RequestLogging *IdentityPlatformConfigMonitoringRequestLogging `field:"optional" json:"requestLogging" yaml:"requestLogging"`
}

