package list


type ListItem struct {
	// An optional comment for the item.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list#comment List#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/list#value List#value}
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

