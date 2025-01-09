//go:build no_runtime_type_checking

package datasentryallkeys

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataSentryAllKeysKeysList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataSentryAllKeysKeysList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataSentryAllKeysKeysList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataSentryAllKeysKeysList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataSentryAllKeysKeysList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataSentryAllKeysKeysList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataSentryAllKeysKeysListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

