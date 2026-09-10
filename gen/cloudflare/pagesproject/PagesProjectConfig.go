package pagesproject

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PagesProjectConfig struct {
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
	// Name of the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pages_project#name PagesProject#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Production branch of the project. Used to identify production deployments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pages_project#production_branch PagesProject#production_branch}
	ProductionBranch *string `field:"required" json:"productionBranch" yaml:"productionBranch"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pages_project#account_id PagesProject#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Configs for the project build process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pages_project#build_config PagesProject#build_config}
	BuildConfig *PagesProjectBuildConfig `field:"optional" json:"buildConfig" yaml:"buildConfig"`
	// Configs for deployments in a project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pages_project#deployment_configs PagesProject#deployment_configs}
	DeploymentConfigs *PagesProjectDeploymentConfigs `field:"optional" json:"deploymentConfigs" yaml:"deploymentConfigs"`
	// Configs for the project source control.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/pages_project#source PagesProject#source}
	Source *PagesProjectSource `field:"optional" json:"source" yaml:"source"`
}

