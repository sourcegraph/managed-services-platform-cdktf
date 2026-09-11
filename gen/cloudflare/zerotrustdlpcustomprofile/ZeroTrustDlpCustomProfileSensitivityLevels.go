package zerotrustdlpcustomprofile


type ZeroTrustDlpCustomProfileSensitivityLevels struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_custom_profile#group_id ZeroTrustDlpCustomProfile#group_id}.
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_dlp_custom_profile#level_id ZeroTrustDlpCustomProfile#level_id}.
	LevelId *string `field:"required" json:"levelId" yaml:"levelId"`
}

