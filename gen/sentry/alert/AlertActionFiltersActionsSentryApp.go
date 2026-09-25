package alert


type AlertActionFiltersActionsSentryApp struct {
	// The numeric Sentry App ID. Use `tostring(data.sentry_app_installation.<name>.sentry_app_id)` to source this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#sentry_app_id Alert#sentry_app_id}
	SentryAppId *string `field:"required" json:"sentryAppId" yaml:"sentryAppId"`
	// Key-value settings passed to the Sentry App action.
	//
	// Specifying `label` preserves the human-readable display name in the Sentry UI for async select fields whose options are paginated by the third-party app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#settings Alert#settings}
	Settings interface{} `field:"optional" json:"settings" yaml:"settings"`
}

