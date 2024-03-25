package googledatacatalogentrygroupiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataCatalogEntryGroupIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_catalog_entry_group_iam_member#entry_group GoogleDataCatalogEntryGroupIamMember#entry_group}.
	EntryGroup *string `field:"required" json:"entryGroup" yaml:"entryGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_catalog_entry_group_iam_member#member GoogleDataCatalogEntryGroupIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_catalog_entry_group_iam_member#role GoogleDataCatalogEntryGroupIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_catalog_entry_group_iam_member#condition GoogleDataCatalogEntryGroupIamMember#condition}
	Condition *GoogleDataCatalogEntryGroupIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_catalog_entry_group_iam_member#id GoogleDataCatalogEntryGroupIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_catalog_entry_group_iam_member#project GoogleDataCatalogEntryGroupIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.21.0/docs/resources/google_data_catalog_entry_group_iam_member#region GoogleDataCatalogEntryGroupIamMember#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

