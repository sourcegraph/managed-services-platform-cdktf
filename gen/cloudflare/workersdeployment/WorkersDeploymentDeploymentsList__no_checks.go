//go:build no_runtime_type_checking

package workersdeployment

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkersDeploymentDeploymentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentDeploymentsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WorkersDeploymentDeploymentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentDeploymentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentDeploymentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WorkersDeploymentDeploymentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWorkersDeploymentDeploymentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

