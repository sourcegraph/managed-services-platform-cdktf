package healthcareconsentstore


type HealthcareConsentStoreTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/healthcare_consent_store#create HealthcareConsentStore#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/healthcare_consent_store#delete HealthcareConsentStore#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/healthcare_consent_store#update HealthcareConsentStore#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
