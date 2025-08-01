package googlecloudrundomainmapping


type GoogleCloudRunDomainMappingSpec struct {
	// The name of the Cloud Run Service that this DomainMapping applies to. The route must exist.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_domain_mapping#route_name GoogleCloudRunDomainMapping#route_name}
	RouteName *string `field:"required" json:"routeName" yaml:"routeName"`
	// The mode of the certificate. Default value: "AUTOMATIC" Possible values: ["NONE", "AUTOMATIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_domain_mapping#certificate_mode GoogleCloudRunDomainMapping#certificate_mode}
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// If set, the mapping will override any mapping set before this spec was set.
	//
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_domain_mapping#force_override GoogleCloudRunDomainMapping#force_override}
	ForceOverride interface{} `field:"optional" json:"forceOverride" yaml:"forceOverride"`
}

