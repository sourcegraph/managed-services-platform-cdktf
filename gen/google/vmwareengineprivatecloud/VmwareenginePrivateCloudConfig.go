package vmwareengineprivatecloud

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VmwareenginePrivateCloudConfig struct {
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
	// The location where the PrivateCloud should reside.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#location VmwareenginePrivateCloud#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// management_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#management_cluster VmwareenginePrivateCloud#management_cluster}
	ManagementCluster *VmwareenginePrivateCloudManagementCluster `field:"required" json:"managementCluster" yaml:"managementCluster"`
	// The ID of the PrivateCloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#name VmwareenginePrivateCloud#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#network_config VmwareenginePrivateCloud#network_config}
	NetworkConfig *VmwareenginePrivateCloudNetworkConfig `field:"required" json:"networkConfig" yaml:"networkConfig"`
	// The number of hours to delay this request.
	//
	// You can set this value to an hour between 0 to 8, where setting it to 0 starts the deletion request immediately. If no value is set, a default value is set at the API Level.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#deletion_delay_hours VmwareenginePrivateCloud#deletion_delay_hours}
	DeletionDelayHours *float64 `field:"optional" json:"deletionDelayHours" yaml:"deletionDelayHours"`
	// User-provided description for this private cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#description VmwareenginePrivateCloud#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#id VmwareenginePrivateCloud#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#project VmwareenginePrivateCloud#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// While set true, deletion_delay_hours value will be sent in the request even for zero value of the field.
	//
	// This field is only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with deletion_delay_hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#send_deletion_delay_hours_if_zero VmwareenginePrivateCloud#send_deletion_delay_hours_if_zero}
	SendDeletionDelayHoursIfZero interface{} `field:"optional" json:"sendDeletionDelayHoursIfZero" yaml:"sendDeletionDelayHoursIfZero"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#timeouts VmwareenginePrivateCloud#timeouts}
	Timeouts *VmwareenginePrivateCloudTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#type VmwareenginePrivateCloud#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

