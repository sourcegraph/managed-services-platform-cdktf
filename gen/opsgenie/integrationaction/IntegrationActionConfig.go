package integrationaction

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationActionConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#integration_id IntegrationAction#integration_id}.
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// acknowledge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#acknowledge IntegrationAction#acknowledge}
	Acknowledge any `field:"optional" json:"acknowledge" yaml:"acknowledge"`
	// add_note block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#add_note IntegrationAction#add_note}
	AddNote any `field:"optional" json:"addNote" yaml:"addNote"`
	// close block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#close IntegrationAction#close}
	Close any `field:"optional" json:"close" yaml:"close"`
	// create block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#create IntegrationAction#create}
	Create any `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#id IntegrationAction#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ignore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/integration_action#ignore IntegrationAction#ignore}
	Ignore any `field:"optional" json:"ignore" yaml:"ignore"`
}
