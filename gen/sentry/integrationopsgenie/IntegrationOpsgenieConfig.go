package integrationopsgenie

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationOpsgenieConfig struct {
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
	// The ID of the Opsgenie integration. Source from the URL `https://<organization>.sentry.io/settings/integrations/opsgenie/<integration-id>/` or use the `sentry_organization_integration` data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/integration_opsgenie#integration_id IntegrationOpsgenie#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The integration key of the Opsgenie service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/integration_opsgenie#integration_key IntegrationOpsgenie#integration_key}
	IntegrationKey *string `field:"required" json:"integrationKey" yaml:"integrationKey"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/integration_opsgenie#organization IntegrationOpsgenie#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The name of the Opsgenie team. In Sentry, this is called Label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/integration_opsgenie#team IntegrationOpsgenie#team}
	Team *string `field:"required" json:"team" yaml:"team"`
}

