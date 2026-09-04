package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettings struct {
	// Add custom headers to allowed requests as key-value pairs.
	//
	// Use header names as keys that map to arrays of header values. Settable only for `http` rules with the action set to `allow`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#add_headers ZeroTrustGatewayPolicy#add_headers}
	AddHeaders interface{} `field:"optional" json:"addHeaders" yaml:"addHeaders"`
	// Set to enable MSP children to bypass this rule.
	//
	// Only parent MSP accounts can set this. this rule. Settable for all types of rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#allow_child_bypass ZeroTrustGatewayPolicy#allow_child_bypass}
	AllowChildBypass interface{} `field:"optional" json:"allowChildBypass" yaml:"allowChildBypass"`
	// Define the settings for the Audit SSH action. Settable only for `l4` rules with `audit_ssh` action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#audit_ssh ZeroTrustGatewayPolicy#audit_ssh}
	AuditSsh *ZeroTrustGatewayPolicyRuleSettingsAuditSsh `field:"optional" json:"auditSsh" yaml:"auditSsh"`
	// Configure browser isolation behavior. Settable only for `http` rules with the action set to `isolate`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#biso_admin_controls ZeroTrustGatewayPolicy#biso_admin_controls}
	BisoAdminControls *ZeroTrustGatewayPolicyRuleSettingsBisoAdminControls `field:"optional" json:"bisoAdminControls" yaml:"bisoAdminControls"`
	// Configure custom block page settings.
	//
	// If missing or null, use the account settings. Settable only for `http` rules with the action set to `block`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#block_page ZeroTrustGatewayPolicy#block_page}
	BlockPage *ZeroTrustGatewayPolicyRuleSettingsBlockPage `field:"optional" json:"blockPage" yaml:"blockPage"`
	// Enable the custom block page. Settable only for `dns` rules with action `block`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#block_page_enabled ZeroTrustGatewayPolicy#block_page_enabled}
	BlockPageEnabled interface{} `field:"optional" json:"blockPageEnabled" yaml:"blockPageEnabled"`
	// Explain why the rule blocks the request.
	//
	// The custom block page shows this text (if enabled). Settable only for `dns`, `l4`, and `http` rules when the action set to `block`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#block_reason ZeroTrustGatewayPolicy#block_reason}
	BlockReason *string `field:"optional" json:"blockReason" yaml:"blockReason"`
	// Set to enable MSP accounts to bypass their parent's rules.
	//
	// Only MSP child accounts can set this. Settable for all types of rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#bypass_parent_rule ZeroTrustGatewayPolicy#bypass_parent_rule}
	BypassParentRule interface{} `field:"optional" json:"bypassParentRule" yaml:"bypassParentRule"`
	// Configure session check behavior. Settable only for `l4` and `http` rules with the action set to `allow`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#check_session ZeroTrustGatewayPolicy#check_session}
	CheckSession *ZeroTrustGatewayPolicyRuleSettingsCheckSession `field:"optional" json:"checkSession" yaml:"checkSession"`
	// Configure custom resolvers to route queries that match the resolver policy.
	//
	// Unused with 'resolve_dns_through_cloudflare' or 'resolve_dns_internally' settings. DNS queries get routed to the address closest to their origin. Only valid when a rule's action set to 'resolve'. Settable only for `dns_resolver` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#dns_resolvers ZeroTrustGatewayPolicy#dns_resolvers}
	DnsResolvers *ZeroTrustGatewayPolicyRuleSettingsDnsResolvers `field:"optional" json:"dnsResolvers" yaml:"dnsResolvers"`
	// Configure how Gateway Proxy traffic egresses.
	//
	// You can enable this setting for rules with Egress actions and filters, or omit it to indicate local egress via WARP IPs. Settable only for `egress` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#egress ZeroTrustGatewayPolicy#egress}
	Egress *ZeroTrustGatewayPolicyRuleSettingsEgress `field:"optional" json:"egress" yaml:"egress"`
	// Configure whether a copy of the HTTP request will be sent to storage when the rule matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#forensic_copy ZeroTrustGatewayPolicy#forensic_copy}
	ForensicCopy *ZeroTrustGatewayPolicyRuleSettingsForensicCopy `field:"optional" json:"forensicCopy" yaml:"forensicCopy"`
	// Ignore category matches at CNAME domains in a response.
	//
	// When off, evaluate categories in this rule against all CNAME domain categories in the response. Settable only for `dns` and `dns_resolver` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#ignore_cname_category_matches ZeroTrustGatewayPolicy#ignore_cname_category_matches}
	IgnoreCnameCategoryMatches interface{} `field:"optional" json:"ignoreCnameCategoryMatches" yaml:"ignoreCnameCategoryMatches"`
	// Specify whether to disable DNSSEC validation (for Allow actions) [INSECURE]. Settable only for `dns` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#insecure_disable_dnssec_validation ZeroTrustGatewayPolicy#insecure_disable_dnssec_validation}
	InsecureDisableDnssecValidation interface{} `field:"optional" json:"insecureDisableDnssecValidation" yaml:"insecureDisableDnssecValidation"`
	// Enable IPs in DNS resolver category blocks.
	//
	// The system blocks only domain name categories unless you enable this setting. Settable only for `dns` and `dns_resolver` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#ip_categories ZeroTrustGatewayPolicy#ip_categories}
	IpCategories interface{} `field:"optional" json:"ipCategories" yaml:"ipCategories"`
	// Indicates whether to include IPs in DNS resolver indicator feed blocks.
	//
	// Default, indicator feeds block only domain names. Settable only for `dns` and `dns_resolver` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#ip_indicator_feeds ZeroTrustGatewayPolicy#ip_indicator_feeds}
	IpIndicatorFeeds interface{} `field:"optional" json:"ipIndicatorFeeds" yaml:"ipIndicatorFeeds"`
	// Send matching traffic to the supplied destination IP address and port.
	//
	// Settable only for `l4` rules with the action set to `l4_override`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#l4override ZeroTrustGatewayPolicy#l4override}
	L4Override *ZeroTrustGatewayPolicyRuleSettingsL4Override `field:"optional" json:"l4Override" yaml:"l4Override"`
	// Configure a notification to display on the user's device when this rule matched.
	//
	// Settable for all types of rules with the action set to `block`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#notification_settings ZeroTrustGatewayPolicy#notification_settings}
	NotificationSettings *ZeroTrustGatewayPolicyRuleSettingsNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
	// Defines a hostname for override, for the matching DNS queries.
	//
	// Settable only for `dns` rules with the action set to `override`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#override_host ZeroTrustGatewayPolicy#override_host}
	OverrideHost *string `field:"optional" json:"overrideHost" yaml:"overrideHost"`
	// Defines a an IP or set of IPs for overriding matched DNS queries.
	//
	// Settable only for `dns` rules with the action set to `override`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#override_ips ZeroTrustGatewayPolicy#override_ips}
	OverrideIps *[]*string `field:"optional" json:"overrideIps" yaml:"overrideIps"`
	// Configure DLP payload logging. Settable only for `http` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#payload_log ZeroTrustGatewayPolicy#payload_log}
	PayloadLog *ZeroTrustGatewayPolicyRuleSettingsPayloadLog `field:"optional" json:"payloadLog" yaml:"payloadLog"`
	// Configure settings that apply to quarantine rules. Settable only for `http` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#quarantine ZeroTrustGatewayPolicy#quarantine}
	Quarantine *ZeroTrustGatewayPolicyRuleSettingsQuarantine `field:"optional" json:"quarantine" yaml:"quarantine"`
	// Apply settings to redirect rules. Settable only for `http` rules with the action set to `redirect`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#redirect ZeroTrustGatewayPolicy#redirect}
	Redirect *ZeroTrustGatewayPolicyRuleSettingsRedirect `field:"optional" json:"redirect" yaml:"redirect"`
	// Configure to forward the query to the internal DNS service, passing the specified 'view_id' as input.
	//
	// Not used when 'dns_resolvers' is specified or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action set to 'resolve'. Settable only for `dns_resolver` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#resolve_dns_internally ZeroTrustGatewayPolicy#resolve_dns_internally}
	ResolveDnsInternally *ZeroTrustGatewayPolicyRuleSettingsResolveDnsInternally `field:"optional" json:"resolveDnsInternally" yaml:"resolveDnsInternally"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS resolver. Cannot set when 'dns_resolvers' specified or 'resolve_dns_internally' is set. Only valid when a rule's action set to 'resolve'. Settable only for `dns_resolver` rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#resolve_dns_through_cloudflare ZeroTrustGatewayPolicy#resolve_dns_through_cloudflare}
	ResolveDnsThroughCloudflare interface{} `field:"optional" json:"resolveDnsThroughCloudflare" yaml:"resolveDnsThroughCloudflare"`
	// Configure behavior when an upstream certificate is invalid or an SSL error occurs.
	//
	// Settable only for `http` rules with the action set to `allow`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_gateway_policy#untrusted_cert ZeroTrustGatewayPolicy#untrusted_cert}
	UntrustedCert *ZeroTrustGatewayPolicyRuleSettingsUntrustedCert `field:"optional" json:"untrustedCert" yaml:"untrustedCert"`
}

