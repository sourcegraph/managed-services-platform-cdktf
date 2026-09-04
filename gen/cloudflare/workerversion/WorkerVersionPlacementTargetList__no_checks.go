//go:build no_runtime_type_checking

package workerversion

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkerVersionPlacementTargetList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkerVersionPlacementTargetList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkerVersionPlacementTargetList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionPlacementTargetList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionPlacementTargetList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionPlacementTargetList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkerVersionPlacementTargetList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkerVersionPlacementTargetListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

