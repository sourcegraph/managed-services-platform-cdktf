//go:build no_runtime_type_checking

package alert

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlertTriggerConditionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlertTriggerConditionsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlertTriggerConditionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlertTriggerConditionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlertTriggerConditionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlertTriggerConditionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlertTriggerConditionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlertTriggerConditionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

