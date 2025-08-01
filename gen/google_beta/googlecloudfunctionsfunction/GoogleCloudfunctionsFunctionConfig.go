package googlecloudfunctionsfunction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudfunctionsFunctionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// A user-defined name of the function. Function names must be unique globally.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#name GoogleCloudfunctionsFunction#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The runtime in which the function is going to run. Eg. "nodejs20", "python37", "go111".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#runtime GoogleCloudfunctionsFunction#runtime}
	Runtime *string `field:"required" json:"runtime" yaml:"runtime"`
	// Memory (in MB), available to the function. Default value is 256. Possible values include 128, 256, 512, 1024, etc.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#available_memory_mb GoogleCloudfunctionsFunction#available_memory_mb}
	AvailableMemoryMb *float64 `field:"optional" json:"availableMemoryMb" yaml:"availableMemoryMb"`
	// A set of key/value environment variable pairs available during build time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#build_environment_variables GoogleCloudfunctionsFunction#build_environment_variables}
	BuildEnvironmentVariables *map[string]*string `field:"optional" json:"buildEnvironmentVariables" yaml:"buildEnvironmentVariables"`
	// The fully-qualified name of the service account to be used for the build step of deploying this function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#build_service_account GoogleCloudfunctionsFunction#build_service_account}
	BuildServiceAccount *string `field:"optional" json:"buildServiceAccount" yaml:"buildServiceAccount"`
	// Name of the Cloud Build Custom Worker Pool that should be used to build the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#build_worker_pool GoogleCloudfunctionsFunction#build_worker_pool}
	BuildWorkerPool *string `field:"optional" json:"buildWorkerPool" yaml:"buildWorkerPool"`
	// Description of the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#description GoogleCloudfunctionsFunction#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docker Registry to use for storing the function's Docker images. Allowed values are ARTIFACT_REGISTRY (default) and CONTAINER_REGISTRY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#docker_registry GoogleCloudfunctionsFunction#docker_registry}
	DockerRegistry *string `field:"optional" json:"dockerRegistry" yaml:"dockerRegistry"`
	// User managed repository created in Artifact Registry optionally with a customer managed encryption key.
	//
	// If specified, deployments will use Artifact Registry for storing images built with Cloud Build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#docker_repository GoogleCloudfunctionsFunction#docker_repository}
	DockerRepository *string `field:"optional" json:"dockerRepository" yaml:"dockerRepository"`
	// Name of the function that will be executed when the Google Cloud Function is triggered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#entry_point GoogleCloudfunctionsFunction#entry_point}
	EntryPoint *string `field:"optional" json:"entryPoint" yaml:"entryPoint"`
	// A set of key/value environment variable pairs to assign to the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#environment_variables GoogleCloudfunctionsFunction#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// event_trigger block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#event_trigger GoogleCloudfunctionsFunction#event_trigger}
	EventTrigger *GoogleCloudfunctionsFunctionEventTrigger `field:"optional" json:"eventTrigger" yaml:"eventTrigger"`
	// The security level for the function. Defaults to SECURE_OPTIONAL. Valid only if trigger_http is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#https_trigger_security_level GoogleCloudfunctionsFunction#https_trigger_security_level}
	HttpsTriggerSecurityLevel *string `field:"optional" json:"httpsTriggerSecurityLevel" yaml:"httpsTriggerSecurityLevel"`
	// URL which triggers function execution. Returned only if trigger_http is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#https_trigger_url GoogleCloudfunctionsFunction#https_trigger_url}
	HttpsTriggerUrl *string `field:"optional" json:"httpsTriggerUrl" yaml:"httpsTriggerUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#id GoogleCloudfunctionsFunction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// String value that controls what traffic can reach the function.
	//
	// Allowed values are ALLOW_ALL and ALLOW_INTERNAL_ONLY. Changes to this field will recreate the cloud function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#ingress_settings GoogleCloudfunctionsFunction#ingress_settings}
	IngressSettings *string `field:"optional" json:"ingressSettings" yaml:"ingressSettings"`
	// Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#kms_key_name GoogleCloudfunctionsFunction#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#labels GoogleCloudfunctionsFunction#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The limit on the maximum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#max_instances GoogleCloudfunctionsFunction#max_instances}
	MaxInstances *float64 `field:"optional" json:"maxInstances" yaml:"maxInstances"`
	// The limit on the minimum number of function instances that may coexist at a given time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#min_instances GoogleCloudfunctionsFunction#min_instances}
	MinInstances *float64 `field:"optional" json:"minInstances" yaml:"minInstances"`
	// Project of the function. If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#project GoogleCloudfunctionsFunction#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Region of function. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#region GoogleCloudfunctionsFunction#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// secret_environment_variables block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#secret_environment_variables GoogleCloudfunctionsFunction#secret_environment_variables}
	SecretEnvironmentVariables interface{} `field:"optional" json:"secretEnvironmentVariables" yaml:"secretEnvironmentVariables"`
	// secret_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#secret_volumes GoogleCloudfunctionsFunction#secret_volumes}
	SecretVolumes interface{} `field:"optional" json:"secretVolumes" yaml:"secretVolumes"`
	// If provided, the self-provided service account to run the function with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#service_account_email GoogleCloudfunctionsFunction#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// The GCS bucket containing the zip archive which contains the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#source_archive_bucket GoogleCloudfunctionsFunction#source_archive_bucket}
	SourceArchiveBucket *string `field:"optional" json:"sourceArchiveBucket" yaml:"sourceArchiveBucket"`
	// The source archive object (file) in archive bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#source_archive_object GoogleCloudfunctionsFunction#source_archive_object}
	SourceArchiveObject *string `field:"optional" json:"sourceArchiveObject" yaml:"sourceArchiveObject"`
	// source_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#source_repository GoogleCloudfunctionsFunction#source_repository}
	SourceRepository *GoogleCloudfunctionsFunctionSourceRepository `field:"optional" json:"sourceRepository" yaml:"sourceRepository"`
	// Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#timeout GoogleCloudfunctionsFunction#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#timeouts GoogleCloudfunctionsFunction#timeouts}
	Timeouts *GoogleCloudfunctionsFunctionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Boolean variable.
	//
	// Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as https_trigger_url. Cannot be used with trigger_bucket and trigger_topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#trigger_http GoogleCloudfunctionsFunction#trigger_http}
	TriggerHttp interface{} `field:"optional" json:"triggerHttp" yaml:"triggerHttp"`
	// The VPC Network Connector that this cloud function can connect to.
	//
	// It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is projects/* /locations/* /connectors/*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#vpc_connector GoogleCloudfunctionsFunction#vpc_connector}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	VpcConnector *string `field:"optional" json:"vpcConnector" yaml:"vpcConnector"`
	// The egress settings for the connector, controlling what traffic is diverted through it.
	//
	// Allowed values are ALL_TRAFFIC and PRIVATE_RANGES_ONLY. Defaults to PRIVATE_RANGES_ONLY. If unset, this field preserves the previously set value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudfunctions_function#vpc_connector_egress_settings GoogleCloudfunctionsFunction#vpc_connector_egress_settings}
	VpcConnectorEgressSettings *string `field:"optional" json:"vpcConnectorEgressSettings" yaml:"vpcConnectorEgressSettings"`
}

