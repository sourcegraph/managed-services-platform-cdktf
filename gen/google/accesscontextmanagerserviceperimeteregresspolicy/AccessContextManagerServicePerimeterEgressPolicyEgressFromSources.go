package accesscontextmanagerserviceperimeteregresspolicy


type AccessContextManagerServicePerimeterEgressPolicyEgressFromSources struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/access_context_manager_service_perimeter_egress_policy#access_level AccessContextManagerServicePerimeterEgressPolicy#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
}

