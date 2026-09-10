package datacloudflareorganization


type DataCloudflareOrganizationFilterContaining struct {
	// Filter the list of organizations to the ones that contain this particular account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#account DataCloudflareOrganization#account}
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Filter the list of organizations to the ones that contain this particular organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#organization DataCloudflareOrganization#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Filter the list of organizations to the ones that contain this particular user.
	//
	// IMPORTANT: Just because an organization "contains" a user is not a
	// representation of any authorization or privilege to manage any resources
	// therein. An organization "containing" a user simply means the user is managed by
	// that organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#user DataCloudflareOrganization#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

