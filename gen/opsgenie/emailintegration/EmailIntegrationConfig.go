package emailintegration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmailIntegrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#email_username EmailIntegration#email_username}.
	EmailUsername *string `field:"required" json:"emailUsername" yaml:"emailUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#name EmailIntegration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#enabled EmailIntegration#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#id EmailIntegration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#ignore_responders_from_payload EmailIntegration#ignore_responders_from_payload}.
	IgnoreRespondersFromPayload interface{} `field:"optional" json:"ignoreRespondersFromPayload" yaml:"ignoreRespondersFromPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#owner_team_id EmailIntegration#owner_team_id}.
	OwnerTeamId *string `field:"optional" json:"ownerTeamId" yaml:"ownerTeamId"`
	// responders block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#responders EmailIntegration#responders}
	Responders interface{} `field:"optional" json:"responders" yaml:"responders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/email_integration#suppress_notifications EmailIntegration#suppress_notifications}.
	SuppressNotifications interface{} `field:"optional" json:"suppressNotifications" yaml:"suppressNotifications"`
}

