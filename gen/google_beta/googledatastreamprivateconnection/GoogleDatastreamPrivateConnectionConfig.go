package googledatastreamprivateconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDatastreamPrivateConnectionConfig struct {
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
	// Display name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#display_name GoogleDatastreamPrivateConnection#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The name of the location this private connection is located in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#location GoogleDatastreamPrivateConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The private connectivity identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#private_connection_id GoogleDatastreamPrivateConnection#private_connection_id}
	PrivateConnectionId *string `field:"required" json:"privateConnectionId" yaml:"privateConnectionId"`
	// If set to true, will skip validations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#create_without_validation GoogleDatastreamPrivateConnection#create_without_validation}
	CreateWithoutValidation interface{} `field:"optional" json:"createWithoutValidation" yaml:"createWithoutValidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#id GoogleDatastreamPrivateConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#labels GoogleDatastreamPrivateConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#project GoogleDatastreamPrivateConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// psc_interface_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#psc_interface_config GoogleDatastreamPrivateConnection#psc_interface_config}
	PscInterfaceConfig *GoogleDatastreamPrivateConnectionPscInterfaceConfig `field:"optional" json:"pscInterfaceConfig" yaml:"pscInterfaceConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#timeouts GoogleDatastreamPrivateConnection#timeouts}
	Timeouts *GoogleDatastreamPrivateConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_private_connection#vpc_peering_config GoogleDatastreamPrivateConnection#vpc_peering_config}
	VpcPeeringConfig *GoogleDatastreamPrivateConnectionVpcPeeringConfig `field:"optional" json:"vpcPeeringConfig" yaml:"vpcPeeringConfig"`
}

