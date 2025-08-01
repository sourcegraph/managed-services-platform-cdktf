package vertexaifeatureonlinestorefeatureview


type VertexAiFeatureOnlineStoreFeatureviewFeatureRegistrySourceFeatureGroups struct {
	// Identifier of the feature group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_feature_online_store_featureview#feature_group_id VertexAiFeatureOnlineStoreFeatureview#feature_group_id}
	FeatureGroupId *string `field:"required" json:"featureGroupId" yaml:"featureGroupId"`
	// Identifiers of features under the feature group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vertex_ai_feature_online_store_featureview#feature_ids VertexAiFeatureOnlineStoreFeatureview#feature_ids}
	FeatureIds *[]*string `field:"required" json:"featureIds" yaml:"featureIds"`
}

