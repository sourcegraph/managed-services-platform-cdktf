package googlestoragecontrolorganizationintelligenceconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleStorageControlOrganizationIntelligenceConfigConfig struct {
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
	// Identifier of the GCP Organization. For GCP org, this field should be organization number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_control_organization_intelligence_config#name GoogleStorageControlOrganizationIntelligenceConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Edition configuration of the Storage Intelligence resource. Valid values are INHERIT, DISABLED, TRIAL and STANDARD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_control_organization_intelligence_config#edition_config GoogleStorageControlOrganizationIntelligenceConfig#edition_config}
	EditionConfig *string `field:"optional" json:"editionConfig" yaml:"editionConfig"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_control_organization_intelligence_config#filter GoogleStorageControlOrganizationIntelligenceConfig#filter}
	Filter *GoogleStorageControlOrganizationIntelligenceConfigFilter `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_control_organization_intelligence_config#id GoogleStorageControlOrganizationIntelligenceConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_storage_control_organization_intelligence_config#timeouts GoogleStorageControlOrganizationIntelligenceConfig#timeouts}
	Timeouts *GoogleStorageControlOrganizationIntelligenceConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

