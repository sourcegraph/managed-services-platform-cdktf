package zerotrustaccesspolicy


type ZeroTrustAccessPolicyInclude struct {
	// Matches any valid Access service token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#any_valid_service_token ZeroTrustAccessPolicy#any_valid_service_token}
	AnyValidServiceToken interface{} `field:"optional" json:"anyValidServiceToken" yaml:"anyValidServiceToken"`
	// auth_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#auth_context ZeroTrustAccessPolicy#auth_context}
	AuthContext interface{} `field:"optional" json:"authContext" yaml:"authContext"`
	// The type of authentication method. Refer to https://datatracker.ietf.org/doc/html/rfc8176#section-2 for possible types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#auth_method ZeroTrustAccessPolicy#auth_method}
	AuthMethod *string `field:"optional" json:"authMethod" yaml:"authMethod"`
	// azure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#azure ZeroTrustAccessPolicy#azure}
	Azure interface{} `field:"optional" json:"azure" yaml:"azure"`
	// Matches any valid client certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#certificate ZeroTrustAccessPolicy#certificate}
	Certificate interface{} `field:"optional" json:"certificate" yaml:"certificate"`
	// Matches a valid client certificate common name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#common_name ZeroTrustAccessPolicy#common_name}
	CommonName *string `field:"optional" json:"commonName" yaml:"commonName"`
	// Overflow field if you need to have multiple common_name rules in a single policy.
	//
	// Use in place of the singular common_name field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#common_names ZeroTrustAccessPolicy#common_names}
	CommonNames *[]*string `field:"optional" json:"commonNames" yaml:"commonNames"`
	// The ID of a device posture integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#device_posture ZeroTrustAccessPolicy#device_posture}
	DevicePosture *[]*string `field:"optional" json:"devicePosture" yaml:"devicePosture"`
	// The email of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#email ZeroTrustAccessPolicy#email}
	Email *[]*string `field:"optional" json:"email" yaml:"email"`
	// The email domain to match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#email_domain ZeroTrustAccessPolicy#email_domain}
	EmailDomain *[]*string `field:"optional" json:"emailDomain" yaml:"emailDomain"`
	// The ID of a previously created email list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#email_list ZeroTrustAccessPolicy#email_list}
	EmailList *[]*string `field:"optional" json:"emailList" yaml:"emailList"`
	// Matches everyone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#everyone ZeroTrustAccessPolicy#everyone}
	Everyone interface{} `field:"optional" json:"everyone" yaml:"everyone"`
	// external_evaluation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#external_evaluation ZeroTrustAccessPolicy#external_evaluation}
	ExternalEvaluation interface{} `field:"optional" json:"externalEvaluation" yaml:"externalEvaluation"`
	// Matches a specific country.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#geo ZeroTrustAccessPolicy#geo}
	Geo *[]*string `field:"optional" json:"geo" yaml:"geo"`
	// github block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#github ZeroTrustAccessPolicy#github}
	Github interface{} `field:"optional" json:"github" yaml:"github"`
	// The ID of a previously created Access group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#group ZeroTrustAccessPolicy#group}
	Group *[]*string `field:"optional" json:"group" yaml:"group"`
	// gsuite block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#gsuite ZeroTrustAccessPolicy#gsuite}
	Gsuite interface{} `field:"optional" json:"gsuite" yaml:"gsuite"`
	// An IPv4 or IPv6 CIDR block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#ip ZeroTrustAccessPolicy#ip}
	Ip *[]*string `field:"optional" json:"ip" yaml:"ip"`
	// The ID of a previously created IP list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#ip_list ZeroTrustAccessPolicy#ip_list}
	IpList *[]*string `field:"optional" json:"ipList" yaml:"ipList"`
	// The ID of a configured identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#login_method ZeroTrustAccessPolicy#login_method}
	LoginMethod *[]*string `field:"optional" json:"loginMethod" yaml:"loginMethod"`
	// okta block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#okta ZeroTrustAccessPolicy#okta}
	Okta interface{} `field:"optional" json:"okta" yaml:"okta"`
	// saml block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#saml ZeroTrustAccessPolicy#saml}
	Saml interface{} `field:"optional" json:"saml" yaml:"saml"`
	// The ID of an Access service token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/zero_trust_access_policy#service_token ZeroTrustAccessPolicy#service_token}
	ServiceToken *[]*string `field:"optional" json:"serviceToken" yaml:"serviceToken"`
}

