package googleprivatecacertificateauthority


type GooglePrivatecaCertificateAuthorityConfigSubjectConfigSubject struct {
	// The common name of the distinguished name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#common_name GooglePrivatecaCertificateAuthority#common_name}
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// The country code of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#country_code GooglePrivatecaCertificateAuthority#country_code}
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// The locality or city of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#locality GooglePrivatecaCertificateAuthority#locality}
	Locality *string `field:"optional" json:"locality" yaml:"locality"`
	// The organization of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#organization GooglePrivatecaCertificateAuthority#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The organizational unit of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#organizational_unit GooglePrivatecaCertificateAuthority#organizational_unit}
	OrganizationalUnit *string `field:"optional" json:"organizationalUnit" yaml:"organizationalUnit"`
	// The postal code of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#postal_code GooglePrivatecaCertificateAuthority#postal_code}
	PostalCode *string `field:"optional" json:"postalCode" yaml:"postalCode"`
	// The province, territory, or regional state of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#province GooglePrivatecaCertificateAuthority#province}
	Province *string `field:"optional" json:"province" yaml:"province"`
	// The street address of the subject.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_certificate_authority#street_address GooglePrivatecaCertificateAuthority#street_address}
	StreetAddress *string `field:"optional" json:"streetAddress" yaml:"streetAddress"`
}

