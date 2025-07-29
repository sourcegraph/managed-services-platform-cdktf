package googlenotebooksenvironment


type GoogleNotebooksEnvironmentContainerImage struct {
	// The path to the container image repository. For example: gcr.io/{project_id}/{imageName}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_environment#repository GoogleNotebooksEnvironment#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The tag of the container image. If not specified, this defaults to the latest tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_environment#tag GoogleNotebooksEnvironment#tag}
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
}

