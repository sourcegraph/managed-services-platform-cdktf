package alert


type AlertActionFiltersActionsPagerduty struct {
	// The ID of the PagerDuty integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The ID of the PagerDuty service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#service_id Alert#service_id}
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// The name of the service to create the ticket in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#service_name Alert#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
	// The PagerDuty severity level for the notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#severity Alert#severity}
	Severity *string `field:"required" json:"severity" yaml:"severity"`
}

