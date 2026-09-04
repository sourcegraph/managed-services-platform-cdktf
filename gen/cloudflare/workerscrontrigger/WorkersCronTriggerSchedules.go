package workerscrontrigger


type WorkersCronTriggerSchedules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_cron_trigger#cron WorkersCronTrigger#cron}.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
}

