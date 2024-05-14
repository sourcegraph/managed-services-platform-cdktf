//go:build no_runtime_type_checking

package team

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TeamMemberList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TeamMemberList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TeamMemberList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TeamMemberList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TeamMemberList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TeamMemberList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TeamMemberList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTeamMemberListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

