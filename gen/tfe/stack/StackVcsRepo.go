package stack


type StackVcsRepo struct {
	// Identifier of the VCS repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#identifier Stack#identifier}
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// The repository branch that Terraform should use. This defaults to the respository's default branch (e.g. main).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#branch Stack#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The installation ID of the GitHub App.
	//
	// This conflicts with `oauth_token_id` and can only be used if `oauth_token_id` is not used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#github_app_installation_id Stack#github_app_installation_id}
	GithubAppInstallationId *string `field:"optional" json:"githubAppInstallationId" yaml:"githubAppInstallationId"`
	// The VCS Connection to use.
	//
	// This ID can be obtained from a `tfe_oauth_client` resource. This conflicts with `github_app_installation_id` and can only be used if `github_app_installation_id` is not used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/stack#oauth_token_id Stack#oauth_token_id}
	OauthTokenId *string `field:"optional" json:"oauthTokenId" yaml:"oauthTokenId"`
}

