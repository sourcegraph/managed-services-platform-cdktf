package googledocumentaiprocessordefaultversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDocumentAiProcessorDefaultVersionConfig struct {
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
	// The processor to set the version on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_document_ai_processor_default_version#processor GoogleDocumentAiProcessorDefaultVersion#processor}
	Processor *string `field:"required" json:"processor" yaml:"processor"`
	// The version to set.
	//
	// Using 'stable' or 'rc' will cause the API to return the latest version in that release channel.
	// Apply 'lifecycle.ignore_changes' to the 'version' field to suppress this diff.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_document_ai_processor_default_version#version GoogleDocumentAiProcessorDefaultVersion#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_document_ai_processor_default_version#id GoogleDocumentAiProcessorDefaultVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.29.0/docs/resources/google_document_ai_processor_default_version#timeouts GoogleDocumentAiProcessorDefaultVersion#timeouts}
	Timeouts *GoogleDocumentAiProcessorDefaultVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

