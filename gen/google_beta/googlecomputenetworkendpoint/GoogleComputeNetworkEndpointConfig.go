package googlecomputenetworkendpoint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeNetworkEndpointConfig struct {
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
	// IPv4 address of network endpoint.
	//
	// The IP address must belong
	// to a VM in GCE (either the primary IP or as part of an aliased IP
	// range).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#ip_address GoogleComputeNetworkEndpoint#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The network endpoint group this endpoint is part of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#network_endpoint_group GoogleComputeNetworkEndpoint#network_endpoint_group}
	NetworkEndpointGroup *string `field:"required" json:"networkEndpointGroup" yaml:"networkEndpointGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#id GoogleComputeNetworkEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name for a specific VM instance that the IP address belongs to.
	//
	// This is required for network endpoints of type GCE_VM_IP_PORT.
	// The instance must be in the same zone of network endpoint group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#instance GoogleComputeNetworkEndpoint#instance}
	Instance *string `field:"optional" json:"instance" yaml:"instance"`
	// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type of 'GCE_VM_IP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#port GoogleComputeNetworkEndpoint#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#project GoogleComputeNetworkEndpoint#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#timeouts GoogleComputeNetworkEndpoint#timeouts}
	Timeouts *GoogleComputeNetworkEndpointTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Zone where the containing network endpoint group is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_network_endpoint#zone GoogleComputeNetworkEndpoint#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

