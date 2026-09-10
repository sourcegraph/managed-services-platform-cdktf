package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicySchedule struct {
	// Specify the time intervals when the rule is active on Fridays, in the increasing order from 00:00-24:00.
	//
	// If this parameter omitted, the rule is deactivated on Fridays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#fri ZeroTrustGatewayPolicy#fri}
	Fri *string `field:"optional" json:"fri" yaml:"fri"`
	// Specify the time intervals when the rule is active on Mondays, in the increasing order from 00:00-24:00(capped at maximum of 6 time splits).
	//
	// If this parameter omitted, the rule is deactivated on Mondays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#mon ZeroTrustGatewayPolicy#mon}
	Mon *string `field:"optional" json:"mon" yaml:"mon"`
	// Specify the time intervals when the rule is active on Saturdays, in the increasing order from 00:00-24:00.
	//
	// If this parameter omitted, the rule is deactivated on Saturdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#sat ZeroTrustGatewayPolicy#sat}
	Sat *string `field:"optional" json:"sat" yaml:"sat"`
	// Specify the time intervals when the rule is active on Sundays, in the increasing order from 00:00-24:00.
	//
	// If this parameter omitted, the rule is deactivated on Sundays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#sun ZeroTrustGatewayPolicy#sun}
	Sun *string `field:"optional" json:"sun" yaml:"sun"`
	// Specify the time intervals when the rule is active on Thursdays, in the increasing order from 00:00-24:00.
	//
	// If this parameter omitted, the rule is deactivated on Thursdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#thu ZeroTrustGatewayPolicy#thu}
	Thu *string `field:"optional" json:"thu" yaml:"thu"`
	// Specify the time zone for rule evaluation.
	//
	// When a [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List) is provided, Gateway always uses the current time for that time zone. When this parameter is omitted, Gateway uses the time zone determined from the user's IP address. Colo time zone is used when the user's IP address does not resolve to a location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#time_zone ZeroTrustGatewayPolicy#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// Specify the time intervals when the rule is active on Tuesdays, in the increasing order from 00:00-24:00.
	//
	// If this parameter omitted, the rule is deactivated on Tuesdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#tue ZeroTrustGatewayPolicy#tue}
	Tue *string `field:"optional" json:"tue" yaml:"tue"`
	// Specify the time intervals when the rule is active on Wednesdays, in the increasing order from 00:00-24:00.
	//
	// If this parameter omitted, the rule is deactivated on Wednesdays. API returns a formatted version of this string, which may cause Terraform drift if a unformatted value is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#wed ZeroTrustGatewayPolicy#wed}
	Wed *string `field:"optional" json:"wed" yaml:"wed"`
}

