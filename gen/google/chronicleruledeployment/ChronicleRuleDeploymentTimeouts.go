package chronicleruledeployment


type ChronicleRuleDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#create ChronicleRuleDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#delete ChronicleRuleDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_rule_deployment#update ChronicleRuleDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

