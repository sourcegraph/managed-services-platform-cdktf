package networkserviceslbtrafficextension


type NetworkServicesLbTrafficExtensionExtensionChainsMatchCondition struct {
	// A Common Expression Language (CEL) expression that is used to match requests for which the extension chain is executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_lb_traffic_extension#cel_expression NetworkServicesLbTrafficExtension#cel_expression}
	CelExpression *string `field:"required" json:"celExpression" yaml:"celExpression"`
}

