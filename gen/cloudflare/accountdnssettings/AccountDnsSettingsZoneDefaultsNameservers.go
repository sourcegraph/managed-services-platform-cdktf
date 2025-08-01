package accountdnssettings


type AccountDnsSettingsZoneDefaultsNameservers struct {
	// Nameserver type Available values: "cloudflare.standard", "cloudflare.standard.random", "custom.account", "custom.tenant".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_dns_settings#type AccountDnsSettings#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

