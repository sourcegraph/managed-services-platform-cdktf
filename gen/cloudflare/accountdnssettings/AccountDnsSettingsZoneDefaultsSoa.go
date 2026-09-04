package accountdnssettings


type AccountDnsSettingsZoneDefaultsSoa struct {
	// Time in seconds of being unable to query the primary server after which secondary servers should stop serving the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings#expire AccountDnsSettings#expire}
	Expire *float64 `field:"optional" json:"expire" yaml:"expire"`
	// The time to live (TTL) for negative caching of records within the zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings#min_ttl AccountDnsSettings#min_ttl}
	MinTtl *float64 `field:"optional" json:"minTtl" yaml:"minTtl"`
	// The primary nameserver, which may be used for outbound zone transfers. If null, a Cloudflare-assigned value will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings#mname AccountDnsSettings#mname}
	Mname *string `field:"optional" json:"mname" yaml:"mname"`
	// Time in seconds after which secondary servers should re-check the SOA record to see if the zone has been updated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings#refresh AccountDnsSettings#refresh}
	Refresh *float64 `field:"optional" json:"refresh" yaml:"refresh"`
	// Time in seconds after which secondary servers should retry queries after the primary server was unresponsive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings#retry AccountDnsSettings#retry}
	Retry *float64 `field:"optional" json:"retry" yaml:"retry"`
	// The email address of the zone administrator, with the first label representing the local part of the email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings#rname AccountDnsSettings#rname}
	Rname *string `field:"optional" json:"rname" yaml:"rname"`
	// The time to live (TTL) of the SOA record itself.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/account_dns_settings#ttl AccountDnsSettings#ttl}
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

