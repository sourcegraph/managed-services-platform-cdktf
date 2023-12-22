//go:build no_runtime_type_checking

package workspacesettings

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkspaceSettingsOverwritesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkspaceSettingsOverwritesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettingsOverwritesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettingsOverwritesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkspaceSettingsOverwritesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkspaceSettingsOverwritesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

