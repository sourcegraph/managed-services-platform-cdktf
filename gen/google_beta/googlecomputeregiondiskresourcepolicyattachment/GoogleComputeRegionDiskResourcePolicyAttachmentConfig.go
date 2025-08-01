package googlecomputeregiondiskresourcepolicyattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeRegionDiskResourcePolicyAttachmentConfig struct {
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
	// The name of the regional disk in which the resource policies are attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_disk_resource_policy_attachment#disk GoogleComputeRegionDiskResourcePolicyAttachment#disk}
	Disk *string `field:"required" json:"disk" yaml:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_disk_resource_policy_attachment#name GoogleComputeRegionDiskResourcePolicyAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_disk_resource_policy_attachment#id GoogleComputeRegionDiskResourcePolicyAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_disk_resource_policy_attachment#project GoogleComputeRegionDiskResourcePolicyAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// A reference to the region where the disk resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_disk_resource_policy_attachment#region GoogleComputeRegionDiskResourcePolicyAttachment#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_region_disk_resource_policy_attachment#timeouts GoogleComputeRegionDiskResourcePolicyAttachment#timeouts}
	Timeouts *GoogleComputeRegionDiskResourcePolicyAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

