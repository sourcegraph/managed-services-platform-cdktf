package aisearchinstance


type AiSearchInstanceCustomMetadata struct {
	// Available values: "text", "number", "boolean".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#data_type AiSearchInstance#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/ai_search_instance#field_name AiSearchInstance#field_name}.
	FieldName *string `field:"required" json:"fieldName" yaml:"fieldName"`
}

