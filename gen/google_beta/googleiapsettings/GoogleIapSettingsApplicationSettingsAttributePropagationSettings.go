package googleiapsettings


type GoogleIapSettingsApplicationSettingsAttributePropagationSettings struct {
	// Whether the provided attribute propagation settings should be evaluated on user requests.
	//
	// If set to true, attributes returned from the expression will be propagated in the set output credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_iap_settings#enable GoogleIapSettings#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Raw string CEL expression.
	//
	// Must return a list of attributes. A maximum of 45 attributes can
	// be selected. Expressions can select different attribute types from attributes:
	// attributes.saml_attributes, attributes.iap_attributes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_iap_settings#expression GoogleIapSettings#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// Which output credentials attributes selected by the CEL expression should be propagated in.
	//
	// All attributes will be fully duplicated in each selected output credential.
	// Possible values are:
	//
	// * 'HEADER': Propagate attributes in the headers with "x-goog-iap-attr-" prefix.
	// * 'JWT': Propagate attributes in the JWT of the form:
	//          "additional_claims": { "my_attribute": ["value1", "value2"] }
	// * 'RCTOKEN': Propagate attributes in the RCToken of the form: "
	//              additional_claims": { "my_attribute": ["value1", "value2"] } Possible values: ["HEADER", "JWT", "RCTOKEN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_iap_settings#output_credentials GoogleIapSettings#output_credentials}
	OutputCredentials *[]*string `field:"optional" json:"outputCredentials" yaml:"outputCredentials"`
}

