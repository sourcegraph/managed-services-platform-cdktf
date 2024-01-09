package projectspikeprotection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectSpikeProtectionConfig struct {
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
	// Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/project_spike_protection#enabled ProjectSpikeProtection#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The slug of the organization the project belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/project_spike_protection#organization ProjectSpikeProtection#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The slug of the project to create the filter for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.12.1/docs/resources/project_spike_protection#project ProjectSpikeProtection#project}
	Project *string `field:"required" json:"project" yaml:"project"`
}

