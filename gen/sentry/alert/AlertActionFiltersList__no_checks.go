//go:build no_runtime_type_checking

package alert

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertActionFiltersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertActionFiltersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertActionFiltersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertActionFiltersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertActionFiltersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertActionFiltersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertActionFiltersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertActionFiltersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

