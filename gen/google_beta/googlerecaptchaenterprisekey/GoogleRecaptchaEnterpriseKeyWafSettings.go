package googlerecaptchaenterprisekey


type GoogleRecaptchaEnterpriseKeyWafSettings struct {
	// Supported WAF features. For more information, see https://cloud.google.com/recaptcha-enterprise/docs/usecase#comparison_of_features. Possible values: CHALLENGE_PAGE, SESSION_TOKEN, ACTION_TOKEN, EXPRESS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_recaptcha_enterprise_key#waf_feature GoogleRecaptchaEnterpriseKey#waf_feature}
	WafFeature *string `field:"required" json:"wafFeature" yaml:"wafFeature"`
	// The WAF service that uses this key. Possible values: CA, FASTLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_recaptcha_enterprise_key#waf_service GoogleRecaptchaEnterpriseKey#waf_service}
	WafService *string `field:"required" json:"wafService" yaml:"wafService"`
}

