package iaptunneliammember


type IapTunnelIamMemberCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iap_tunnel_iam_member#expression IapTunnelIamMember#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iap_tunnel_iam_member#title IapTunnelIamMember#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iap_tunnel_iam_member#description IapTunnelIamMember#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

