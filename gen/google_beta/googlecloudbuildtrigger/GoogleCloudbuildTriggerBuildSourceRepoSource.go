package googlecloudbuildtrigger


type GoogleCloudbuildTriggerBuildSourceRepoSource struct {
	// Name of the Cloud Source Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#repo_name GoogleCloudbuildTrigger#repo_name}
	RepoName *string `field:"required" json:"repoName" yaml:"repoName"`
	// Regex matching branches to build.
	//
	// Exactly one a of branch name, tag, or commit SHA must be provided.
	// The syntax of the regular expressions accepted is the syntax accepted by RE2 and
	// described at https://github.com/google/re2/wiki/Syntax
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#branch_name GoogleCloudbuildTrigger#branch_name}
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
	// Explicit commit SHA to build. Exactly one a of branch name, tag, or commit SHA must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#commit_sha GoogleCloudbuildTrigger#commit_sha}
	CommitSha *string `field:"optional" json:"commitSha" yaml:"commitSha"`
	// Directory, relative to the source root, in which to run the build.
	//
	// This must be a relative path. If a step's dir is specified and is an absolute path,
	// this value is ignored for that step's execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#dir GoogleCloudbuildTrigger#dir}
	Dir *string `field:"optional" json:"dir" yaml:"dir"`
	// Only trigger a build if the revision regex does NOT match the revision regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#invert_regex GoogleCloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
	// ID of the project that owns the Cloud Source Repository. If omitted, the project ID requesting the build is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#project_id GoogleCloudbuildTrigger#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Substitutions to use in a triggered build. Should only be used with triggers.run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#substitutions GoogleCloudbuildTrigger#substitutions}
	Substitutions *map[string]*string `field:"optional" json:"substitutions" yaml:"substitutions"`
	// Regex matching tags to build.
	//
	// Exactly one a of branch name, tag, or commit SHA must be provided.
	// The syntax of the regular expressions accepted is the syntax accepted by RE2 and
	// described at https://github.com/google/re2/wiki/Syntax
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#tag_name GoogleCloudbuildTrigger#tag_name}
	TagName *string `field:"optional" json:"tagName" yaml:"tagName"`
}

