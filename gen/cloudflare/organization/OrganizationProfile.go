package organization


type OrganizationProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/organization#business_address Organization#business_address}.
	BusinessAddress *string `field:"required" json:"businessAddress" yaml:"businessAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/organization#business_email Organization#business_email}.
	BusinessEmail *string `field:"required" json:"businessEmail" yaml:"businessEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/organization#business_name Organization#business_name}.
	BusinessName *string `field:"required" json:"businessName" yaml:"businessName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/organization#business_phone Organization#business_phone}.
	BusinessPhone *string `field:"required" json:"businessPhone" yaml:"businessPhone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/organization#external_metadata Organization#external_metadata}.
	ExternalMetadata *string `field:"required" json:"externalMetadata" yaml:"externalMetadata"`
}

