//go:build no_runtime_type_checking

package maintenance

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MaintenanceTimeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MaintenanceTimeList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MaintenanceTimeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MaintenanceTimeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MaintenanceTimeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MaintenanceTimeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MaintenanceTimeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMaintenanceTimeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

