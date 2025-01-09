package allprojectsspikeprotection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AllProjectsSpikeProtectionConfig struct {
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
	// Toggle the browser-extensions, localhost, filtered-transaction, or web-crawlers filter on or off for all projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/all_projects_spike_protection#enabled AllProjectsSpikeProtection#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/all_projects_spike_protection#organization AllProjectsSpikeProtection#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The slugs of the projects to enable or disable spike protection for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/all_projects_spike_protection#projects AllProjectsSpikeProtection#projects}
	Projects *[]*string `field:"required" json:"projects" yaml:"projects"`
}

