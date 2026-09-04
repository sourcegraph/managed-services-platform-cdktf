package datacloudflarecertificatepack


type DataCloudflareCertificatePackFilter struct {
	// Include Certificate Packs of all statuses, not just active ones. Available values: "all".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/certificate_pack#status DataCloudflareCertificatePack#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

