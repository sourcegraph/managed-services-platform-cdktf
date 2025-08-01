package workersscript


type WorkersScriptMigrationsStepsTransferredClasses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#from WorkersScript#from}.
	From *string `field:"optional" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#from_script WorkersScript#from_script}.
	FromScript *string `field:"optional" json:"fromScript" yaml:"fromScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_script#to WorkersScript#to}.
	To *string `field:"optional" json:"to" yaml:"to"`
}

