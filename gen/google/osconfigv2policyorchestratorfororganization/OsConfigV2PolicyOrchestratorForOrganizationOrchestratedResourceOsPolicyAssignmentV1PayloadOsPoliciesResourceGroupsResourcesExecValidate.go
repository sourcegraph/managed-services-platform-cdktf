package osconfigv2policyorchestratorfororganization


type OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecValidate struct {
	// Required. The script interpreter to use. Possible values: INTERPRETER_UNSPECIFIED NONE SHELL POWERSHELL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#interpreter OsConfigV2PolicyOrchestratorForOrganization#interpreter}
	Interpreter *string `field:"required" json:"interpreter" yaml:"interpreter"`
	// Optional arguments to pass to the source during execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#args OsConfigV2PolicyOrchestratorForOrganization#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#file OsConfigV2PolicyOrchestratorForOrganization#file}
	File *OsConfigV2PolicyOrchestratorForOrganizationOrchestratedResourceOsPolicyAssignmentV1PayloadOsPoliciesResourceGroupsResourcesExecValidateFile `field:"optional" json:"file" yaml:"file"`
	// Only recorded for enforce Exec.
	//
	// Path to an output file (that is created by this Exec) whose
	// content will be recorded in OSPolicyResourceCompliance after a
	// successful run. Absence or failure to read this file will result in
	// this ExecResource being non-compliant. Output file size is limited to
	// 500K bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#output_file_path OsConfigV2PolicyOrchestratorForOrganization#output_file_path}
	OutputFilePath *string `field:"optional" json:"outputFilePath" yaml:"outputFilePath"`
	// An inline script. The size of the script is limited to 32KiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/os_config_v2_policy_orchestrator_for_organization#script OsConfigV2PolicyOrchestratorForOrganization#script}
	Script *string `field:"optional" json:"script" yaml:"script"`
}

