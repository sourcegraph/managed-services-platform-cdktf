package teamprojectaccess


type TeamProjectAccessProjectAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#settings TeamProjectAccess#settings}.
	Settings *string `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#teams TeamProjectAccess#teams}.
	Teams *string `field:"optional" json:"teams" yaml:"teams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_project_access#variable_sets TeamProjectAccess#variable_sets}.
	VariableSets *string `field:"optional" json:"variableSets" yaml:"variableSets"`
}

