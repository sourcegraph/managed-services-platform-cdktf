package googlenotebooksruntime


type GoogleNotebooksRuntimeVirtualMachine struct {
	// virtual_machine_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_notebooks_runtime#virtual_machine_config GoogleNotebooksRuntime#virtual_machine_config}
	VirtualMachineConfig *GoogleNotebooksRuntimeVirtualMachineVirtualMachineConfig `field:"optional" json:"virtualMachineConfig" yaml:"virtualMachineConfig"`
}

