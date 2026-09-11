package workflow


type WorkflowLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/workflow#steps Workflow#steps}.
	Steps *float64 `field:"optional" json:"steps" yaml:"steps"`
}

