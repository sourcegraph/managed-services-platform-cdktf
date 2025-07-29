package googleapigeesecurityprofilev2


type GoogleApigeeSecurityProfileV2ProfileAssessmentConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apigee_security_profile_v2#assessment GoogleApigeeSecurityProfileV2#assessment}.
	Assessment *string `field:"required" json:"assessment" yaml:"assessment"`
	// The weight of the assessment. Possible values: ["MINOR", "MODERATE", "MAJOR"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_apigee_security_profile_v2#weight GoogleApigeeSecurityProfileV2#weight}
	Weight *string `field:"required" json:"weight" yaml:"weight"`
}

