//go:build no_runtime_type_checking

package integrationaction

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationActionAcknowledgeFilterConditionsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionAcknowledgeFilterConditionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeFilterConditionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeFilterConditionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeFilterConditionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeFilterConditionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationActionAcknowledgeFilterConditionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

