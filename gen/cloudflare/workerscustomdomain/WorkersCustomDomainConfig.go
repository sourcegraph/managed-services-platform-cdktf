package workerscustomdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkersCustomDomainConfig struct {
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
	// Hostname of the domain.
	//
	// Can be either the zone apex or a subdomain of the zone. Requests to this hostname will be routed to the configured Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_custom_domain#hostname WorkersCustomDomain#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// Name of the Worker associated with the domain. Requests to the configured hostname will be routed to this Worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_custom_domain#service WorkersCustomDomain#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_custom_domain#account_id WorkersCustomDomain#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Worker environment associated with the domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_custom_domain#environment WorkersCustomDomain#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// ID of the zone containing the domain hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_custom_domain#zone_id WorkersCustomDomain#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
	// Name of the zone containing the domain hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workers_custom_domain#zone_name WorkersCustomDomain#zone_name}
	ZoneName *string `field:"optional" json:"zoneName" yaml:"zoneName"`
}

