package zonednssettings


type ZoneDnsSettingsNameservers struct {
	// Configured nameserver set to be used for this zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone_dns_settings#ns_set ZoneDnsSettings#ns_set}
	NsSet *float64 `field:"optional" json:"nsSet" yaml:"nsSet"`
	// Nameserver type Available values: "cloudflare.standard", "custom.account", "custom.tenant", "custom.zone".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zone_dns_settings#type ZoneDnsSettings#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

