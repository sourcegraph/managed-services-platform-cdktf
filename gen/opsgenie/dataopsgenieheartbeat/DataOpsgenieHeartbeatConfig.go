package dataopsgenieheartbeat

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpsgenieHeartbeatConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#name DataOpsgenieHeartbeat#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#alert_message DataOpsgenieHeartbeat#alert_message}.
	AlertMessage *string `field:"optional" json:"alertMessage" yaml:"alertMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#alert_priority DataOpsgenieHeartbeat#alert_priority}.
	AlertPriority *string `field:"optional" json:"alertPriority" yaml:"alertPriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#alert_tags DataOpsgenieHeartbeat#alert_tags}.
	AlertTags *[]*string `field:"optional" json:"alertTags" yaml:"alertTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#description DataOpsgenieHeartbeat#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#enabled DataOpsgenieHeartbeat#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#id DataOpsgenieHeartbeat#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#interval DataOpsgenieHeartbeat#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#interval_unit DataOpsgenieHeartbeat#interval_unit}.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/data-sources/heartbeat#owner_team_id DataOpsgenieHeartbeat#owner_team_id}.
	OwnerTeamId *string `field:"optional" json:"ownerTeamId" yaml:"ownerTeamId"`
}

