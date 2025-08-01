package datacloudflareaccountdnssettingsinternalviews


type DataCloudflareAccountDnsSettingsInternalViewsName struct {
	// Substring of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_dns_settings_internal_views#contains DataCloudflareAccountDnsSettingsInternalViews#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// Suffix of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_dns_settings_internal_views#endswith DataCloudflareAccountDnsSettingsInternalViews#endswith}
	Endswith *string `field:"optional" json:"endswith" yaml:"endswith"`
	// Exact value of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_dns_settings_internal_views#exact DataCloudflareAccountDnsSettingsInternalViews#exact}
	Exact *string `field:"optional" json:"exact" yaml:"exact"`
	// Prefix of the DNS view name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/account_dns_settings_internal_views#startswith DataCloudflareAccountDnsSettingsInternalViews#startswith}
	Startswith *string `field:"optional" json:"startswith" yaml:"startswith"`
}

