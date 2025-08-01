package firestoreindex


type FirestoreIndexTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firestore_index#create FirestoreIndex#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firestore_index#delete FirestoreIndex#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

