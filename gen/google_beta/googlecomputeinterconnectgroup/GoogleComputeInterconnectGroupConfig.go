package googlecomputeinterconnectgroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeInterconnectGroupConfig struct {
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
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_interconnect_group#intent GoogleComputeInterconnectGroup#intent}
	Intent *GoogleComputeInterconnectGroupIntent `field:"required" json:"intent" yaml:"intent"`
	// Name of the resource.
	//
	// Provided by the client when the resource is created. The name must be
	// 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
	// character must be a lowercase letter, and all following characters must be a dash,
	// lowercase letter, or digit, except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_interconnect_group#name GoogleComputeInterconnectGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// An optional description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_interconnect_group#description GoogleComputeInterconnectGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_interconnect_group#id GoogleComputeInterconnectGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// interconnects block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_interconnect_group#interconnects GoogleComputeInterconnectGroup#interconnects}
	Interconnects interface{} `field:"optional" json:"interconnects" yaml:"interconnects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_interconnect_group#project GoogleComputeInterconnectGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_interconnect_group#timeouts GoogleComputeInterconnectGroup#timeouts}
	Timeouts *GoogleComputeInterconnectGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

