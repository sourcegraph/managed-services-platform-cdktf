package chroniclereferencelist


type ChronicleReferenceListEntries struct {
	// Required. The value of the entry. Maximum length is 512 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_reference_list#value ChronicleReferenceList#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

