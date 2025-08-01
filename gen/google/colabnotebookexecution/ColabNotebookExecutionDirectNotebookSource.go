package colabnotebookexecution


type ColabNotebookExecutionDirectNotebookSource struct {
	// The base64-encoded contents of the input notebook file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_notebook_execution#content ColabNotebookExecution#content}
	Content *string `field:"required" json:"content" yaml:"content"`
}

