//go:build no_runtime_type_checking

package registrymodule

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegistryModuleTestConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RegistryModuleTestConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RegistryModuleTestConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryModuleTestConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryModuleTestConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RegistryModuleTestConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRegistryModuleTestConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

