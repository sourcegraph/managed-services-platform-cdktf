package workerversion


type WorkerVersionPlacementTarget struct {
	// TCP host:port for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#host WorkerVersion#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// HTTP hostname for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#hostname WorkerVersion#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Cloud region in format 'provider:region'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#region WorkerVersion#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

