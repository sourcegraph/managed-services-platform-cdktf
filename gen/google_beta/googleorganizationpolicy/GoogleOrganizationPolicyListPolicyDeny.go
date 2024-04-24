package googleorganizationpolicy


type GoogleOrganizationPolicyListPolicyDeny struct {
	// The policy allows or denies all values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_organization_policy#all GoogleOrganizationPolicy#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// The policy can define specific values that are allowed or denied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_organization_policy#values GoogleOrganizationPolicy#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

