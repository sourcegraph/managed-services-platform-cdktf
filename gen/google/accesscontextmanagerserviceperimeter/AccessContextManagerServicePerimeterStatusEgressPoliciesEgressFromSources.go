package accesscontextmanagerserviceperimeter


type AccessContextManagerServicePerimeterStatusEgressPoliciesEgressFromSources struct {
	// An AccessLevel resource name that allows resources outside the ServicePerimeter to be accessed from the inside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_service_perimeter#access_level AccessContextManagerServicePerimeter#access_level}
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// A Google Cloud resource that is allowed to egress the perimeter.
	//
	// Requests from these resources are allowed to access data outside the perimeter.
	// Currently only projects are allowed. Project format: 'projects/{project_number}'.
	// The resource may be in any Google Cloud organization, not just the
	// organization that the perimeter is defined in. '*' is not allowed, the
	// case of allowing all Google Cloud resources only is not supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_service_perimeter#resource AccessContextManagerServicePerimeter#resource}
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

