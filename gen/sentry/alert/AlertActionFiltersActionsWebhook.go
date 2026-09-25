package alert


type AlertActionFiltersActionsWebhook struct {
	// The slug of the integration service to notify. Use the special value `webhooks` to notify all legacy plugin webhooks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#service Alert#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

