package googlenetworkservicesedgecacheservice


type GoogleNetworkServicesEdgeCacheServiceRoutingPathMatcherRouteRuleUrlRedirect struct {
	// The host that will be used in the redirect response instead of the one that was supplied in the request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#host_redirect GoogleNetworkServicesEdgeCacheService#host_redirect}
	HostRedirect *string `field:"optional" json:"hostRedirect" yaml:"hostRedirect"`
	// If set to true, the URL scheme in the redirected request is set to https.
	//
	// If set to false, the URL scheme of the redirected request will remain the same as that of the request.
	//
	// This can only be set if there is at least one (1) edgeSslCertificate set on the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#https_redirect GoogleNetworkServicesEdgeCacheService#https_redirect}
	HttpsRedirect interface{} `field:"optional" json:"httpsRedirect" yaml:"httpsRedirect"`
	// The path that will be used in the redirect response instead of the one that was supplied in the request.
	//
	// pathRedirect cannot be supplied together with prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.
	//
	// The path value must be between 1 and 1024 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#path_redirect GoogleNetworkServicesEdgeCacheService#path_redirect}
	PathRedirect *string `field:"optional" json:"pathRedirect" yaml:"pathRedirect"`
	// The prefix that replaces the prefixMatch specified in the routeRule, retaining the remaining portion of the URL before redirecting the request.
	//
	// prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or neither. If neither is supplied, the path of the original request will be used for the redirect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#prefix_redirect GoogleNetworkServicesEdgeCacheService#prefix_redirect}
	PrefixRedirect *string `field:"optional" json:"prefixRedirect" yaml:"prefixRedirect"`
	// The HTTP Status code to use for this RedirectAction.
	//
	// The supported values are:
	//
	// - 'MOVED_PERMANENTLY_DEFAULT', which is the default value and corresponds to 301.
	// - 'FOUND', which corresponds to 302.
	// - 'SEE_OTHER' which corresponds to 303.
	// - 'TEMPORARY_REDIRECT', which corresponds to 307. in this case, the request method will be retained.
	// - 'PERMANENT_REDIRECT', which corresponds to 308. in this case, the request method will be retained. Possible values: ["MOVED_PERMANENTLY_DEFAULT", "FOUND", "SEE_OTHER", "TEMPORARY_REDIRECT", "PERMANENT_REDIRECT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#redirect_response_code GoogleNetworkServicesEdgeCacheService#redirect_response_code}
	RedirectResponseCode *string `field:"optional" json:"redirectResponseCode" yaml:"redirectResponseCode"`
	// If set to true, any accompanying query portion of the original URL is removed prior to redirecting the request.
	//
	// If set to false, the query portion of the original URL is retained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_services_edge_cache_service#strip_query GoogleNetworkServicesEdgeCacheService#strip_query}
	StripQuery interface{} `field:"optional" json:"stripQuery" yaml:"stripQuery"`
}

