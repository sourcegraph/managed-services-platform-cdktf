package issuealert


type IssueAlertFiltersV2EventAttribute struct {
	// Valid values are: `message`, `platform`, `environment`, `type`, `error.handled`, `error.unhandled`, `error.main_thread`, `exception.type`, `exception.value`, `user.id`, `user.email`, `user.username`, `user.ip_address`, `http.method`, `http.url`, `http.status_code`, `sdk.name`, `stacktrace.code`, `stacktrace.module`, `stacktrace.filename`, `stacktrace.abs_path`, `stacktrace.package`, `unreal.crash_type`, `app.in_foreground`, `os.distribution_name`, `os.distribution_version`, and `symbolicated_in_app`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#attribute IssueAlert#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// The comparison operator.
	//
	// Valid values are: `CONTAINS`, `ENDS_WITH`, `EQUAL`, `GREATER_OR_EQUAL`, `GREATER`, `IS_SET`, `IS_IN`, `LESS_OR_EQUAL`, `LESS`, `NOT_CONTAINS`, `NOT_ENDS_WITH`, `NOT_EQUAL`, `NOT_SET`, `NOT_STARTS_WITH`, `NOT_IN`, and `STARTS_WITH`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#match IssueAlert#match}
	Match *string `field:"required" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/issue_alert#value IssueAlert#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

