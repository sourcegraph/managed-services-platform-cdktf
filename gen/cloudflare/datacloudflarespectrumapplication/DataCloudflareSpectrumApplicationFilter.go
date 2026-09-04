package datacloudflarespectrumapplication


type DataCloudflareSpectrumApplicationFilter struct {
	// Sets the direction by which results are ordered. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/spectrum_application#direction DataCloudflareSpectrumApplication#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Application field by which results are ordered. Available values: "protocol", "app_id", "created_on", "modified_on", "dns".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/spectrum_application#order DataCloudflareSpectrumApplication#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
}

