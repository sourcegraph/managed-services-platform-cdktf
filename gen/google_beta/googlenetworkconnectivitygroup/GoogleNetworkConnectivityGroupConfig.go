package googlenetworkconnectivitygroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkConnectivityGroupConfig struct {
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
	// The name of the hub. Hub names must be unique. They use the following form: projects/{projectNumber}/locations/global/hubs/{hubId}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#hub GoogleNetworkConnectivityGroup#hub}
	Hub *string `field:"required" json:"hub" yaml:"hub"`
	// The name of the group. Group names must be unique. Possible values: ["default", "center", "edge"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#name GoogleNetworkConnectivityGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// auto_accept block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#auto_accept GoogleNetworkConnectivityGroup#auto_accept}
	AutoAccept *GoogleNetworkConnectivityGroupAutoAccept `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// An optional description of the group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#description GoogleNetworkConnectivityGroup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#id GoogleNetworkConnectivityGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#labels GoogleNetworkConnectivityGroup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#project GoogleNetworkConnectivityGroup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_group#timeouts GoogleNetworkConnectivityGroup#timeouts}
	Timeouts *GoogleNetworkConnectivityGroupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

