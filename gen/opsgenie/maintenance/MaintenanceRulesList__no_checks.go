//go:build no_runtime_type_checking

package maintenance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaintenanceRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaintenanceRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaintenanceRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

