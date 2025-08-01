package workflowsworkflow


type WorkflowsWorkflowTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workflows_workflow#create WorkflowsWorkflow#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workflows_workflow#delete WorkflowsWorkflow#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workflows_workflow#update WorkflowsWorkflow#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

