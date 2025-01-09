//go:build no_runtime_type_checking

package workspacesettings

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspaceSettings) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateImportFromParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateMoveToIdParameters(id *string) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettings) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateWorkspaceSettings_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateWorkspaceSettings_IsConstructParameters(x interface{}) error {
	return nil
}

func validateWorkspaceSettings_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateWorkspaceSettings_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetAgentPoolIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetExecutionModeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetGlobalRemoteStateParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetRemoteStateConsumerIdsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettings) validateSetWorkspaceIdParameters(val *string) error {
	return nil
}

func validateNewWorkspaceSettingsParameters(scope constructs.Construct, id *string, config *WorkspaceSettingsConfig) error {
	return nil
}

