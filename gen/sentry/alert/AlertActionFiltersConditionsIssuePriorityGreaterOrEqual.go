package alert


type AlertActionFiltersConditionsIssuePriorityGreaterOrEqual struct {
	// The priority the issue must be for the alert to fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#comparison Alert#comparison}
	Comparison *float64 `field:"required" json:"comparison" yaml:"comparison"`
}

