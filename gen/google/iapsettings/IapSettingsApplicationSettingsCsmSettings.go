package iapsettings


type IapSettingsApplicationSettingsCsmSettings struct {
	// Audience claim set in the generated RCToken. This value is not validated by IAP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iap_settings#rctoken_aud IapSettings#rctoken_aud}
	RctokenAud *string `field:"optional" json:"rctokenAud" yaml:"rctokenAud"`
}

