package datalosspreventiondiscoveryconfig


type DataLossPreventionDiscoveryConfigActionsTagResourcesTagConditions struct {
	// sensitivity_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#sensitivity_score DataLossPreventionDiscoveryConfig#sensitivity_score}
	SensitivityScore *DataLossPreventionDiscoveryConfigActionsTagResourcesTagConditionsSensitivityScore `field:"optional" json:"sensitivityScore" yaml:"sensitivityScore"`
	// tag block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/data_loss_prevention_discovery_config#tag DataLossPreventionDiscoveryConfig#tag}
	Tag *DataLossPreventionDiscoveryConfigActionsTagResourcesTagConditionsTag `field:"optional" json:"tag" yaml:"tag"`
}

