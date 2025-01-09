//go:build no_runtime_type_checking

package datatferegistryproviders

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfeRegistryProvidersProvidersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfeRegistryProvidersProvidersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfeRegistryProvidersProvidersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfeRegistryProvidersProvidersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfeRegistryProvidersProvidersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfeRegistryProvidersProvidersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfeRegistryProvidersProvidersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

