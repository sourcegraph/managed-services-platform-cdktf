//go:build no_runtime_type_checking

package escalation

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EscalationRulesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EscalationRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EscalationRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EscalationRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEscalationRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

