package notificationpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationPolicyConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#filter NotificationPolicy#filter}
	Filter *NotificationPolicyFilter `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#name NotificationPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#team_id NotificationPolicy#team_id}.
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// auto_close_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#auto_close_action NotificationPolicy#auto_close_action}
	AutoCloseAction *NotificationPolicyAutoCloseAction `field:"optional" json:"autoCloseAction" yaml:"autoCloseAction"`
	// auto_restart_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#auto_restart_action NotificationPolicy#auto_restart_action}
	AutoRestartAction *NotificationPolicyAutoRestartAction `field:"optional" json:"autoRestartAction" yaml:"autoRestartAction"`
	// de_duplication_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#de_duplication_action NotificationPolicy#de_duplication_action}
	DeDuplicationAction *NotificationPolicyDeDuplicationAction `field:"optional" json:"deDuplicationAction" yaml:"deDuplicationAction"`
	// delay_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#delay_action NotificationPolicy#delay_action}
	DelayAction *NotificationPolicyDelayAction `field:"optional" json:"delayAction" yaml:"delayAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#enabled NotificationPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#id NotificationPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#policy_description NotificationPolicy#policy_description}.
	PolicyDescription *string `field:"optional" json:"policyDescription" yaml:"policyDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#suppress NotificationPolicy#suppress}.
	Suppress interface{} `field:"optional" json:"suppress" yaml:"suppress"`
	// time_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/notification_policy#time_restriction NotificationPolicy#time_restriction}
	TimeRestriction *NotificationPolicyTimeRestriction `field:"optional" json:"timeRestriction" yaml:"timeRestriction"`
}

