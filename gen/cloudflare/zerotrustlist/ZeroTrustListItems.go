package zerotrustlist


type ZeroTrustListItems struct {
	// Provide the list item description (optional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_list#description ZeroTrustList#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specify the item value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_list#value ZeroTrustList#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

