package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsTagResourcesTagConditionsSensitivityScore struct {
	// The sensitivity score applied to the resource. Possible values: ["SENSITIVITY_LOW", "SENSITIVITY_MODERATE", "SENSITIVITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#score DataLossPreventionDiscoveryConfig#score}
	Score *string `field:"required" json:"score" yaml:"score"`
}

