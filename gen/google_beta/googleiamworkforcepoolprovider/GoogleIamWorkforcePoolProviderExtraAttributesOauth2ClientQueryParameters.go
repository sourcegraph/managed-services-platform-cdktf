package googleiamworkforcepoolprovider


type GoogleIamWorkforcePoolProviderExtraAttributesOauth2ClientQueryParameters struct {
	// The filter used to request specific records from IdP.
	//
	// In case of attributes type as AZURE_AD_GROUPS_MAIL and AZURE_AD_GROUPS_ID, it represents the
	// filter used to request specific groups for users from IdP. By default, all of the groups associated with the user are fetched. The
	// groups should be security enabled. See https://learn.microsoft.com/en-us/graph/search-query-parameter for more details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_iam_workforce_pool_provider#filter GoogleIamWorkforcePoolProvider#filter}
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
}

