package discoveryenginedatastore


type DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesOcrParsingConfig struct {
	// If true, will use native text instead of OCR text on pages containing native text.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/discovery_engine_data_store#use_native_text DiscoveryEngineDataStore#use_native_text}
	UseNativeText interface{} `field:"optional" json:"useNativeText" yaml:"useNativeText"`
}

