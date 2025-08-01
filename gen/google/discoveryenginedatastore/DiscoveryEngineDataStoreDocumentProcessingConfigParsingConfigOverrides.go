package discoveryenginedatastore


type DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/discovery_engine_data_store#file_type DiscoveryEngineDataStore#file_type}.
	FileType *string `field:"required" json:"fileType" yaml:"fileType"`
	// digital_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/discovery_engine_data_store#digital_parsing_config DiscoveryEngineDataStore#digital_parsing_config}
	DigitalParsingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesDigitalParsingConfig `field:"optional" json:"digitalParsingConfig" yaml:"digitalParsingConfig"`
	// layout_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/discovery_engine_data_store#layout_parsing_config DiscoveryEngineDataStore#layout_parsing_config}
	LayoutParsingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesLayoutParsingConfig `field:"optional" json:"layoutParsingConfig" yaml:"layoutParsingConfig"`
	// ocr_parsing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/discovery_engine_data_store#ocr_parsing_config DiscoveryEngineDataStore#ocr_parsing_config}
	OcrParsingConfig *DiscoveryEngineDataStoreDocumentProcessingConfigParsingConfigOverridesOcrParsingConfig `field:"optional" json:"ocrParsingConfig" yaml:"ocrParsingConfig"`
}

