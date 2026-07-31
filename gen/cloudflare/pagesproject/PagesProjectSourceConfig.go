package pagesproject


type PagesProjectSourceConfig struct {
	// Whether to enable automatic deployments when pushing to the source repository.
	//
	// When disabled, no deployments (production or preview) will be triggered automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#deployments_enabled PagesProject#deployments_enabled}
	DeploymentsEnabled interface{} `field:"optional" json:"deploymentsEnabled" yaml:"deploymentsEnabled"`
	// The owner of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#owner PagesProject#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// The owner ID of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#owner_id PagesProject#owner_id}
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// A list of paths that should be excluded from triggering a preview deployment. Wildcard syntax (`*`) is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#path_excludes PagesProject#path_excludes}
	PathExcludes *[]*string `field:"optional" json:"pathExcludes" yaml:"pathExcludes"`
	// A list of paths that should be watched to trigger a preview deployment. Wildcard syntax (`*`) is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#path_includes PagesProject#path_includes}
	PathIncludes *[]*string `field:"optional" json:"pathIncludes" yaml:"pathIncludes"`
	// Whether to enable PR comments.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#pr_comments_enabled PagesProject#pr_comments_enabled}
	PrCommentsEnabled interface{} `field:"optional" json:"prCommentsEnabled" yaml:"prCommentsEnabled"`
	// A list of branches that should not trigger a preview deployment.
	//
	// Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#preview_branch_excludes PagesProject#preview_branch_excludes}
	PreviewBranchExcludes *[]*string `field:"optional" json:"previewBranchExcludes" yaml:"previewBranchExcludes"`
	// A list of branches that should trigger a preview deployment.
	//
	// Wildcard syntax (`*`) is supported. Must be used with `preview_deployment_setting` set to `custom`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#preview_branch_includes PagesProject#preview_branch_includes}
	PreviewBranchIncludes *[]*string `field:"optional" json:"previewBranchIncludes" yaml:"previewBranchIncludes"`
	// Controls whether commits to preview branches trigger a preview deployment. Available values: "all", "none", "custom".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#preview_deployment_setting PagesProject#preview_deployment_setting}
	PreviewDeploymentSetting *string `field:"optional" json:"previewDeploymentSetting" yaml:"previewDeploymentSetting"`
	// The production branch of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#production_branch PagesProject#production_branch}
	ProductionBranch *string `field:"optional" json:"productionBranch" yaml:"productionBranch"`
	// Whether to trigger a production deployment on commits to the production branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#production_deployments_enabled PagesProject#production_deployments_enabled}
	ProductionDeploymentsEnabled interface{} `field:"optional" json:"productionDeploymentsEnabled" yaml:"productionDeploymentsEnabled"`
	// The ID of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#repo_id PagesProject#repo_id}
	RepoId *string `field:"optional" json:"repoId" yaml:"repoId"`
	// The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/pages_project#repo_name PagesProject#repo_name}
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
}

