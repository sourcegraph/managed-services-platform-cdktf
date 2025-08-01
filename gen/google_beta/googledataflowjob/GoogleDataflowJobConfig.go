package googledataflowjob

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataflowJobConfig struct {
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
	// A unique name for the resource, required by Dataflow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#name GoogleDataflowJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A writeable location on Google Cloud Storage for the Dataflow job to dump its temporary data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#temp_gcs_location GoogleDataflowJob#temp_gcs_location}
	TempGcsLocation *string `field:"required" json:"tempGcsLocation" yaml:"tempGcsLocation"`
	// The Google Cloud Storage path to the Dataflow job template.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#template_gcs_path GoogleDataflowJob#template_gcs_path}
	TemplateGcsPath *string `field:"required" json:"templateGcsPath" yaml:"templateGcsPath"`
	// List of experiments that should be used by the job. An example value is ["enable_stackdriver_agent_metrics"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#additional_experiments GoogleDataflowJob#additional_experiments}
	AdditionalExperiments *[]*string `field:"optional" json:"additionalExperiments" yaml:"additionalExperiments"`
	// Indicates if the job should use the streaming engine feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#enable_streaming_engine GoogleDataflowJob#enable_streaming_engine}
	EnableStreamingEngine interface{} `field:"optional" json:"enableStreamingEngine" yaml:"enableStreamingEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#id GoogleDataflowJob#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The configuration for VM IPs. Options are "WORKER_IP_PUBLIC" or "WORKER_IP_PRIVATE".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#ip_configuration GoogleDataflowJob#ip_configuration}
	IpConfiguration *string `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// The name for the Cloud KMS key for the job. Key format is: projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#kms_key_name GoogleDataflowJob#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
	// User labels to be specified for the job.
	//
	// Keys and values should follow the restrictions specified in the labeling restrictions page. NOTE: This field is non-authoritative, and will only manage the labels present in your configuration.
	// 				Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#labels GoogleDataflowJob#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The machine type to use for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#machine_type GoogleDataflowJob#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// The number of workers permitted to work on the job. More workers may improve processing speed at additional cost.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#max_workers GoogleDataflowJob#max_workers}
	MaxWorkers *float64 `field:"optional" json:"maxWorkers" yaml:"maxWorkers"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#network GoogleDataflowJob#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// One of "drain" or "cancel". Specifies behavior of deletion during terraform destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#on_delete GoogleDataflowJob#on_delete}
	OnDelete *string `field:"optional" json:"onDelete" yaml:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#parameters GoogleDataflowJob#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The project in which the resource belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#project GoogleDataflowJob#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region in which the created job should run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#region GoogleDataflowJob#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The Service Account email used to create the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#service_account_email GoogleDataflowJob#service_account_email}
	ServiceAccountEmail *string `field:"optional" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// If true, treat DRAINING and CANCELLING as terminal job states and do not wait for further changes before removing from terraform state and moving on.
	//
	// WARNING: this will lead to job name conflicts if you do not ensure that the job names are different, e.g. by embedding a release ID or by using a random_id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#skip_wait_on_job_termination GoogleDataflowJob#skip_wait_on_job_termination}
	SkipWaitOnJobTermination interface{} `field:"optional" json:"skipWaitOnJobTermination" yaml:"skipWaitOnJobTermination"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#subnetwork GoogleDataflowJob#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#timeouts GoogleDataflowJob#timeouts}
	Timeouts *GoogleDataflowJobTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Only applicable when updating a pipeline.
	//
	// Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#transform_name_mapping GoogleDataflowJob#transform_name_mapping}
	TransformNameMapping *map[string]*string `field:"optional" json:"transformNameMapping" yaml:"transformNameMapping"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataflow_job#zone GoogleDataflowJob#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

