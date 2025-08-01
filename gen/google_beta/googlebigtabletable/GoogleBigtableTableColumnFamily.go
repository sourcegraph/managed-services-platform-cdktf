package googlebigtabletable


type GoogleBigtableTableColumnFamily struct {
	// The name of the column family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_table#family GoogleBigtableTable#family}
	Family *string `field:"required" json:"family" yaml:"family"`
	// The type of the column family.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigtable_table#type GoogleBigtableTable#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

