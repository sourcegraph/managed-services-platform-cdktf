package iamoauthclient


type IamOauthClientTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iam_oauth_client#create IamOauthClient#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iam_oauth_client#delete IamOauthClient#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/iam_oauth_client#update IamOauthClient#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

