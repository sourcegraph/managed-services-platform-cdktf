package googlenetworkconnectivityinternalrange


type GoogleNetworkConnectivityInternalRangeMigration struct {
	// Resource path as an URI of the source resource, for example a subnet.
	//
	// The project for the source resource should match the project for the
	// InternalRange.
	// An example /projects/{project}/regions/{region}/subnetworks/{subnet}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_internal_range#source GoogleNetworkConnectivityInternalRange#source}
	Source *string `field:"required" json:"source" yaml:"source"`
	// Resource path of the target resource.
	//
	// The target project can be
	// different, as in the cases when migrating to peer networks. The resource
	// may not exist yet.
	// For example /projects/{project}/regions/{region}/subnetworks/{subnet}
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_internal_range#target GoogleNetworkConnectivityInternalRange#target}
	Target *string `field:"required" json:"target" yaml:"target"`
}

