package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluation struct {
	// The API endpoint containing your business logic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#evaluate_url ZeroTrustAccessApplication#evaluate_url}
	EvaluateUrl *string `field:"required" json:"evaluateUrl" yaml:"evaluateUrl"`
	// The API endpoint containing the key that Access uses to verify that the response came from your API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application#keys_url ZeroTrustAccessApplication#keys_url}
	KeysUrl *string `field:"required" json:"keysUrl" yaml:"keysUrl"`
}

