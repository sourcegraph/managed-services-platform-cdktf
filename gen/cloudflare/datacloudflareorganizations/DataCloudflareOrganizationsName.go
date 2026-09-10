package datacloudflareorganizations


type DataCloudflareOrganizationsName struct {
	// (case-insensitive) Filter the list of organizations to where the name contains a particular string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#contains DataCloudflareOrganizations#contains}
	Contains *string `field:"optional" json:"contains" yaml:"contains"`
	// (case-insensitive) Filter the list of organizations to where the name ends with a particular string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#ends_with DataCloudflareOrganizations#ends_with}
	EndsWith *string `field:"optional" json:"endsWith" yaml:"endsWith"`
	// (case-insensitive) Filter the list of organizations to where the name starts with a particular string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organizations#starts_with DataCloudflareOrganizations#starts_with}
	StartsWith *string `field:"optional" json:"startsWith" yaml:"startsWith"`
}

