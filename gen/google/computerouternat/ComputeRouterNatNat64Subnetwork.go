package computerouternat


type ComputeRouterNatNat64Subnetwork struct {
	// Self-link of the subnetwork resource that will use NAT64.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_router_nat#name ComputeRouterNat#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

