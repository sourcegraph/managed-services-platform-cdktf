package googlediscoveryenginesearchengine


type GoogleDiscoveryEngineSearchEngineCommonConfig struct {
	// The name of the company, business or entity that is associated with the engine.
	//
	// Setting this may help improve LLM related features.cd
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_discovery_engine_search_engine#company_name GoogleDiscoveryEngineSearchEngine#company_name}
	CompanyName *string `field:"optional" json:"companyName" yaml:"companyName"`
}

