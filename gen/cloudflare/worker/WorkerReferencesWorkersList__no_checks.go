//go:build no_runtime_type_checking

package worker

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerReferencesWorkersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerReferencesWorkersList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerReferencesWorkersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerReferencesWorkersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerReferencesWorkersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerReferencesWorkersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerReferencesWorkersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

