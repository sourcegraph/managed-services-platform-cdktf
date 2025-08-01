package datacloudflareemailroutingaddress


type DataCloudflareEmailRoutingAddressFilter struct {
	// Sorts results in an ascending or descending order. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/email_routing_address#direction DataCloudflareEmailRoutingAddress#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Filter by verified destination addresses.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/email_routing_address#verified DataCloudflareEmailRoutingAddress#verified}
	Verified interface{} `field:"optional" json:"verified" yaml:"verified"`
}

