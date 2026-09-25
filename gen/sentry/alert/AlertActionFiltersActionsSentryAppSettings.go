package alert


type AlertActionFiltersActionsSentryAppSettings struct {
	// The name of the setting field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#name Alert#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the setting field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#value Alert#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// The human-readable display label for the value.
	//
	// Required for async select fields whose option list is paginated by the third-party app — without it the field may appear blank under certain conditions in the Sentry UI after apply.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#label Alert#label}
	Label *string `field:"optional" json:"label" yaml:"label"`
}

