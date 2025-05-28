package bigtableappprofile


type BigtableAppProfileStandardIsolation struct {
	// The priority of requests sent using this app profile. Possible values: ["PRIORITY_LOW", "PRIORITY_MEDIUM", "PRIORITY_HIGH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/bigtable_app_profile#priority BigtableAppProfile#priority}
	Priority *string `field:"required" json:"priority" yaml:"priority"`
}

