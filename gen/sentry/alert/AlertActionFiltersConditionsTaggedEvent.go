package alert


type AlertActionFiltersConditionsTaggedEvent struct {
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#key Alert#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The comparison operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#match Alert#match}
	Match *string `field:"required" json:"match" yaml:"match"`
	// A string. Not required when match is `is` or `ns`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

