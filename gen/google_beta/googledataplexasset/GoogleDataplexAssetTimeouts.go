package googledataplexasset


type GoogleDataplexAssetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dataplex_asset#create GoogleDataplexAsset#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dataplex_asset#delete GoogleDataplexAsset#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_dataplex_asset#update GoogleDataplexAsset#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
