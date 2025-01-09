package googlenetworksecurityinterceptendpointgroupassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityInterceptEndpointGroupAssociationConfig struct {
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
	// Immutable. The Intercept Endpoint Group that this resource is connected to. Format is: 'projects/{project}/locations/global/interceptEndpointGroups/{interceptEndpointGroup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#intercept_endpoint_group GoogleNetworkSecurityInterceptEndpointGroupAssociation#intercept_endpoint_group}
	InterceptEndpointGroup *string `field:"required" json:"interceptEndpointGroup" yaml:"interceptEndpointGroup"`
	// The location of the Intercept Endpoint Group Association, currently restricted to 'global'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#location GoogleNetworkSecurityInterceptEndpointGroupAssociation#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Immutable. The VPC network associated. Format: 'projects/{project}/global/networks/{network}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#network GoogleNetworkSecurityInterceptEndpointGroupAssociation#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#id GoogleNetworkSecurityInterceptEndpointGroupAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ID of the Intercept Endpoint Group Association.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#intercept_endpoint_group_association_id GoogleNetworkSecurityInterceptEndpointGroupAssociation#intercept_endpoint_group_association_id}
	InterceptEndpointGroupAssociationId *string `field:"optional" json:"interceptEndpointGroupAssociationId" yaml:"interceptEndpointGroupAssociationId"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#labels GoogleNetworkSecurityInterceptEndpointGroupAssociation#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#project GoogleNetworkSecurityInterceptEndpointGroupAssociation#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_intercept_endpoint_group_association#timeouts GoogleNetworkSecurityInterceptEndpointGroupAssociation#timeouts}
	Timeouts *GoogleNetworkSecurityInterceptEndpointGroupAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

