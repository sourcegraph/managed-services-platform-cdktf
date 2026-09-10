package workerversion


type WorkerVersionPlacement struct {
	// TCP host and port for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#host WorkerVersion#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// HTTP hostname for targeted placement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#hostname WorkerVersion#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Enables [Smart Placement](https://developers.cloudflare.com/workers/configuration/smart-placement). Available values: "smart", "targeted".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#mode WorkerVersion#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Cloud region for targeted placement in format 'provider:region'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#region WorkerVersion#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Array of placement targets (currently limited to single target).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/worker_version#target WorkerVersion#target}
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

