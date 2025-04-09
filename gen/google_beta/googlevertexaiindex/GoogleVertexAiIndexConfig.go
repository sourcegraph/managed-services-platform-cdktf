package googlevertexaiindex

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleVertexAiIndexConfig struct {
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
	// The display name of the Index.
	//
	// The name can be up to 128 characters long and can consist of any UTF-8 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#display_name GoogleVertexAiIndex#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The description of the Index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#description GoogleVertexAiIndex#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#id GoogleVertexAiIndex#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The update method to use with this Index.
	//
	// The value must be the followings. If not set, BATCH_UPDATE will be used by default.
	// * BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
	// * STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#index_update_method GoogleVertexAiIndex#index_update_method}
	IndexUpdateMethod *string `field:"optional" json:"indexUpdateMethod" yaml:"indexUpdateMethod"`
	// The labels with user-defined metadata to organize your Indexes.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#labels GoogleVertexAiIndex#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#metadata GoogleVertexAiIndex#metadata}
	Metadata *GoogleVertexAiIndexMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#project GoogleVertexAiIndex#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the index. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#region GoogleVertexAiIndex#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_vertex_ai_index#timeouts GoogleVertexAiIndex#timeouts}
	Timeouts *GoogleVertexAiIndexTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

