//go:build no_runtime_type_checking

package accounttoken

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccountTokenPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AccountTokenPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccountTokenPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccountTokenPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccountTokenPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

