package alertpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#message AlertPolicy#message}.
	Message *string `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#name AlertPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#actions AlertPolicy#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#alert_description AlertPolicy#alert_description}.
	AlertDescription *string `field:"optional" json:"alertDescription" yaml:"alertDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#alias AlertPolicy#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#continue_policy AlertPolicy#continue_policy}.
	ContinuePolicy interface{} `field:"optional" json:"continuePolicy" yaml:"continuePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#enabled AlertPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#entity AlertPolicy#entity}.
	Entity *string `field:"optional" json:"entity" yaml:"entity"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#filter AlertPolicy#filter}
	Filter *AlertPolicyFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#id AlertPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#ignore_original_actions AlertPolicy#ignore_original_actions}.
	IgnoreOriginalActions interface{} `field:"optional" json:"ignoreOriginalActions" yaml:"ignoreOriginalActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#ignore_original_details AlertPolicy#ignore_original_details}.
	IgnoreOriginalDetails interface{} `field:"optional" json:"ignoreOriginalDetails" yaml:"ignoreOriginalDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#ignore_original_responders AlertPolicy#ignore_original_responders}.
	IgnoreOriginalResponders interface{} `field:"optional" json:"ignoreOriginalResponders" yaml:"ignoreOriginalResponders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#ignore_original_tags AlertPolicy#ignore_original_tags}.
	IgnoreOriginalTags interface{} `field:"optional" json:"ignoreOriginalTags" yaml:"ignoreOriginalTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#policy_description AlertPolicy#policy_description}.
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#priority AlertPolicy#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// responders block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#responders AlertPolicy#responders}
	Responders interface{} `field:"optional" json:"responders" yaml:"responders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#source AlertPolicy#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#tags AlertPolicy#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#team_id AlertPolicy#team_id}.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// time_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/alert_policy#time_restriction AlertPolicy#time_restriction}
	TimeRestriction *AlertPolicyTimeRestriction `field:"optional" json:"timeRestriction" yaml:"timeRestriction"`
}

