//go:build no_runtime_type_checking

package apishield

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiShieldMessagesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ApiShieldMessagesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ApiShieldMessagesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ApiShieldMessagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ApiShieldMessagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ApiShieldMessagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewApiShieldMessagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

