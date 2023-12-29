//go:build no_runtime_type_checking

package alertpolicy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertPolicyRespondersList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertPolicyRespondersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyRespondersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyRespondersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyRespondersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertPolicyRespondersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertPolicyRespondersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

