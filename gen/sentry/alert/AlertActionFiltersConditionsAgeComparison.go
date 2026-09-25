package alert


type AlertActionFiltersConditionsAgeComparison struct {
	// The type of comparison to perform. Valid values are: `older`, and `newer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#comparison_type Alert#comparison_type}
	ComparisonType *string `field:"required" json:"comparisonType" yaml:"comparisonType"`
	// The unit of time for the age comparison. Valid values are: `minute`, `hour`, `day`, and `week`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#time Alert#time}
	Time *string `field:"required" json:"time" yaml:"time"`
	// The value of the age comparison.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

