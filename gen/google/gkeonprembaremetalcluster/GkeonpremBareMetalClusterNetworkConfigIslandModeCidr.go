package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterNetworkConfigIslandModeCidr struct {
	// All pods in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#pod_address_cidr_blocks GkeonpremBareMetalCluster#pod_address_cidr_blocks}
	PodAddressCidrBlocks *[]*string `field:"required" json:"podAddressCidrBlocks" yaml:"podAddressCidrBlocks"`
	// All services in the cluster are assigned an RFC1918 IPv4 address from these ranges.
	//
	// This field cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#service_address_cidr_blocks GkeonpremBareMetalCluster#service_address_cidr_blocks}
	ServiceAddressCidrBlocks *[]*string `field:"required" json:"serviceAddressCidrBlocks" yaml:"serviceAddressCidrBlocks"`
}

