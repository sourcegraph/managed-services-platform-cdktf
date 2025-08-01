package healthcareworkspace


type HealthcareWorkspaceSettings struct {
	// Project IDs for data projects hosted in a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/healthcare_workspace#data_project_ids HealthcareWorkspace#data_project_ids}
	DataProjectIds *[]*string `field:"required" json:"dataProjectIds" yaml:"dataProjectIds"`
}

