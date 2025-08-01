package googlecloudfunctions2function


type GoogleCloudfunctions2FunctionServiceConfig struct {
	// Whether 100% of traffic is routed to the latest revision. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#all_traffic_on_latest_revision GoogleCloudfunctions2Function#all_traffic_on_latest_revision}
	AllTrafficOnLatestRevision interface{} `field:"optional" json:"allTrafficOnLatestRevision" yaml:"allTrafficOnLatestRevision"`
	// The number of CPUs used in a single container instance. Default value is calculated from available memory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#available_cpu GoogleCloudfunctions2Function#available_cpu}
	AvailableCpu *string `field:"optional" json:"availableCpu" yaml:"availableCpu"`
	// The amount of memory available for a function.
	//
	// Defaults to 256M. Supported units are k, M, G, Mi, Gi. If no unit is
	// supplied the value is interpreted as bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#available_memory GoogleCloudfunctions2Function#available_memory}
	AvailableMemory *string `field:"optional" json:"availableMemory" yaml:"availableMemory"`
	// The binary authorization policy to be checked when deploying the Cloud Run service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#binary_authorization_policy GoogleCloudfunctions2Function#binary_authorization_policy}
	BinaryAuthorizationPolicy *string `field:"optional" json:"binaryAuthorizationPolicy" yaml:"binaryAuthorizationPolicy"`
	// Environment variables that shall be available during function execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#environment_variables GoogleCloudfunctions2Function#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Available ingress settings. Defaults to "ALLOW_ALL" if unspecified. Default value: "ALLOW_ALL" Possible values: ["ALLOW_ALL", "ALLOW_INTERNAL_ONLY", "ALLOW_INTERNAL_AND_GCLB"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#ingress_settings GoogleCloudfunctions2Function#ingress_settings}
	IngressSettings *string `field:"optional" json:"ingressSettings" yaml:"ingressSettings"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#max_instance_count GoogleCloudfunctions2Function#max_instance_count}
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Sets the maximum number of concurrent requests that each instance can receive. Defaults to 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#max_instance_request_concurrency GoogleCloudfunctions2Function#max_instance_request_concurrency}
	MaxInstanceRequestConcurrency *float64 `field:"optional" json:"maxInstanceRequestConcurrency" yaml:"maxInstanceRequestConcurrency"`
	// The limit on the minimum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#min_instance_count GoogleCloudfunctions2Function#min_instance_count}
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// secret_environment_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#secret_environment_variables GoogleCloudfunctions2Function#secret_environment_variables}
	SecretEnvironmentVariables interface{} `field:"optional" json:"secretEnvironmentVariables" yaml:"secretEnvironmentVariables"`
	// secret_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#secret_volumes GoogleCloudfunctions2Function#secret_volumes}
	SecretVolumes interface{} `field:"optional" json:"secretVolumes" yaml:"secretVolumes"`
	// Name of the service associated with a Function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#service GoogleCloudfunctions2Function#service}
	Service *string `field:"optional" json:"service" yaml:"service"`
	// The email of the service account for this function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#service_account_email GoogleCloudfunctions2Function#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The function execution timeout.
	//
	// Execution is considered failed and
	// can be terminated if the function is not completed at the end of the
	// timeout period. Defaults to 60 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#timeout_seconds GoogleCloudfunctions2Function#timeout_seconds}
	TimeoutSeconds *float64 `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
	// The Serverless VPC Access connector that this cloud function can connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#vpc_connector GoogleCloudfunctions2Function#vpc_connector}
	VpcConnector *string `field:"optional" json:"vpcConnector" yaml:"vpcConnector"`
	// Available egress settings. Possible values: ["VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED", "PRIVATE_RANGES_ONLY", "ALL_TRAFFIC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions2_function#vpc_connector_egress_settings GoogleCloudfunctions2Function#vpc_connector_egress_settings}
	VpcConnectorEgressSettings *string `field:"optional" json:"vpcConnectorEgressSettings" yaml:"vpcConnectorEgressSettings"`
}

