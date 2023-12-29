//go:build no_runtime_type_checking

package schedulerotation

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduleRotationParticipantList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_ScheduleRotationParticipantList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ScheduleRotationParticipantList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ScheduleRotationParticipantList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ScheduleRotationParticipantList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ScheduleRotationParticipantList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewScheduleRotationParticipantListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

