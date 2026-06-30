package googlefirestorefield

type GoogleFirestoreFieldIndexConfig struct {
	// indexes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_firestore_field#indexes GoogleFirestoreField#indexes}
	Indexes any `field:"optional" json:"indexes" yaml:"indexes"`
}
