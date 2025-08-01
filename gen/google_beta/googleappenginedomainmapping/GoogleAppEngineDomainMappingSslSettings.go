package googleappenginedomainmapping


type GoogleAppEngineDomainMappingSslSettings struct {
	// SSL management type for this domain.
	//
	// If 'AUTOMATIC', a managed certificate is automatically provisioned.
	// If 'MANUAL', 'certificateId' must be manually specified in order to configure SSL for this domain. Possible values: ["AUTOMATIC", "MANUAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_domain_mapping#ssl_management_type GoogleAppEngineDomainMapping#ssl_management_type}
	SslManagementType *string `field:"required" json:"sslManagementType" yaml:"sslManagementType"`
	// ID of the AuthorizedCertificate resource configuring SSL for the application.
	//
	// Clearing this field will
	// remove SSL support.
	// By default, a managed certificate is automatically created for every domain mapping. To omit SSL support
	// or to configure SSL manually, specify 'SslManagementType.MANUAL' on a 'CREATE' or 'UPDATE' request. You must be
	// authorized to administer the 'AuthorizedCertificate' resource to manually map it to a DomainMapping resource.
	// Example: 12345.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_app_engine_domain_mapping#certificate_id GoogleAppEngineDomainMapping#certificate_id}
	CertificateId *string `field:"optional" json:"certificateId" yaml:"certificateId"`
}

