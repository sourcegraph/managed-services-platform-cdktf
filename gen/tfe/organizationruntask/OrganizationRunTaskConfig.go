package organizationruntask

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRunTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#name OrganizationRunTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#url OrganizationRunTask#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#category OrganizationRunTask#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#description OrganizationRunTask#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#enabled OrganizationRunTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#hmac_key OrganizationRunTask#hmac_key}.
	HmacKey *string `field:"optional" json:"hmacKey" yaml:"hmacKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#id OrganizationRunTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/organization_run_task#organization OrganizationRunTask#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
}

