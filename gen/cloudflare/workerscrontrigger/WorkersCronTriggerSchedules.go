package workerscrontrigger


type WorkersCronTriggerSchedules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/workers_cron_trigger#cron WorkersCronTrigger#cron}.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
}

