package googlecomputeurlmap


type GoogleComputeUrlMapPathMatcherRouteRulesMatchRulesQueryParameterMatches struct {
	// The name of the query parameter to match.
	//
	// The query parameter must exist in the
	// request, in the absence of which the request match fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#name GoogleComputeUrlMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The queryParameterMatch matches if the value of the parameter exactly matches the contents of exactMatch.
	//
	// Only one of presentMatch, exactMatch and regexMatch
	// must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#exact_match GoogleComputeUrlMap#exact_match}
	ExactMatch *string `field:"optional" json:"exactMatch" yaml:"exactMatch"`
	// Specifies that the queryParameterMatch matches if the request contains the query parameter, irrespective of whether the parameter has a value or not.
	//
	// Only one of
	// presentMatch, exactMatch and regexMatch must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#present_match GoogleComputeUrlMap#present_match}
	PresentMatch interface{} `field:"optional" json:"presentMatch" yaml:"presentMatch"`
	// The queryParameterMatch matches if the value of the parameter matches the regular expression specified by regexMatch.
	//
	// For the regular expression grammar,
	// please see en.cppreference.com/w/cpp/regex/ecmascript  Only one of presentMatch,
	// exactMatch and regexMatch must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_url_map#regex_match GoogleComputeUrlMap#regex_match}
	RegexMatch *string `field:"optional" json:"regexMatch" yaml:"regexMatch"`
}

