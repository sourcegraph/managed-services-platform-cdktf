//go:build no_runtime_type_checking

package registryprovider

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegistryProvider) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_RegistryProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateRegistryProvider_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateRegistryProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRegistryProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRegistryProvider_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetNamespaceParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetOrganizationParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_RegistryProvider) validateSetRegistryNameParameters(val *string) error {
	return nil
}

func validateNewRegistryProviderParameters(scope constructs.Construct, id *string, config *RegistryProviderConfig) error {
	return nil
}

