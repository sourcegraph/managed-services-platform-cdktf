//go:build no_runtime_type_checking

package integrationaction

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationActionAcknowledgeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionAcknowledgeList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionAcknowledgeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAcknowledgeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationActionAcknowledgeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

