package dataretentionpolicy


type DataRetentionPolicyDeleteOlderThan struct {
	// Number of days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/data_retention_policy#days DataRetentionPolicy#days}
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

