//go:build no_runtime_type_checking

package worker

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerReferencesDomainsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerReferencesDomainsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerReferencesDomainsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerReferencesDomainsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerReferencesDomainsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerReferencesDomainsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerReferencesDomainsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

