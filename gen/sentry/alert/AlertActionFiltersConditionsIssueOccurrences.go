package alert


type AlertActionFiltersConditionsIssueOccurrences struct {
	// A positive integer representing how many times the issue has to happen before the alert will fire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

