//go:build no_runtime_type_checking

package maintenance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaintenanceRulesEntityList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MaintenanceRulesEntityList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaintenanceRulesEntityList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesEntityList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesEntityList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesEntityList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaintenanceRulesEntityList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaintenanceRulesEntityListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

