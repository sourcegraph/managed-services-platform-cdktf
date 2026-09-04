package organizationprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationProfileAConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/organization_profile#business_address OrganizationProfileA#business_address}.
	BusinessAddress *string `field:"required" json:"businessAddress" yaml:"businessAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/organization_profile#business_email OrganizationProfileA#business_email}.
	BusinessEmail *string `field:"required" json:"businessEmail" yaml:"businessEmail"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/organization_profile#business_name OrganizationProfileA#business_name}.
	BusinessName *string `field:"required" json:"businessName" yaml:"businessName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/organization_profile#business_phone OrganizationProfileA#business_phone}.
	BusinessPhone *string `field:"required" json:"businessPhone" yaml:"businessPhone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/organization_profile#external_metadata OrganizationProfileA#external_metadata}.
	ExternalMetadata *string `field:"required" json:"externalMetadata" yaml:"externalMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/organization_profile#organization_id OrganizationProfileA#organization_id}.
	OrganizationId *string `field:"required" json:"organizationId" yaml:"organizationId"`
}

