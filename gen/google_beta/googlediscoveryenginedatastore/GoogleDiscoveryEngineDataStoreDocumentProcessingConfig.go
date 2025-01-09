package googlediscoveryenginedatastore


type GoogleDiscoveryEngineDataStoreDocumentProcessingConfig struct {
	// chunking_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_discovery_engine_data_store#chunking_config GoogleDiscoveryEngineDataStore#chunking_config}
	ChunkingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigChunkingConfig `field:"optional" json:"chunkingConfig" yaml:"chunkingConfig"`
	// default_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_discovery_engine_data_store#default_parsing_config GoogleDiscoveryEngineDataStore#default_parsing_config}
	DefaultParsingConfig *GoogleDiscoveryEngineDataStoreDocumentProcessingConfigDefaultParsingConfig `field:"optional" json:"defaultParsingConfig" yaml:"defaultParsingConfig"`
	// parsing_config_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_discovery_engine_data_store#parsing_config_overrides GoogleDiscoveryEngineDataStore#parsing_config_overrides}
	ParsingConfigOverrides interface{} `field:"optional" json:"parsingConfigOverrides" yaml:"parsingConfigOverrides"`
}

