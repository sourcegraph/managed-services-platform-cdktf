package googlecomputeinstantsnapshotiambinding


type GoogleComputeInstantSnapshotIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instant_snapshot_iam_binding#expression GoogleComputeInstantSnapshotIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instant_snapshot_iam_binding#title GoogleComputeInstantSnapshotIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_compute_instant_snapshot_iam_binding#description GoogleComputeInstantSnapshotIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

