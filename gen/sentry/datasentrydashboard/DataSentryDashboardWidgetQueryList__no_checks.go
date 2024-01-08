//go:build no_runtime_type_checking

package datasentrydashboard

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSentryDashboardWidgetQueryList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSentryDashboardWidgetQueryList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSentryDashboardWidgetQueryList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSentryDashboardWidgetQueryList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSentryDashboardWidgetQueryList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSentryDashboardWidgetQueryListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

