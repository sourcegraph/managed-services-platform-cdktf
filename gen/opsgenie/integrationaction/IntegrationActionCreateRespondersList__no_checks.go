//go:build no_runtime_type_checking

package integrationaction

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationActionCreateRespondersList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionCreateRespondersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateRespondersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateRespondersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateRespondersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateRespondersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationActionCreateRespondersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

