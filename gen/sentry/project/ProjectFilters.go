package project


type ProjectFilters struct {
	// Filter events from these IP addresses. (e.g. 127.0.0.1 or 10.0.0.0/8).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/project#blacklisted_ips Project#blacklisted_ips}
	BlacklistedIps *[]*string `field:"optional" json:"blacklistedIps" yaml:"blacklistedIps"`
	// Filter events by error messages. Allows [glob pattern matching](https://en.wikipedia.org/wiki/Glob_(programming)). (e.g. TypeError* or *: integer division or modulo by zero).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/project#error_messages Project#error_messages}
	ErrorMessages *[]*string `field:"optional" json:"errorMessages" yaml:"errorMessages"`
	// Filter events from these releases. Allows [glob pattern matching](https://en.wikipedia.org/wiki/Glob_(programming)). (e.g. 1.* or [!3].[0-9].*).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/project#releases Project#releases}
	Releases *[]*string `field:"optional" json:"releases" yaml:"releases"`
}

