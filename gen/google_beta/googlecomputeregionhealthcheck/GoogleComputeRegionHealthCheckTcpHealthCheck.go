package googlecomputeregionhealthcheck


type GoogleComputeRegionHealthCheckTcpHealthCheck struct {
	// The TCP port number for the TCP health check request. The default value is 80.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_health_check#port GoogleComputeRegionHealthCheck#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Port name as defined in InstanceGroup#NamedPort#name. If both port and port_name are defined, port takes precedence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_health_check#port_name GoogleComputeRegionHealthCheck#port_name}
	PortName *string `field:"optional" json:"portName" yaml:"portName"`
	// Specifies how port is selected for health checking, can be one of the following values:.
	//
	// * 'USE_FIXED_PORT': The port number in 'port' is used for health checking.
	//
	//   * 'USE_NAMED_PORT': The 'portName' is used for health checking.
	//
	//   * 'USE_SERVING_PORT': For NetworkEndpointGroup, the port specified for each
	//   network endpoint is used for health checking. For other backends, the
	//   port or named port specified in the Backend Service is used for health
	//   checking.
	//
	// If not specified, TCP health check follows behavior specified in 'port' and
	// 'portName' fields. Possible values: ["USE_FIXED_PORT", "USE_NAMED_PORT", "USE_SERVING_PORT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_health_check#port_specification GoogleComputeRegionHealthCheck#port_specification}
	PortSpecification *string `field:"optional" json:"portSpecification" yaml:"portSpecification"`
	// Specifies the type of proxy header to append before sending data to the backend.
	//
	// Default value: "NONE" Possible values: ["NONE", "PROXY_V1"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_health_check#proxy_header GoogleComputeRegionHealthCheck#proxy_header}
	ProxyHeader *string `field:"optional" json:"proxyHeader" yaml:"proxyHeader"`
	// The application data to send once the TCP connection has been established (default value is empty).
	//
	// If both request and response are
	// empty, the connection establishment alone will indicate health. The request
	// data can only be ASCII.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_health_check#request GoogleComputeRegionHealthCheck#request}
	Request *string `field:"optional" json:"request" yaml:"request"`
	// The bytes to match against the beginning of the response data.
	//
	// If left empty
	// (the default value), any response will indicate health. The response data
	// can only be ASCII.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_health_check#response GoogleComputeRegionHealthCheck#response}
	Response *string `field:"optional" json:"response" yaml:"response"`
}

