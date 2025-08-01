package googlecloudrunv2job

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudRunV2JobConfig struct {
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
	// The location of the cloud run job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#location GoogleCloudRunV2Job#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of the Job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#name GoogleCloudRunV2Job#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#template GoogleCloudRunV2Job#template}
	Template *GoogleCloudRunV2JobTemplate `field:"required" json:"template" yaml:"template"`
	// Unstructured key value map that may be set by external tools to store and arbitrary metadata.
	//
	// They are not queryable and should be preserved when modifying objects.
	//
	// Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected on new resources.
	// All system annotations in v1 now have a corresponding field in v2 Job.
	//
	// This field follows Kubernetes annotations' namespacing, limits, and rules.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#annotations GoogleCloudRunV2Job#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#binary_authorization GoogleCloudRunV2Job#binary_authorization}
	BinaryAuthorization *GoogleCloudRunV2JobBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// Arbitrary identifier for the API client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#client GoogleCloudRunV2Job#client}
	Client *string `field:"optional" json:"client" yaml:"client"`
	// Arbitrary version identifier for the API client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#client_version GoogleCloudRunV2Job#client_version}
	ClientVersion *string `field:"optional" json:"clientVersion" yaml:"clientVersion"`
	// Whether Terraform will be prevented from destroying the job.
	//
	// Defaults to true.
	// When a'terraform destroy' or 'terraform apply' would delete the job,
	// the command will fail if this field is not set to false in Terraform state.
	// When the field is set to true or unset in Terraform state, a 'terraform apply'
	// or 'terraform destroy' that would delete the job will fail.
	// When the field is set to false, deleting the job is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#deletion_protection GoogleCloudRunV2Job#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#id GoogleCloudRunV2Job#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Unstructured key value map that can be used to organize and categorize objects.
	//
	// User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component,
	// environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels.
	//
	// Cloud Run API v2 does not support labels with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.
	// All system labels in v1 now have a corresponding field in v2 Job.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#labels GoogleCloudRunV2Job#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.
	//
	// For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: ["UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#launch_stage GoogleCloudRunV2Job#launch_stage}
	LaunchStage *string `field:"optional" json:"launchStage" yaml:"launchStage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#project GoogleCloudRunV2Job#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A unique string used as a suffix creating a new execution upon job create or update.
	//
	// The Job will become ready when the execution is successfully completed.
	// The sum of job name and token length must be fewer than 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#run_execution_token GoogleCloudRunV2Job#run_execution_token}
	RunExecutionToken *string `field:"optional" json:"runExecutionToken" yaml:"runExecutionToken"`
	// A unique string used as a suffix creating a new execution upon job create or update.
	//
	// The Job will become ready when the execution is successfully started.
	// The sum of job name and token length must be fewer than 63 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#start_execution_token GoogleCloudRunV2Job#start_execution_token}
	StartExecutionToken *string `field:"optional" json:"startExecutionToken" yaml:"startExecutionToken"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_job#timeouts GoogleCloudRunV2Job#timeouts}
	Timeouts *GoogleCloudRunV2JobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

