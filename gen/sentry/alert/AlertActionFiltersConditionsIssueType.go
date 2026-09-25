package alert


type AlertActionFiltersConditionsIssueType struct {
	// If `true`, matches when the issue type equals `value`. If `false`, matches when it does not equal `value`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#include Alert#include}
	Include interface{} `field:"required" json:"include" yaml:"include"`
	// The issue type slug (e.g. `performance_large_http_payload`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

