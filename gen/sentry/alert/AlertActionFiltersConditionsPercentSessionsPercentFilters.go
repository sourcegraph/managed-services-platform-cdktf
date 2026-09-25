package alert


type AlertActionFiltersConditionsPercentSessionsPercentFilters struct {
	// The attribute of the filter. Conflicts with `key`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#attribute Alert#attribute}
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// The key of the filter. Conflicts with `attribute`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#key Alert#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The match type of the filter.
	//
	// Valid values are: `co`, `ew`, `eq`, `gte`, `gt`, `is`, `in`, `lte`, `lt`, `nc`, `new`, `ne`, `ns`, `nsw`, `nin`, and `sw`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#match Alert#match}
	Match *string `field:"optional" json:"match" yaml:"match"`
	// The value of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

