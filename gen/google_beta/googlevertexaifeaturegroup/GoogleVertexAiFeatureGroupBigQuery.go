package googlevertexaifeaturegroup


type GoogleVertexAiFeatureGroupBigQuery struct {
	// big_query_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_vertex_ai_feature_group#big_query_source GoogleVertexAiFeatureGroup#big_query_source}
	BigQuerySource *GoogleVertexAiFeatureGroupBigQueryBigQuerySource `field:"required" json:"bigQuerySource" yaml:"bigQuerySource"`
	// Columns to construct entityId / row keys. Currently only supports 1 entity_id_column. If not provided defaults to entityId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_vertex_ai_feature_group#entity_id_columns GoogleVertexAiFeatureGroup#entity_id_columns}
	EntityIdColumns *[]*string `field:"optional" json:"entityIdColumns" yaml:"entityIdColumns"`
}

