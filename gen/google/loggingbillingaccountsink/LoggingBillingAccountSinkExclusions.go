package loggingbillingaccountsink


type LoggingBillingAccountSinkExclusions struct {
	// An advanced logs filter that matches the log entries to be excluded.
	//
	// By using the sample function, you can exclude less than 100% of the matching log entries
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_billing_account_sink#filter LoggingBillingAccountSink#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// A client-assigned identifier, such as "load-balancer-exclusion".
	//
	// Identifiers are limited to 100 characters and can include only letters, digits, underscores, hyphens, and periods. First character has to be alphanumeric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_billing_account_sink#name LoggingBillingAccountSink#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A description of this exclusion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_billing_account_sink#description LoggingBillingAccountSink#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If set to True, then this exclusion is disabled and it does not exclude any log entries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/logging_billing_account_sink#disabled LoggingBillingAccountSink#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
}

