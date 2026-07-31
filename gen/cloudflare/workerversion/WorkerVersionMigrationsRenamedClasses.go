package workerversion


type WorkerVersionMigrationsRenamedClasses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/worker_version#from WorkerVersion#from}.
	From *string `field:"optional" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/worker_version#to WorkerVersion#to}.
	To *string `field:"optional" json:"to" yaml:"to"`
}

