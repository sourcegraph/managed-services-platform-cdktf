package alert


type AlertActionFiltersConditionsAssignedTo struct {
	// The internal ID of the user or team. Only required if the target type is `Member` or `Team`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#target_id Alert#target_id}
	TargetId *string `field:"optional" json:"targetId" yaml:"targetId"`
	// Who the issue is assigned to. Valid values are: `Unassigned`, `Member`, and `Team`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#target_type Alert#target_type}
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
}

