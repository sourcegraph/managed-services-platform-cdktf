package accesscontextmanagerserviceperimeteregresspolicy


type AccessContextManagerServicePerimeterEgressPolicyEgressTo struct {
	// A list of external resources that are allowed to be accessed.
	//
	// A request
	// matches if it contains an external resource in this list (Example:
	// s3://bucket/path). Currently '*' is not allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_service_perimeter_egress_policy#external_resources AccessContextManagerServicePerimeterEgressPolicy#external_resources}
	ExternalResources *[]*string `field:"optional" json:"externalResources" yaml:"externalResources"`
	// operations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_service_perimeter_egress_policy#operations AccessContextManagerServicePerimeterEgressPolicy#operations}
	Operations interface{} `field:"optional" json:"operations" yaml:"operations"`
	// A list of resources, currently only projects in the form 'projects/<projectnumber>', that match this to stanza.
	//
	// A request matches
	// if it contains a resource in this list. If * is specified for resources,
	// then this 'EgressTo' rule will authorize access to all resources outside
	// the perimeter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_service_perimeter_egress_policy#resources AccessContextManagerServicePerimeterEgressPolicy#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// A list of IAM roles that represent the set of operations that the sources specified in the corresponding 'EgressFrom' are allowed to perform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_service_perimeter_egress_policy#roles AccessContextManagerServicePerimeterEgressPolicy#roles}
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
}

