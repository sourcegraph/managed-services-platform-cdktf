package alert


type AlertActionFiltersConditionsEventAttribute struct {
	// The attribute to evaluate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#attribute Alert#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The match type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#match Alert#match}
	Match *string `field:"required" json:"match" yaml:"match"`
	// The value to compare against. Not required when `match` is `is` or `ns`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

