package zerotrustgatewaypolicy


type ZeroTrustGatewayPolicyRuleSettingsRedirect struct {
	// Specify the URI to which the user is redirected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#target_uri ZeroTrustGatewayPolicy#target_uri}
	TargetUri *string `field:"required" json:"targetUri" yaml:"targetUri"`
	// Specify whether to pass the context information as query parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#include_context ZeroTrustGatewayPolicy#include_context}
	IncludeContext interface{} `field:"optional" json:"includeContext" yaml:"includeContext"`
	// Specify whether to append the path and query parameters from the original request to target_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_policy#preserve_path_and_query ZeroTrustGatewayPolicy#preserve_path_and_query}
	PreservePathAndQuery interface{} `field:"optional" json:"preservePathAndQuery" yaml:"preservePathAndQuery"`
}

