package datacloudflareusergroup


type DataCloudflareUserGroupFilter struct {
	// The sort order of returned user groups by name (ascending or descending). Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/user_group#direction DataCloudflareUserGroup#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// A string used for searching for user groups containing that substring.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/user_group#fuzzy_name DataCloudflareUserGroup#fuzzy_name}
	FuzzyName *string `field:"optional" json:"fuzzyName" yaml:"fuzzyName"`
	// ID of the user group to be fetched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/user_group#id DataCloudflareUserGroup#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the user group to be fetched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/user_group#name DataCloudflareUserGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

