package alert


type AlertActionFiltersConditionsIssueCategory struct {
	// The issue category to filter to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// If `true`, the condition matches when the issue category is equal to `value`.
	//
	// If `false`, the condition matches when it is not equal. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#include Alert#include}
	Include interface{} `field:"optional" json:"include" yaml:"include"`
}

