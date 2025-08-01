package spectrumapplication


type SpectrumApplicationDns struct {
	// The name of the DNS record associated with the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/spectrum_application#name SpectrumApplication#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The type of DNS record associated with the application. Available values: "CNAME", "ADDRESS".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/spectrum_application#type SpectrumApplication#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

