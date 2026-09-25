package alert


type AlertActionFiltersConditionsIssuePriorityDeescalating struct {
	// The minimum priority threshold required to trigger a de-escalation event.
	//
	// The rule triggers when the historical peak priority meets this threshold, and the current priority drops below it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#comparison Alert#comparison}
	Comparison *float64 `field:"required" json:"comparison" yaml:"comparison"`
}

