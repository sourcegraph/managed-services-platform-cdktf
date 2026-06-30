package googlecomputeimage

type GoogleComputeImageShieldedInstanceInitialState struct {
	// dbs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_image#dbs GoogleComputeImage#dbs}
	Dbs any `field:"optional" json:"dbs" yaml:"dbs"`
	// dbxs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_image#dbxs GoogleComputeImage#dbxs}
	Dbxs any `field:"optional" json:"dbxs" yaml:"dbxs"`
	// keks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_image#keks GoogleComputeImage#keks}
	Keks any `field:"optional" json:"keks" yaml:"keks"`
	// pk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_image#pk GoogleComputeImage#pk}
	Pk *GoogleComputeImageShieldedInstanceInitialStatePk `field:"optional" json:"pk" yaml:"pk"`
}
