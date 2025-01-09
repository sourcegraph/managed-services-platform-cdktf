package team


type TeamOrganizationAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#access_secret_teams Team#access_secret_teams}.
	AccessSecretTeams interface{} `field:"optional" json:"accessSecretTeams" yaml:"accessSecretTeams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_agent_pools Team#manage_agent_pools}.
	ManageAgentPools interface{} `field:"optional" json:"manageAgentPools" yaml:"manageAgentPools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_membership Team#manage_membership}.
	ManageMembership interface{} `field:"optional" json:"manageMembership" yaml:"manageMembership"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_modules Team#manage_modules}.
	ManageModules interface{} `field:"optional" json:"manageModules" yaml:"manageModules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_organization_access Team#manage_organization_access}.
	ManageOrganizationAccess interface{} `field:"optional" json:"manageOrganizationAccess" yaml:"manageOrganizationAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_policies Team#manage_policies}.
	ManagePolicies interface{} `field:"optional" json:"managePolicies" yaml:"managePolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_policy_overrides Team#manage_policy_overrides}.
	ManagePolicyOverrides interface{} `field:"optional" json:"managePolicyOverrides" yaml:"managePolicyOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_projects Team#manage_projects}.
	ManageProjects interface{} `field:"optional" json:"manageProjects" yaml:"manageProjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_providers Team#manage_providers}.
	ManageProviders interface{} `field:"optional" json:"manageProviders" yaml:"manageProviders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_run_tasks Team#manage_run_tasks}.
	ManageRunTasks interface{} `field:"optional" json:"manageRunTasks" yaml:"manageRunTasks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_teams Team#manage_teams}.
	ManageTeams interface{} `field:"optional" json:"manageTeams" yaml:"manageTeams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_vcs_settings Team#manage_vcs_settings}.
	ManageVcsSettings interface{} `field:"optional" json:"manageVcsSettings" yaml:"manageVcsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#manage_workspaces Team#manage_workspaces}.
	ManageWorkspaces interface{} `field:"optional" json:"manageWorkspaces" yaml:"manageWorkspaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#read_projects Team#read_projects}.
	ReadProjects interface{} `field:"optional" json:"readProjects" yaml:"readProjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/team#read_workspaces Team#read_workspaces}.
	ReadWorkspaces interface{} `field:"optional" json:"readWorkspaces" yaml:"readWorkspaces"`
}

