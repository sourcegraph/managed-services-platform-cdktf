//go:build no_runtime_type_checking

package escalation

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EscalationRepeatList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EscalationRepeatList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EscalationRepeatList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EscalationRepeatList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EscalationRepeatList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EscalationRepeatList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEscalationRepeatListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

