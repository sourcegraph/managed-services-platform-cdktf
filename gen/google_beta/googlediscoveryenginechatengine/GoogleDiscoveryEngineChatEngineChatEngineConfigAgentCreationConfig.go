package googlediscoveryenginechatengine


type GoogleDiscoveryEngineChatEngineChatEngineConfigAgentCreationConfig struct {
	// The default language of the agent as a language tag.
	//
	// See [Language Support](https://cloud.google.com/dialogflow/docs/reference/language) for a list of the currently supported language codes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_chat_engine#default_language_code GoogleDiscoveryEngineChatEngine#default_language_code}
	DefaultLanguageCode *string `field:"required" json:"defaultLanguageCode" yaml:"defaultLanguageCode"`
	// The time zone of the agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York, Europe/Paris.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_chat_engine#time_zone GoogleDiscoveryEngineChatEngine#time_zone}
	TimeZone *string `field:"required" json:"timeZone" yaml:"timeZone"`
	// Name of the company, organization or other entity that the agent represents.
	//
	// Used for knowledge connector LLM prompt and for knowledge search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_chat_engine#business GoogleDiscoveryEngineChatEngine#business}
	Business *string `field:"optional" json:"business" yaml:"business"`
	// Agent location for Agent creation, currently supported values: global/us/eu, it needs to be the same region as the Chat Engine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_discovery_engine_chat_engine#location GoogleDiscoveryEngineChatEngine#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
}

