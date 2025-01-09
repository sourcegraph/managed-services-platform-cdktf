//go:build no_runtime_type_checking

package stack

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Stack) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateImportFromParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (s *jsiiProxy_Stack) validateMoveToIdParameters(id *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (s *jsiiProxy_Stack) validatePutVcsRepoParameters(value *StackVcsRepo) error {
	return nil
}

func validateStack_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateStack_IsConstructParameters(x interface{}) error {
	return nil
}

func validateStack_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateStack_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Stack) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Stack) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Stack) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stack) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Stack) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stack) validateSetProjectIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Stack) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func validateNewStackParameters(scope constructs.Construct, id *string, config *StackConfig) error {
	return nil
}

