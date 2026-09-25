package alert


type AlertActionFiltersConditionsLevel struct {
	// The level to compare against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#level Alert#level}
	Level *float64 `field:"required" json:"level" yaml:"level"`
	// The comparison operator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#match Alert#match}
	Match *string `field:"required" json:"match" yaml:"match"`
}

