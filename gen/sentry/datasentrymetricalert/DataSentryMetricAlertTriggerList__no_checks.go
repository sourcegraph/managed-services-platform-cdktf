//go:build no_runtime_type_checking

package datasentrymetricalert

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSentryMetricAlertTriggerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSentryMetricAlertTriggerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSentryMetricAlertTriggerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSentryMetricAlertTriggerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSentryMetricAlertTriggerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSentryMetricAlertTriggerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
