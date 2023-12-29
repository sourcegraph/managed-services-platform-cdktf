//go:build no_runtime_type_checking

package emailintegration

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmailIntegrationRespondersList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmailIntegrationRespondersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmailIntegrationRespondersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EmailIntegrationRespondersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmailIntegrationRespondersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmailIntegrationRespondersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmailIntegrationRespondersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

