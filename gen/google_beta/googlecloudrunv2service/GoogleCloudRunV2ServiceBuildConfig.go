package googlecloudrunv2service


type GoogleCloudRunV2ServiceBuildConfig struct {
	// The base image used to build the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#base_image GoogleCloudRunV2Service#base_image}
	BaseImage *string `field:"optional" json:"baseImage" yaml:"baseImage"`
	// Sets whether the function will receive automatic base image updates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#enable_automatic_updates GoogleCloudRunV2Service#enable_automatic_updates}
	EnableAutomaticUpdates interface{} `field:"optional" json:"enableAutomaticUpdates" yaml:"enableAutomaticUpdates"`
	// User-provided build-time environment variables for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#environment_variables GoogleCloudRunV2Service#environment_variables}
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The name of the function (as defined in source code) that will be executed.
	//
	// Defaults to the resource name suffix, if not specified. For backward compatibility, if function with given name is not found, then the system will try to use function named "function".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#function_target GoogleCloudRunV2Service#function_target}
	FunctionTarget *string `field:"optional" json:"functionTarget" yaml:"functionTarget"`
	// Artifact Registry URI to store the built image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#image_uri GoogleCloudRunV2Service#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// Service account to be used for building the container. The format of this field is 'projects/{projectId}/serviceAccounts/{serviceAccountEmail}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#service_account GoogleCloudRunV2Service#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// The Cloud Storage bucket URI where the function source code is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#source_location GoogleCloudRunV2Service#source_location}
	SourceLocation *string `field:"optional" json:"sourceLocation" yaml:"sourceLocation"`
	// Name of the Cloud Build Custom Worker Pool that should be used to build the Cloud Run function.
	//
	// The format of this field is 'projects/{project}/locations/{region}/workerPools/{workerPool}' where {project} and {region} are the project id and region respectively where the worker pool is defined and {workerPool} is the short name of the worker pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_service#worker_pool GoogleCloudRunV2Service#worker_pool}
	WorkerPool *string `field:"optional" json:"workerPool" yaml:"workerPool"`
}

