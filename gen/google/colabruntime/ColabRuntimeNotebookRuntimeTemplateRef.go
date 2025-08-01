package colabruntime


type ColabRuntimeNotebookRuntimeTemplateRef struct {
	// The resource name of the NotebookRuntimeTemplate based on which a NotebookRuntime will be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime#notebook_runtime_template ColabRuntime#notebook_runtime_template}
	NotebookRuntimeTemplate *string `field:"required" json:"notebookRuntimeTemplate" yaml:"notebookRuntimeTemplate"`
}

