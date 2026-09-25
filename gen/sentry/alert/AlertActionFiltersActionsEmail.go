package alert


type AlertActionFiltersActionsEmail struct {
	// The type of recipient to notify. Valid values are: `issue_owners`, `team`, and `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#target_type Alert#target_type}
	TargetType *string `field:"required" json:"targetType" yaml:"targetType"`
	// The type of fallthrough to apply when choosing to notify issue owners.
	//
	// Only required if the target type is `issue_owners`. Valid values are: `AllMembers`, `ActiveMembers`, and `NoOne`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#fallthrough_type Alert#fallthrough_type}
	FallthroughType *string `field:"optional" json:"fallthroughType" yaml:"fallthroughType"`
	// The internal ID of the user or team. Only required if the target type is `team` or `user`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#target_id Alert#target_id}
	TargetId *string `field:"optional" json:"targetId" yaml:"targetId"`
}

