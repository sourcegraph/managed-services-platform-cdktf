package googletagstagkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleTagsTagKeyConfig struct {
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
	// Input only. The resource name of the new TagKey's parent. Must be of the form organizations/{org_id} or projects/{project_id_or_number}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_tags_tag_key#parent GoogleTagsTagKey#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Input only.
	//
	// The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace.
	//
	// The short name can have a maximum length of 256 characters. The permitted character set for the shortName includes all UTF-8 encoded Unicode characters except single quotes ('), double quotes ("), backslashes (\\), and forward slashes (/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_tags_tag_key#short_name GoogleTagsTagKey#short_name}
	ShortName *string `field:"required" json:"shortName" yaml:"shortName"`
	// User-assigned description of the TagKey. Must not exceed 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_tags_tag_key#description GoogleTagsTagKey#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_tags_tag_key#id GoogleTagsTagKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. A purpose cannot be changed once set.
	//
	// A purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag. Possible values: ["GCE_FIREWALL"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_tags_tag_key#purpose GoogleTagsTagKey#purpose}
	Purpose *string `field:"optional" json:"purpose" yaml:"purpose"`
	// Optional. Purpose data cannot be changed once set.
	//
	// Purpose data corresponds to the policy system that the tag is intended for. For example, the GCE_FIREWALL purpose expects data in the following format: 'network = "<project-name>/<vpc-name>"'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_tags_tag_key#purpose_data GoogleTagsTagKey#purpose_data}
	PurposeData *map[string]*string `field:"optional" json:"purposeData" yaml:"purposeData"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_tags_tag_key#timeouts GoogleTagsTagKey#timeouts}
	Timeouts *GoogleTagsTagKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

