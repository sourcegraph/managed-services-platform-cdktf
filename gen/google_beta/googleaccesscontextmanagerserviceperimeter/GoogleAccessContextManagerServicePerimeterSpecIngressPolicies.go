package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterSpecIngressPolicies struct {
	// ingress_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_access_context_manager_service_perimeter#ingress_from GoogleAccessContextManagerServicePerimeter#ingress_from}
	IngressFrom *GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressFrom `field:"optional" json:"ingressFrom" yaml:"ingressFrom"`
	// ingress_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_access_context_manager_service_perimeter#ingress_to GoogleAccessContextManagerServicePerimeter#ingress_to}
	IngressTo *GoogleAccessContextManagerServicePerimeterSpecIngressPoliciesIngressTo `field:"optional" json:"ingressTo" yaml:"ingressTo"`
	// Human readable title. Must be unique within the perimeter. Does not affect behavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_access_context_manager_service_perimeter#title GoogleAccessContextManagerServicePerimeter#title}
	Title *string `field:"optional" json:"title" yaml:"title"`
}

