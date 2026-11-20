package teamprojectaccess

type TeamProjectAccessWorkspaceAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#create TeamProjectAccess#create}.
	Create any `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#delete TeamProjectAccess#delete}.
	Delete any `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#locking TeamProjectAccess#locking}.
	Locking any `field:"optional" json:"locking" yaml:"locking"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#move TeamProjectAccess#move}.
	Move any `field:"optional" json:"move" yaml:"move"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#runs TeamProjectAccess#runs}.
	Runs *string `field:"optional" json:"runs" yaml:"runs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#run_tasks TeamProjectAccess#run_tasks}.
	RunTasks any `field:"optional" json:"runTasks" yaml:"runTasks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#sentinel_mocks TeamProjectAccess#sentinel_mocks}.
	SentinelMocks *string `field:"optional" json:"sentinelMocks" yaml:"sentinelMocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#state_versions TeamProjectAccess#state_versions}.
	StateVersions *string `field:"optional" json:"stateVersions" yaml:"stateVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#variables TeamProjectAccess#variables}.
	Variables *string `field:"optional" json:"variables" yaml:"variables"`
}
