package alert

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AlertConfig struct {
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
	// The filters to run before the action will fire and the action(s) to fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#action_filters Alert#action_filters}
	ActionFilters interface{} `field:"required" json:"actionFilters" yaml:"actionFilters"`
	// How often the alert should fire in minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#frequency_minutes Alert#frequency_minutes}
	FrequencyMinutes *float64 `field:"required" json:"frequencyMinutes" yaml:"frequencyMinutes"`
	// The IDs of the monitors to create alerts for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#monitor_ids Alert#monitor_ids}
	MonitorIds *[]*string `field:"required" json:"monitorIds" yaml:"monitorIds"`
	// The name of this alert.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#name Alert#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization slug or internal ID to create the alert for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#organization Alert#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Whether the alert is enabled. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#enabled Alert#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The environment to filter alerts to. Omit or set to `null` to apply to all environments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#environment Alert#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// ⚠️ The trigger condition types listed here are not natively supported by this provider and may be deprecated by Sentry in a future API version.
	//
	// Trigger condition types present on this alert that are not representable in `trigger_conditions` (e.g. `new_high_priority_issue`, `existing_high_priority_issue`, `issue_resolution_change`). When omitted from config these will be removed on the next apply. Set explicitly to preserve them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#legacy_trigger_conditions Alert#legacy_trigger_conditions}
	LegacyTriggerConditions *[]*string `field:"optional" json:"legacyTriggerConditions" yaml:"legacyTriggerConditions"`
	// The conditions on which the alert will trigger.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#trigger_conditions Alert#trigger_conditions}
	TriggerConditions interface{} `field:"optional" json:"triggerConditions" yaml:"triggerConditions"`
}

