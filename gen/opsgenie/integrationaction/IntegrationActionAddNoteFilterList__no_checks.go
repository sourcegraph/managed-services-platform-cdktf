//go:build no_runtime_type_checking

package integrationaction

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IntegrationActionAddNoteFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IntegrationActionAddNoteFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAddNoteFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAddNoteFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAddNoteFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IntegrationActionAddNoteFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIntegrationActionAddNoteFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

