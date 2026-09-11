//go:build no_runtime_type_checking

package usergroupmembers

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_UserGroupMembersMembersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (u *jsiiProxy_UserGroupMembersMembersList) validateGetParameters(index *float64) error {
	return nil
}

func (u *jsiiProxy_UserGroupMembersMembersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_UserGroupMembersMembersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_UserGroupMembersMembersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_UserGroupMembersMembersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_UserGroupMembersMembersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewUserGroupMembersMembersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

