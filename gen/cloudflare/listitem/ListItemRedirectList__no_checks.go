//go:build no_runtime_type_checking

package listitem

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_ListItemRedirectList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_ListItemRedirectList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_ListItemRedirectList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ListItemRedirectList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ListItemRedirectList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ListItemRedirectList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ListItemRedirectList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewListItemRedirectListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

