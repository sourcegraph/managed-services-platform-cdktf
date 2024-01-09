//go:build no_runtime_type_checking

package metricalert

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MetricAlertTriggerList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MetricAlertTriggerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MetricAlertTriggerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMetricAlertTriggerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

