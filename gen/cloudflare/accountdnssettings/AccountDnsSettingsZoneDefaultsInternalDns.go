package accountdnssettings


type AccountDnsSettingsZoneDefaultsInternalDns struct {
	// The ID of the zone to fallback to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_dns_settings#reference_zone_id AccountDnsSettings#reference_zone_id}
	ReferenceZoneId *string `field:"optional" json:"referenceZoneId" yaml:"referenceZoneId"`
}

