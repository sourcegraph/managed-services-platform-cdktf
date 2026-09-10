//go:build no_runtime_type_checking

package usergroup

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserGroupPoliciesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserGroupPoliciesList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserGroupPoliciesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserGroupPoliciesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserGroupPoliciesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

