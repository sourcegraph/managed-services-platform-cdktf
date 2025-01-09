package googlenetworksecuritymirroringdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetworkSecurityMirroringDeploymentConfig struct {
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
	// Required. Immutable. The regional load balancer which the mirrored traffic should be forwarded to. Format is: projects/{project}/regions/{region}/forwardingRules/{forwardingRule}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#forwarding_rule GoogleNetworkSecurityMirroringDeployment#forwarding_rule}
	ForwardingRule *string `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// Resource ID segment making up resource 'name'.
	//
	// It identifies the resource within its parent collection as described in https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringDeployment'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#location GoogleNetworkSecurityMirroringDeployment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Required. Immutable. The Mirroring Deployment Group that this resource is part of. Format is: 'projects/{project}/locations/global/mirroringDeploymentGroups/{mirroringDeploymentGroup}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#mirroring_deployment_group GoogleNetworkSecurityMirroringDeployment#mirroring_deployment_group}
	MirroringDeploymentGroup *string `field:"required" json:"mirroringDeploymentGroup" yaml:"mirroringDeploymentGroup"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and mirroring_deployment_id from the method_signature of Create RPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#mirroring_deployment_id GoogleNetworkSecurityMirroringDeployment#mirroring_deployment_id}
	MirroringDeploymentId *string `field:"required" json:"mirroringDeploymentId" yaml:"mirroringDeploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#id GoogleNetworkSecurityMirroringDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#labels GoogleNetworkSecurityMirroringDeployment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#project GoogleNetworkSecurityMirroringDeployment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_network_security_mirroring_deployment#timeouts GoogleNetworkSecurityMirroringDeployment#timeouts}
	Timeouts *GoogleNetworkSecurityMirroringDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

