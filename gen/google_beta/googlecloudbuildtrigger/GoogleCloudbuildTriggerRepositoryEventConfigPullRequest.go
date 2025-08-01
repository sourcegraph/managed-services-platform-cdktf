package googlecloudbuildtrigger


type GoogleCloudbuildTriggerRepositoryEventConfigPullRequest struct {
	// Regex of branches to match.
	//
	// The syntax of the regular expressions accepted is the syntax accepted by
	// RE2 and described at https://github.com/google/re2/wiki/Syntax
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#branch GoogleCloudbuildTrigger#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Configure builds to run whether a repository owner or collaborator need to comment '/gcbrun'. Possible values: ["COMMENTS_DISABLED", "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#comment_control GoogleCloudbuildTrigger#comment_control}
	CommentControl *string `field:"optional" json:"commentControl" yaml:"commentControl"`
	// If true, branches that do NOT match the git_ref will trigger a build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#invert_regex GoogleCloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
}

