package alert


type AlertActionFiltersActionsGithub struct {
	// The ID of the GitHub integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#integration_id Alert#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The name of the repository to create the issue in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#repo Alert#repo}
	Repo *string `field:"required" json:"repo" yaml:"repo"`
	// The GitHub user to assign the issue to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#assignee Alert#assignee}
	Assignee *string `field:"optional" json:"assignee" yaml:"assignee"`
	// A list of labels to assign to the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#labels Alert#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
}

