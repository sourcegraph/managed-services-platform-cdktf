package googlecloudrunserviceiammember


type GoogleCloudRunServiceIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_service_iam_member#expression GoogleCloudRunServiceIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_service_iam_member#title GoogleCloudRunServiceIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_service_iam_member#description GoogleCloudRunServiceIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

