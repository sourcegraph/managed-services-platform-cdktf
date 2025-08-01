package apitoken


type ApiTokenCondition struct {
	// Client IP restrictions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/api_token#request_ip ApiToken#request_ip}
	RequestIp *ApiTokenConditionRequestIp `field:"optional" json:"requestIp" yaml:"requestIp"`
}

