package accesscontextmanagerserviceperimeterdryrunegresspolicy


type AccessContextManagerServicePerimeterDryRunEgressPolicyEgressFromSources struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/access_context_manager_service_perimeter_dry_run_egress_policy#access_level AccessContextManagerServicePerimeterDryRunEgressPolicy#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
}

