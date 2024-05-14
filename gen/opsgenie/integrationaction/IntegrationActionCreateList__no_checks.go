//go:build no_runtime_type_checking

package integrationaction

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationActionCreateList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionCreateList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionCreateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionCreateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationActionCreateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

