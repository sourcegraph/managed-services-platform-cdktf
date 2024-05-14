//go:build no_runtime_type_checking

package metricalert

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MetricAlertTriggerActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (m *jsiiProxy_MetricAlertTriggerActionList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MetricAlertTriggerActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMetricAlertTriggerActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

