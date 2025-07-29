package googlevertexaifeatureonlinestorefeatureviewiampolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_policy#feature_online_store GoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy#feature_online_store}.
	FeatureOnlineStore *string `field:"required" json:"featureOnlineStore" yaml:"featureOnlineStore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_policy#feature_view GoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy#feature_view}.
	FeatureView *string `field:"required" json:"featureView" yaml:"featureView"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_policy#policy_data GoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy#policy_data}.
	PolicyData *string `field:"required" json:"policyData" yaml:"policyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_policy#id GoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_policy#project GoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_feature_online_store_featureview_iam_policy#region GoogleVertexAiFeatureOnlineStoreFeatureviewIamPolicy#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

