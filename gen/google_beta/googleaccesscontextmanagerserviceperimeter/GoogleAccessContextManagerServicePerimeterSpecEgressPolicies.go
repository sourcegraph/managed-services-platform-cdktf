package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterSpecEgressPolicies struct {
	// egress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_access_context_manager_service_perimeter#egress_from GoogleAccessContextManagerServicePerimeter#egress_from}
	EgressFrom *GoogleAccessContextManagerServicePerimeterSpecEgressPoliciesEgressFrom `field:"optional" json:"egressFrom" yaml:"egressFrom"`
	// egress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_access_context_manager_service_perimeter#egress_to GoogleAccessContextManagerServicePerimeter#egress_to}
	EgressTo *GoogleAccessContextManagerServicePerimeterSpecEgressPoliciesEgressTo `field:"optional" json:"egressTo" yaml:"egressTo"`
	// Human readable title. Must be unique within the perimeter. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_access_context_manager_service_perimeter#title GoogleAccessContextManagerServicePerimeter#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

