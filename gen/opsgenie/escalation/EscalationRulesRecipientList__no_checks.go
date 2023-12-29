//go:build no_runtime_type_checking

package escalation

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EscalationRulesRecipientList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EscalationRulesRecipientList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesRecipientList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesRecipientList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesRecipientList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesRecipientList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEscalationRulesRecipientListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

