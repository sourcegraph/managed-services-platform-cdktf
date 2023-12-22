//go:build no_runtime_type_checking

package nocodemodule

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NoCodeModuleVariableOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NoCodeModuleVariableOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NoCodeModuleVariableOptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NoCodeModuleVariableOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NoCodeModuleVariableOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NoCodeModuleVariableOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNoCodeModuleVariableOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

