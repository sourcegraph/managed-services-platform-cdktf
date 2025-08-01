package googlemonitoringservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleMonitoringServiceConfig struct {
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
	// An optional service ID to use. If not given, the server will generate a service ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_service#service_id GoogleMonitoringService#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// basic_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_service#basic_service GoogleMonitoringService#basic_service}
	BasicService *GoogleMonitoringServiceBasicService `field:"optional" json:"basicService" yaml:"basicService"`
	// Name used for UI elements listing this Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_service#display_name GoogleMonitoringService#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_service#id GoogleMonitoringService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_service#project GoogleMonitoringService#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_service#timeouts GoogleMonitoringService#timeouts}
	Timeouts *GoogleMonitoringServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Labels which have been used to annotate the service.
	//
	// Label keys must start
	// with a letter. Label keys and values may contain lowercase letters,
	// numbers, underscores, and dashes. Label keys and values have a maximum
	// length of 63 characters, and must be less than 128 bytes in size. Up to 64
	// label entries may be stored. For labels which do not have a semantic value,
	// the empty string may be supplied for the label value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_monitoring_service#user_labels GoogleMonitoringService#user_labels}
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
}

