package googleintegrationconnectorsendpointattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIntegrationConnectorsEndpointAttachmentConfig struct {
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
	// Location in which Endpoint Attachment needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#location GoogleIntegrationConnectorsEndpointAttachment#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Name of Endpoint Attachment needs to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#name GoogleIntegrationConnectorsEndpointAttachment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The path of the service attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#service_attachment GoogleIntegrationConnectorsEndpointAttachment#service_attachment}
	ServiceAttachment *string `field:"required" json:"serviceAttachment" yaml:"serviceAttachment"`
	// Description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#description GoogleIntegrationConnectorsEndpointAttachment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Enable global access for endpoint attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#endpoint_global_access GoogleIntegrationConnectorsEndpointAttachment#endpoint_global_access}
	EndpointGlobalAccess interface{} `field:"optional" json:"endpointGlobalAccess" yaml:"endpointGlobalAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#id GoogleIntegrationConnectorsEndpointAttachment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Resource labels to represent user provided metadata.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#labels GoogleIntegrationConnectorsEndpointAttachment#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#project GoogleIntegrationConnectorsEndpointAttachment#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integration_connectors_endpoint_attachment#timeouts GoogleIntegrationConnectorsEndpointAttachment#timeouts}
	Timeouts *GoogleIntegrationConnectorsEndpointAttachmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

