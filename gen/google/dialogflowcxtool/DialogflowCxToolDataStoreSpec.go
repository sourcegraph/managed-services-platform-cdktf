package dialogflowcxtool


type DialogflowCxToolDataStoreSpec struct {
	// data_store_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_tool#data_store_connections DialogflowCxTool#data_store_connections}
	DataStoreConnections interface{} `field:"required" json:"dataStoreConnections" yaml:"dataStoreConnections"`
	// fallback_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dialogflow_cx_tool#fallback_prompt DialogflowCxTool#fallback_prompt}
	FallbackPrompt *DialogflowCxToolDataStoreSpecFallbackPrompt `field:"required" json:"fallbackPrompt" yaml:"fallbackPrompt"`
}

