package googlecomputebackendservice


type GoogleComputeBackendServiceLocalityLbPoliciesCustomPolicy struct {
	// Identifies the custom policy.
	//
	// The value should match the type the custom implementation is registered
	// with on the gRPC clients. It should follow protocol buffer
	// message naming conventions and include the full path (e.g.
	// myorg.CustomLbPolicy). The maximum length is 256 characters.
	//
	// Note that specifying the same custom policy more than once for a
	// backend is not a valid configuration and will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#name GoogleComputeBackendService#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional, arbitrary JSON object with configuration data, understood by a locally installed custom policy implementation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_backend_service#data GoogleComputeBackendService#data}
	Data *string `field:"optional" json:"data" yaml:"data"`
}

