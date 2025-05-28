package googleaccesscontextmanagerserviceperimeteringresspolicy


type GoogleAccessContextManagerServicePerimeterIngressPolicyIngressToOperations struct {
	// method_selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_access_context_manager_service_perimeter_ingress_policy#method_selectors GoogleAccessContextManagerServicePerimeterIngressPolicy#method_selectors}
	MethodSelectors interface{} `field:"optional" json:"methodSelectors" yaml:"methodSelectors"`
	// The name of the API whose methods or permissions the 'IngressPolicy' or 'EgressPolicy' want to allow.
	//
	// A single 'ApiOperation' with 'serviceName'
	// field set to '*' will allow all methods AND permissions for all services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_access_context_manager_service_perimeter_ingress_policy#service_name GoogleAccessContextManagerServicePerimeterIngressPolicy#service_name}
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

