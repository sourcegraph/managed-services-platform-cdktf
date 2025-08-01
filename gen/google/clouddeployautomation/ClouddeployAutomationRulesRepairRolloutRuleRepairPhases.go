package clouddeployautomation


type ClouddeployAutomationRulesRepairRolloutRuleRepairPhases struct {
	// retry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_automation#retry ClouddeployAutomation#retry}
	Retry *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRetry `field:"optional" json:"retry" yaml:"retry"`
	// rollback block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_automation#rollback ClouddeployAutomation#rollback}
	Rollback *ClouddeployAutomationRulesRepairRolloutRuleRepairPhasesRollback `field:"optional" json:"rollback" yaml:"rollback"`
}

