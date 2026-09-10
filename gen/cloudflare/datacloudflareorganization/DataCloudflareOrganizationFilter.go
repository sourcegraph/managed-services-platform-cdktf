package datacloudflareorganization


type DataCloudflareOrganizationFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#containing DataCloudflareOrganization#containing}.
	Containing *DataCloudflareOrganizationFilterContaining `field:"optional" json:"containing" yaml:"containing"`
	// Only return organizations with the specified IDs (ex. id=foo&id=bar). Send multiple elements by repeating the query value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#id DataCloudflareOrganization#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *[]*string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#name DataCloudflareOrganization#name}.
	Name *DataCloudflareOrganizationFilterName `field:"optional" json:"name" yaml:"name"`
	// The amount of items to return. Defaults to 10.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#page_size DataCloudflareOrganization#page_size}
	PageSize *float64 `field:"optional" json:"pageSize" yaml:"pageSize"`
	// An opaque token returned from the last list response that when provided will retrieve the next page.
	//
	// Parameters used to filter the retrieved list must remain in subsequent
	// requests with a page token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#page_token DataCloudflareOrganization#page_token}
	PageToken *string `field:"optional" json:"pageToken" yaml:"pageToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/organization#parent DataCloudflareOrganization#parent}.
	Parent *DataCloudflareOrganizationFilterParent `field:"optional" json:"parent" yaml:"parent"`
}

