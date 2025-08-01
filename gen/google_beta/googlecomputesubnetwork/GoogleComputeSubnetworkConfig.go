package googlecomputesubnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeSubnetworkConfig struct {
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
	// The name of the resource, provided by the client when initially creating the resource.
	//
	// The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which
	// means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#name GoogleComputeSubnetwork#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The network this subnet belongs to. Only networks that are in the distributed mode can have subnetworks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#network GoogleComputeSubnetwork#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Typically packets destined to IPs within the subnetwork range that do not match existing resources are dropped and prevented from leaving the VPC.
	//
	// Setting this field to true will allow these packets to match dynamic routes injected
	// via BGP even if their destinations match existing subnet ranges.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#allow_subnet_cidr_routes_overlap GoogleComputeSubnetwork#allow_subnet_cidr_routes_overlap}
	AllowSubnetCidrRoutesOverlap interface{} `field:"optional" json:"allowSubnetCidrRoutesOverlap" yaml:"allowSubnetCidrRoutesOverlap"`
	// An optional description of this resource.
	//
	// Provide this property when
	// you create the resource. This field can be set only at resource
	// creation time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#description GoogleComputeSubnetwork#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether to enable flow logging for this subnetwork.
	//
	// If this field is not explicitly set,
	// it will not appear in get listings. If not set the default behavior is determined by the
	// org policy, if there is no org policy specified, then it will default to disabled.
	// This field isn't supported if the subnet purpose field is set to REGIONAL_MANAGED_PROXY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#enable_flow_logs GoogleComputeSubnetwork#enable_flow_logs}
	EnableFlowLogs interface{} `field:"optional" json:"enableFlowLogs" yaml:"enableFlowLogs"`
	// The range of external IPv6 addresses that are owned by this subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#external_ipv6_prefix GoogleComputeSubnetwork#external_ipv6_prefix}
	ExternalIpv6Prefix *string `field:"optional" json:"externalIpv6Prefix" yaml:"externalIpv6Prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#id GoogleComputeSubnetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The range of internal addresses that are owned by this subnetwork.
	//
	// Provide this property when you create the subnetwork. For example,
	// 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
	// non-overlapping within a network. Only IPv4 is supported.
	// Field is optional when 'reserved_internal_range' is defined, otherwise required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#ip_cidr_range GoogleComputeSubnetwork#ip_cidr_range}
	IpCidrRange *string `field:"optional" json:"ipCidrRange" yaml:"ipCidrRange"`
	// Resource reference of a PublicDelegatedPrefix.
	//
	// The PDP must be a sub-PDP
	// in EXTERNAL_IPV6_SUBNETWORK_CREATION mode.
	// Use one of the following formats to specify a sub-PDP when creating an
	// IPv6 NetLB forwarding rule using BYOIP:
	// Full resource URL, as in:
	//   * 'https://www.googleapis.com/compute/v1/projects/{{projectId}}/regions/{{region}}/publicDelegatedPrefixes/{{sub-pdp-name}}'
	// Partial URL, as in:
	//   * 'projects/{{projectId}}/regions/region/publicDelegatedPrefixes/{{sub-pdp-name}}'
	//   * 'regions/{{region}}/publicDelegatedPrefixes/{{sub-pdp-name}}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#ip_collection GoogleComputeSubnetwork#ip_collection}
	IpCollection *string `field:"optional" json:"ipCollection" yaml:"ipCollection"`
	// The access type of IPv6 address this subnet holds.
	//
	// It's immutable and can only be specified during creation
	// or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
	// cannot enable direct path. Possible values: ["EXTERNAL", "INTERNAL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#ipv6_access_type GoogleComputeSubnetwork#ipv6_access_type}
	Ipv6AccessType *string `field:"optional" json:"ipv6AccessType" yaml:"ipv6AccessType"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#log_config GoogleComputeSubnetwork#log_config}
	LogConfig *GoogleComputeSubnetworkLogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
	// params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#params GoogleComputeSubnetwork#params}
	Params *GoogleComputeSubnetworkParams `field:"optional" json:"params" yaml:"params"`
	// When enabled, VMs in this subnetwork without external IP addresses can access Google APIs and services by using Private Google Access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#private_ip_google_access GoogleComputeSubnetwork#private_ip_google_access}
	PrivateIpGoogleAccess interface{} `field:"optional" json:"privateIpGoogleAccess" yaml:"privateIpGoogleAccess"`
	// The private IPv6 google access type for the VMs in this subnet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#private_ipv6_google_access GoogleComputeSubnetwork#private_ipv6_google_access}
	PrivateIpv6GoogleAccess *string `field:"optional" json:"privateIpv6GoogleAccess" yaml:"privateIpv6GoogleAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#project GoogleComputeSubnetwork#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The purpose of the resource.
	//
	// This field can be either 'PRIVATE', 'REGIONAL_MANAGED_PROXY', 'GLOBAL_MANAGED_PROXY', 'PRIVATE_SERVICE_CONNECT', 'PEER_MIGRATION' or 'PRIVATE_NAT'([Beta](https://terraform.io/docs/providers/google/guides/provider_versions.html)).
	// A subnet with purpose set to 'REGIONAL_MANAGED_PROXY' is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
	// A subnetwork in a given region with purpose set to 'GLOBAL_MANAGED_PROXY' is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
	// A subnetwork with purpose set to 'PRIVATE_SERVICE_CONNECT' reserves the subnet for hosting a Private Service Connect published service.
	// A subnetwork with purpose set to 'PEER_MIGRATION' is a user created subnetwork that is reserved for migrating resources from one peered network to another.
	// A subnetwork with purpose set to 'PRIVATE_NAT' is used as source range for Private NAT gateways.
	// Note that 'REGIONAL_MANAGED_PROXY' is the preferred setting for all regional Envoy load balancers.
	// If unspecified, the purpose defaults to 'PRIVATE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#purpose GoogleComputeSubnetwork#purpose}
	Purpose *string `field:"optional" json:"purpose" yaml:"purpose"`
	// The GCP region for this subnetwork.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#region GoogleComputeSubnetwork#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The ID of the reserved internal range. Must be prefixed with 'networkconnectivity.googleapis.com' E.g. 'networkconnectivity.googleapis.com/projects/{project}/locations/global/internalRanges/{rangeId}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#reserved_internal_range GoogleComputeSubnetwork#reserved_internal_range}
	ReservedInternalRange *string `field:"optional" json:"reservedInternalRange" yaml:"reservedInternalRange"`
	// The role of subnetwork.
	//
	// Currently, this field is only used when 'purpose' is 'REGIONAL_MANAGED_PROXY'.
	// The value can be set to 'ACTIVE' or 'BACKUP'.
	// An 'ACTIVE' subnetwork is one that is currently being used for Envoy-based load balancers in a region.
	// A 'BACKUP' subnetwork is one that is ready to be promoted to 'ACTIVE' or is currently draining. Possible values: ["ACTIVE", "BACKUP"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#role GoogleComputeSubnetwork#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
	// secondary_ip_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#secondary_ip_range GoogleComputeSubnetwork#secondary_ip_range}
	SecondaryIpRange interface{} `field:"optional" json:"secondaryIpRange" yaml:"secondaryIpRange"`
	// Controls the removal behavior of secondary_ip_range.
	//
	// When false, removing secondary_ip_range from config will not produce a diff as
	// the provider will default to the API's value.
	// When true, the provider will treat removing secondary_ip_range as sending an
	// empty list of secondary IP ranges to the API.
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#send_secondary_ip_range_if_empty GoogleComputeSubnetwork#send_secondary_ip_range_if_empty}
	SendSecondaryIpRangeIfEmpty interface{} `field:"optional" json:"sendSecondaryIpRangeIfEmpty" yaml:"sendSecondaryIpRangeIfEmpty"`
	// The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
	//
	// If not specified IPV4_ONLY will be used. Possible values: ["IPV4_ONLY", "IPV4_IPV6", "IPV6_ONLY"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#stack_type GoogleComputeSubnetwork#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_subnetwork#timeouts GoogleComputeSubnetwork#timeouts}
	Timeouts *GoogleComputeSubnetworkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

