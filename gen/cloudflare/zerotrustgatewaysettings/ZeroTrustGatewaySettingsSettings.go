package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettings struct {
	// Specify activity log settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#activity_log ZeroTrustGatewaySettings#activity_log}
	ActivityLog *ZeroTrustGatewaySettingsSettingsActivityLog `field:"optional" json:"activityLog" yaml:"activityLog"`
	// Specify anti-virus settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#antivirus ZeroTrustGatewaySettings#antivirus}
	Antivirus *ZeroTrustGatewaySettingsSettingsAntivirus `field:"optional" json:"antivirus" yaml:"antivirus"`
	// Specify block page layout settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#block_page ZeroTrustGatewaySettings#block_page}
	BlockPage *ZeroTrustGatewaySettingsSettingsBlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// Specify the DLP inspection mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#body_scanning ZeroTrustGatewaySettings#body_scanning}
	BodyScanning *ZeroTrustGatewaySettingsSettingsBodyScanning `field:"optional" json:"bodyScanning" yaml:"bodyScanning"`
	// Specify Clientless Browser Isolation settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#browser_isolation ZeroTrustGatewaySettings#browser_isolation}
	BrowserIsolation *ZeroTrustGatewaySettingsSettingsBrowserIsolation `field:"optional" json:"browserIsolation" yaml:"browserIsolation"`
	// Specify certificate settings for Gateway TLS interception. If unset, the Cloudflare Root CA handles interception.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#certificate ZeroTrustGatewaySettings#certificate}
	Certificate *ZeroTrustGatewaySettingsSettingsCertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Specify custom certificate settings for BYO-PKI. This field is deprecated; use `certificate` instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#custom_certificate ZeroTrustGatewaySettings#custom_certificate}
	CustomCertificate *ZeroTrustGatewaySettingsSettingsCustomCertificate `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// Configures user email settings for firewall policies.
	//
	// When you enable this, the system standardizes email addresses in the identity portion of the rule to match extended email variants in firewall policies. When you disable this setting, the system matches email addresses exactly as you provide them. Enable this setting if your email uses `.` or `+` modifiers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#extended_email_matching ZeroTrustGatewaySettings#extended_email_matching}
	ExtendedEmailMatching *ZeroTrustGatewaySettingsSettingsExtendedEmailMatching `field:"optional" json:"extendedEmailMatching" yaml:"extendedEmailMatching"`
	// Specify FIPS settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#fips ZeroTrustGatewaySettings#fips}
	Fips *ZeroTrustGatewaySettingsSettingsFips `field:"optional" json:"fips" yaml:"fips"`
	// Enable host selection in egress policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#host_selector ZeroTrustGatewaySettings#host_selector}
	HostSelector *ZeroTrustGatewaySettingsSettingsHostSelector `field:"optional" json:"hostSelector" yaml:"hostSelector"`
	// Define the proxy inspection mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#inspection ZeroTrustGatewaySettings#inspection}
	Inspection *ZeroTrustGatewaySettingsSettingsInspection `field:"optional" json:"inspection" yaml:"inspection"`
	// Specify whether to detect protocols from the initial bytes of client traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#protocol_detection ZeroTrustGatewaySettings#protocol_detection}
	ProtocolDetection *ZeroTrustGatewaySettingsSettingsProtocolDetection `field:"optional" json:"protocolDetection" yaml:"protocolDetection"`
	// Specify whether to enable the sandbox.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#sandbox ZeroTrustGatewaySettings#sandbox}
	Sandbox *ZeroTrustGatewaySettingsSettingsSandbox `field:"optional" json:"sandbox" yaml:"sandbox"`
	// Specify whether to inspect encrypted HTTP traffic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_settings#tls_decrypt ZeroTrustGatewaySettings#tls_decrypt}
	TlsDecrypt *ZeroTrustGatewaySettingsSettingsTlsDecrypt `field:"optional" json:"tlsDecrypt" yaml:"tlsDecrypt"`
}

