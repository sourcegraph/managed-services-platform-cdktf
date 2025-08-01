package accounttoken


type AccountTokenCondition struct {
	// Client IP restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/account_token#request_ip AccountToken#request_ip}
	RequestIp *AccountTokenConditionRequestIp `field:"optional" json:"requestIp" yaml:"requestIp"`
}

