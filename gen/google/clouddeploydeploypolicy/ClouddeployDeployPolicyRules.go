package clouddeploydeploypolicy


type ClouddeployDeployPolicyRules struct {
	// rollout_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_deploy_policy#rollout_restriction ClouddeployDeployPolicy#rollout_restriction}
	RolloutRestriction *ClouddeployDeployPolicyRulesRolloutRestriction `field:"optional" json:"rolloutRestriction" yaml:"rolloutRestriction"`
}

