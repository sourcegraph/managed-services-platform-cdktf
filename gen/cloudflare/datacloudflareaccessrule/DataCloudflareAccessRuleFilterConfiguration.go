package datacloudflareaccessrule


type DataCloudflareAccessRuleFilterConfiguration struct {
	// Defines the target to search in existing rules. Available values: "ip", "ip_range", "asn", "country".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/access_rule#target DataCloudflareAccessRule#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Defines the target value to search for in existing rules: an IP address, an IP address range, or a country code, depending on the provided `configuration.target`. Notes: You can search for a single IPv4 address, an IP address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1 alpha-2 country code.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/access_rule#value DataCloudflareAccessRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

