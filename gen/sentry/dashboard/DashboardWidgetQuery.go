package dashboard


type DashboardWidgetQuery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#aggregates Dashboard#aggregates}.
	Aggregates *[]*string `field:"optional" json:"aggregates" yaml:"aggregates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#columns Dashboard#columns}.
	Columns *[]*string `field:"optional" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#conditions Dashboard#conditions}.
	Conditions *string `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#field_aliases Dashboard#field_aliases}.
	FieldAliases *[]*string `field:"optional" json:"fieldAliases" yaml:"fieldAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#fields Dashboard#fields}.
	Fields *[]*string `field:"optional" json:"fields" yaml:"fields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#name Dashboard#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/dashboard#order_by Dashboard#order_by}.
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
}

