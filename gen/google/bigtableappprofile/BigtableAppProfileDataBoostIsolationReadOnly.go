package bigtableappprofile


type BigtableAppProfileDataBoostIsolationReadOnly struct {
	// The Compute Billing Owner for this Data Boost App Profile. Possible values: ["HOST_PAYS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigtable_app_profile#compute_billing_owner BigtableAppProfile#compute_billing_owner}
	ComputeBillingOwner *string `field:"required" json:"computeBillingOwner" yaml:"computeBillingOwner"`
}

