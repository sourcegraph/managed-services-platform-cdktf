package datacloudflarecertificatepack


type DataCloudflareCertificatePackFilter struct {
	// Specify the deployment environment for the certificate packs. Available values: "staging", "production".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/certificate_pack#deploy DataCloudflareCertificatePack#deploy}
	Deploy *string `field:"optional" json:"deploy" yaml:"deploy"`
	// Include Certificate Packs of all statuses, not just active ones. Available values: "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/certificate_pack#status DataCloudflareCertificatePack#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

