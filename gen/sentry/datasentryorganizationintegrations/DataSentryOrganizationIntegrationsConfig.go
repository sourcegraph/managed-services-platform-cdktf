package datasentryorganizationintegrations

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataSentryOrganizationIntegrationsConfig struct {
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
	// The organization the resource belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/data-sources/organization_integrations#organization DataSentryOrganizationIntegrations#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// Specific integration provider to filter by such as `slack`. See [the list of supported providers](https://docs.sentry.io/integrations/).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/data-sources/organization_integrations#provider_key DataSentryOrganizationIntegrations#provider_key}
	ProviderKey *string `field:"optional" json:"providerKey" yaml:"providerKey"`
}

