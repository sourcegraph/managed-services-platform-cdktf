package googlemigrationcenterpreferenceset


type GoogleMigrationCenterPreferenceSetVirtualMachinePreferencesSoleTenancyPreferences struct {
	// Commitment plan to consider when calculating costs for virtual machine insights and recommendations.
	//
	// If you are unsure which value to set, a 3 year commitment plan is often a good value to start with. Possible values: 'COMMITMENT_PLAN_UNSPECIFIED', 'ON_DEMAND', 'COMMITMENT_1_YEAR', 'COMMITMENT_3_YEAR'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_migration_center_preference_set#commitment_plan GoogleMigrationCenterPreferenceSet#commitment_plan}
	CommitmentPlan *string `field:"optional" json:"commitmentPlan" yaml:"commitmentPlan"`
	// CPU overcommit ratio. Acceptable values are between 1.0 and 2.0 inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_migration_center_preference_set#cpu_overcommit_ratio GoogleMigrationCenterPreferenceSet#cpu_overcommit_ratio}
	CpuOvercommitRatio *float64 `field:"optional" json:"cpuOvercommitRatio" yaml:"cpuOvercommitRatio"`
	// Sole Tenancy nodes maintenance policy. Possible values: 'HOST_MAINTENANCE_POLICY_UNSPECIFIED', 'HOST_MAINTENANCE_POLICY_DEFAULT', 'HOST_MAINTENANCE_POLICY_RESTART_IN_PLACE', 'HOST_MAINTENANCE_POLICY_MIGRATE_WITHIN_NODE_GROUP'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_migration_center_preference_set#host_maintenance_policy GoogleMigrationCenterPreferenceSet#host_maintenance_policy}
	HostMaintenancePolicy *string `field:"optional" json:"hostMaintenancePolicy" yaml:"hostMaintenancePolicy"`
	// node_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_migration_center_preference_set#node_types GoogleMigrationCenterPreferenceSet#node_types}
	NodeTypes interface{} `field:"optional" json:"nodeTypes" yaml:"nodeTypes"`
}

