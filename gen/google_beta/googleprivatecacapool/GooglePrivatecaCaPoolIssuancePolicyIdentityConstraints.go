package googleprivatecacapool


type GooglePrivatecaCaPoolIssuancePolicyIdentityConstraints struct {
	// If this is set, the SubjectAltNames extension may be copied from a certificate request into the signed certificate.
	//
	// Otherwise, the requested SubjectAltNames will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#allow_subject_alt_names_passthrough GooglePrivatecaCaPool#allow_subject_alt_names_passthrough}
	AllowSubjectAltNamesPassthrough interface{} `field:"required" json:"allowSubjectAltNamesPassthrough" yaml:"allowSubjectAltNamesPassthrough"`
	// If this is set, the Subject field may be copied from a certificate request into the signed certificate.
	//
	// Otherwise, the requested Subject will be discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#allow_subject_passthrough GooglePrivatecaCaPool#allow_subject_passthrough}
	AllowSubjectPassthrough interface{} `field:"required" json:"allowSubjectPassthrough" yaml:"allowSubjectPassthrough"`
	// cel_expression block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_privateca_ca_pool#cel_expression GooglePrivatecaCaPool#cel_expression}
	CelExpression *GooglePrivatecaCaPoolIssuancePolicyIdentityConstraintsCelExpression `field:"optional" json:"celExpression" yaml:"celExpression"`
}

