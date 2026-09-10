package datacloudflareorganization


type DataCloudflareOrganizationFilterParent struct {
	// Filter the list of organizations to the ones that are a sub-organization of the specified organization.
	//
	// "null" is a valid value to provide for this parameter. It means "where
	// an organization has no parent (i.e. it is a 'root' organization)."
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#id DataCloudflareOrganization#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

