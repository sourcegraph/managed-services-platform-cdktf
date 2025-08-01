package googlecloudbuildtrigger


type GoogleCloudbuildTriggerRepositoryEventConfigPush struct {
	// Regex of branches to match.
	//
	// The syntax of the regular expressions accepted is the syntax accepted by
	// RE2 and described at https://github.com/google/re2/wiki/Syntax
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#branch GoogleCloudbuildTrigger#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// If true, only trigger a build if the revision regex does NOT match the git_ref regex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#invert_regex GoogleCloudbuildTrigger#invert_regex}
	InvertRegex interface{} `field:"optional" json:"invertRegex" yaml:"invertRegex"`
	// Regex of tags to match.
	//
	// The syntax of the regular expressions accepted is the syntax accepted by
	// RE2 and described at https://github.com/google/re2/wiki/Syntax
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuild_trigger#tag GoogleCloudbuildTrigger#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

