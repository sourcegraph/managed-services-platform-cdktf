package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsAntivirus struct {
	// Specify whether to enable anti-virus scanning on downloads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#enabled_download_phase ZeroTrustGatewaySettings#enabled_download_phase}
	EnabledDownloadPhase interface{} `field:"optional" json:"enabledDownloadPhase" yaml:"enabledDownloadPhase"`
	// Specify whether to enable anti-virus scanning on uploads.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#enabled_upload_phase ZeroTrustGatewaySettings#enabled_upload_phase}
	EnabledUploadPhase interface{} `field:"optional" json:"enabledUploadPhase" yaml:"enabledUploadPhase"`
	// Specify whether to block requests for unscannable files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#fail_closed ZeroTrustGatewaySettings#fail_closed}
	FailClosed interface{} `field:"optional" json:"failClosed" yaml:"failClosed"`
	// Configure the message the user's device shows during an antivirus scan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#notification_settings ZeroTrustGatewaySettings#notification_settings}
	NotificationSettings *ZeroTrustGatewaySettingsSettingsAntivirusNotificationSettings `field:"optional" json:"notificationSettings" yaml:"notificationSettings"`
}

