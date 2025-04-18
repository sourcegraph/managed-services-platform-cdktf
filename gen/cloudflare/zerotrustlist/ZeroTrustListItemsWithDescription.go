package zerotrustlist


type ZeroTrustListItemsWithDescription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_list#description ZeroTrustList#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_list#value ZeroTrustList#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

