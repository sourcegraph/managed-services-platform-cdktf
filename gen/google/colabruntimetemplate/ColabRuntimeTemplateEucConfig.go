package colabruntimetemplate


type ColabRuntimeTemplateEucConfig struct {
	// Disable end user credential access for the runtime.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_runtime_template#euc_disabled ColabRuntimeTemplate#euc_disabled}
	EucDisabled interface{} `field:"optional" json:"eucDisabled" yaml:"eucDisabled"`
}

