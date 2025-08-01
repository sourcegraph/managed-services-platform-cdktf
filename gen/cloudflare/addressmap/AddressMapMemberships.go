package addressmap


type AddressMapMemberships struct {
	// The identifier for the membership (eg. a zone or account tag).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/address_map#identifier AddressMap#identifier}
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// The type of the membership. Available values: "zone", "account".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/address_map#kind AddressMap#kind}
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
}

