package googlestoragehmackey


type GoogleStorageHmacKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_storage_hmac_key#create GoogleStorageHmacKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_storage_hmac_key#delete GoogleStorageHmacKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_storage_hmac_key#update GoogleStorageHmacKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
