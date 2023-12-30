//go:build no_runtime_type_checking

package integrationaction

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationActionCreateFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionCreateFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationActionCreateFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

