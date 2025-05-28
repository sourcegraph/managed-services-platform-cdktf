package googlesccmanagementprojectsecurityhealthanalyticscustommodule


type GoogleSccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigResourceSelector struct {
	// The resource types to run the detector on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_scc_management_project_security_health_analytics_custom_module#resource_types GoogleSccManagementProjectSecurityHealthAnalyticsCustomModule#resource_types}
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
}

