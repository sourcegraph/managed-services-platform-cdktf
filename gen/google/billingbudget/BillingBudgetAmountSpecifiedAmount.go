package billingbudget


type BillingBudgetAmountSpecifiedAmount struct {
	// The 3-letter currency code defined in ISO 4217.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/billing_budget#currency_code BillingBudget#currency_code}
	CurrencyCode *string `field:"optional" json:"currencyCode" yaml:"currencyCode"`
	// Number of nano (10^-9) units of the amount.
	//
	// The value must be between -999,999,999 and +999,999,999
	// inclusive. If units is positive, nanos must be positive or
	// zero. If units is zero, nanos can be positive, zero, or
	// negative. If units is negative, nanos must be negative or
	// zero. For example $-1.75 is represented as units=-1 and
	// nanos=-750,000,000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/billing_budget#nanos BillingBudget#nanos}
	Nanos *float64 `field:"optional" json:"nanos" yaml:"nanos"`
	// The whole units of the amount. For example if currencyCode is "USD", then 1 unit is one US dollar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/billing_budget#units BillingBudget#units}
	Units *string `field:"optional" json:"units" yaml:"units"`
}

