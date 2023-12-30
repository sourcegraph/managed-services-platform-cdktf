//go:build no_runtime_type_checking

package notificationrule

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NotificationRuleCriteriaList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NotificationRuleCriteriaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NotificationRuleCriteriaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NotificationRuleCriteriaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NotificationRuleCriteriaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NotificationRuleCriteriaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNotificationRuleCriteriaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

