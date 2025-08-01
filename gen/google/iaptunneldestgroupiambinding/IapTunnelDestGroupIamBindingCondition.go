package iaptunneldestgroupiambinding


type IapTunnelDestGroupIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iap_tunnel_dest_group_iam_binding#expression IapTunnelDestGroupIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iap_tunnel_dest_group_iam_binding#title IapTunnelDestGroupIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iap_tunnel_dest_group_iam_binding#description IapTunnelDestGroupIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

