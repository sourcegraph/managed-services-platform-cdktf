package googleeventarcenrollment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleEventarcEnrollmentConfig struct {
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
	// A CEL expression identifying which messages this enrollment applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#cel_match GoogleEventarcEnrollment#cel_match}
	CelMatch *string `field:"required" json:"celMatch" yaml:"celMatch"`
	// Destination is the Pipeline that the Enrollment is delivering to.
	//
	// It must
	// point to the full resource name of a Pipeline. Format:
	// "projects/{PROJECT_ID}/locations/{region}/pipelines/{PIPELINE_ID)"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#destination GoogleEventarcEnrollment#destination}
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The user-provided ID to be assigned to the Enrollment. It should match the format '^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#enrollment_id GoogleEventarcEnrollment#enrollment_id}
	EnrollmentId *string `field:"required" json:"enrollmentId" yaml:"enrollmentId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#location GoogleEventarcEnrollment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Resource name of the message bus identifying the source of the messages. It matches the form projects/{project}/locations/{location}/messageBuses/{messageBus}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#message_bus GoogleEventarcEnrollment#message_bus}
	MessageBus *string `field:"required" json:"messageBus" yaml:"messageBus"`
	// Resource annotations.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#annotations GoogleEventarcEnrollment#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Resource display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#display_name GoogleEventarcEnrollment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#id GoogleEventarcEnrollment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#labels GoogleEventarcEnrollment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#project GoogleEventarcEnrollment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_enrollment#timeouts GoogleEventarcEnrollment#timeouts}
	Timeouts *GoogleEventarcEnrollmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

