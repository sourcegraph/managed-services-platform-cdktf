package accessgroup


type AccessGroupExcludeExternalEvaluation struct {
	// The API endpoint containing your business logic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_group#evaluate_url AccessGroup#evaluate_url}
	EvaluateUrl *string `field:"optional" json:"evaluateUrl" yaml:"evaluateUrl"`
	// The API endpoint containing the key that Access uses to verify that the response came from your API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.50.0/docs/resources/access_group#keys_url AccessGroup#keys_url}
	KeysUrl *string `field:"optional" json:"keysUrl" yaml:"keysUrl"`
}

