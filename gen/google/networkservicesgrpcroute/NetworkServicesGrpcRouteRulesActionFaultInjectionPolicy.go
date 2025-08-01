package networkservicesgrpcroute


type NetworkServicesGrpcRouteRulesActionFaultInjectionPolicy struct {
	// abort block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_grpc_route#abort NetworkServicesGrpcRoute#abort}
	Abort *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyAbort `field:"optional" json:"abort" yaml:"abort"`
	// delay block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_grpc_route#delay NetworkServicesGrpcRoute#delay}
	Delay *NetworkServicesGrpcRouteRulesActionFaultInjectionPolicyDelay `field:"optional" json:"delay" yaml:"delay"`
}

