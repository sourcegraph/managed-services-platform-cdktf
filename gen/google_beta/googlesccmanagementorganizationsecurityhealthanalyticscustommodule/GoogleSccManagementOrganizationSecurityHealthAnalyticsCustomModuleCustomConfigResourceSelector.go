package googlesccmanagementorganizationsecurityhealthanalyticscustommodule


type GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector struct {
	// The resource types to run the detector on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_scc_management_organization_security_health_analytics_custom_module#resource_types GoogleSccManagementOrganizationSecurityHealthAnalyticsCustomModule#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

