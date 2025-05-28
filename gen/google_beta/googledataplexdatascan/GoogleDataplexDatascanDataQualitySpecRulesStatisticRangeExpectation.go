package googledataplexdatascan


type GoogleDataplexDatascanDataQualitySpecRulesStatisticRangeExpectation struct {
	// column statistics. Possible values: ["STATISTIC_UNDEFINED", "MEAN", "MIN", "MAX"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dataplex_datascan#statistic GoogleDataplexDatascan#statistic}
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
	// The maximum column statistic value allowed for a row to pass this validation.
	//
	// At least one of minValue and maxValue need to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dataplex_datascan#max_value GoogleDataplexDatascan#max_value}
	MaxValue *string `field:"optional" json:"maxValue" yaml:"maxValue"`
	// The minimum column statistic value allowed for a row to pass this validation.
	//
	// At least one of minValue and maxValue need to be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dataplex_datascan#min_value GoogleDataplexDatascan#min_value}
	MinValue *string `field:"optional" json:"minValue" yaml:"minValue"`
	// Whether column statistic needs to be strictly lesser than ('<') the maximum, or if equality is allowed.
	//
	// Only relevant if a maxValue has been defined. Default = false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dataplex_datascan#strict_max_enabled GoogleDataplexDatascan#strict_max_enabled}
	StrictMaxEnabled interface{} `field:"optional" json:"strictMaxEnabled" yaml:"strictMaxEnabled"`
	// Whether column statistic needs to be strictly greater than ('>') the minimum, or if equality is allowed.
	//
	// Only relevant if a minValue has been defined. Default = false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_dataplex_datascan#strict_min_enabled GoogleDataplexDatascan#strict_min_enabled}
	StrictMinEnabled interface{} `field:"optional" json:"strictMinEnabled" yaml:"strictMinEnabled"`
}

