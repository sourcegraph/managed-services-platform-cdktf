package googlecomputeaddress

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeAddressConfig struct {
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
	// Name of the resource.
	//
	// The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#name GoogleComputeAddress#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The static external IP address represented by this resource.
	//
	// The IP address must be inside the specified subnetwork,
	// if any. Set by the API if undefined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#address GoogleComputeAddress#address}
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The type of address to reserve.
	//
	// Note: if you set this argument's value as 'INTERNAL' you need to leave the 'network_tier' argument unset in that resource block. Default value: "EXTERNAL" Possible values: ["INTERNAL", "EXTERNAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#address_type GoogleComputeAddress#address_type}
	AddressType *string `field:"optional" json:"addressType" yaml:"addressType"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#description GoogleComputeAddress#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#id GoogleComputeAddress#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The endpoint type of this address, which should be VM or NETLB.
	//
	// This is
	// used for deciding which type of endpoint this address can be used after
	// the external IPv6 address reservation. Possible values: ["VM", "NETLB"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#ipv6_endpoint_type GoogleComputeAddress#ipv6_endpoint_type}
	Ipv6EndpointType *string `field:"optional" json:"ipv6EndpointType" yaml:"ipv6EndpointType"`
	// The IP Version that will be used by this address. The default value is 'IPV4'. Possible values: ["IPV4", "IPV6"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#ip_version GoogleComputeAddress#ip_version}
	IpVersion *string `field:"optional" json:"ipVersion" yaml:"ipVersion"`
	// Labels to apply to this address.  A list of key->value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#labels GoogleComputeAddress#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The URL of the network in which to reserve the address.
	//
	// This field
	// can only be used with INTERNAL type with the VPC_PEERING and
	// IPSEC_INTERCONNECT purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#network GoogleComputeAddress#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// The networking tier used for configuring this address.
	//
	// If this field is not
	// specified, it is assumed to be PREMIUM.
	// This argument should not be used when configuring Internal addresses, because [network tier cannot be set for internal traffic; it's always Premium](https://cloud.google.com/network-tiers/docs/overview). Possible values: ["PREMIUM", "STANDARD"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#network_tier GoogleComputeAddress#network_tier}
	NetworkTier *string `field:"optional" json:"networkTier" yaml:"networkTier"`
	// The prefix length if the resource represents an IP range.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#prefix_length GoogleComputeAddress#prefix_length}
	PrefixLength *float64 `field:"optional" json:"prefixLength" yaml:"prefixLength"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#project GoogleComputeAddress#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The purpose of this resource, which can be one of the following values.
	//
	// * GCE_ENDPOINT for addresses that are used by VM instances, alias IP
	// ranges, load balancers, and similar resources.
	//
	// * SHARED_LOADBALANCER_VIP for an address that can be used by multiple
	// internal load balancers.
	//
	// * VPC_PEERING for addresses that are reserved for VPC peer networks.
	//
	// * IPSEC_INTERCONNECT for addresses created from a private IP range that
	// are reserved for a VLAN attachment in an HA VPN over Cloud Interconnect
	// configuration. These addresses are regional resources.
	//
	// * PRIVATE_SERVICE_CONNECT for a private network address that is used to
	// configure Private Service Connect. Only global internal addresses can use
	// this purpose.
	//
	// This should only be set when using an Internal address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#purpose GoogleComputeAddress#purpose}
	Purpose *string `field:"optional" json:"purpose" yaml:"purpose"`
	// The Region in which the created address should reside. If it is not provided, the provider region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#region GoogleComputeAddress#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The URL of the subnetwork in which to reserve the address.
	//
	// If an IP
	// address is specified, it must be within the subnetwork's IP range.
	// This field can only be used with INTERNAL type with
	// GCE_ENDPOINT/DNS_RESOLVER purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#subnetwork GoogleComputeAddress#subnetwork}
	Subnetwork *string `field:"optional" json:"subnetwork" yaml:"subnetwork"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_address#timeouts GoogleComputeAddress#timeouts}
	Timeouts *GoogleComputeAddressTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

