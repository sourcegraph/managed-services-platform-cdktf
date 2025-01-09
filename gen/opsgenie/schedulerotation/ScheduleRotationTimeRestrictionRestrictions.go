package schedulerotation


type ScheduleRotationTimeRestrictionRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#end_day ScheduleRotation#end_day}.
	EndDay *string `field:"required" json:"endDay" yaml:"endDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#end_hour ScheduleRotation#end_hour}.
	EndHour *float64 `field:"required" json:"endHour" yaml:"endHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#end_min ScheduleRotation#end_min}.
	EndMin *float64 `field:"required" json:"endMin" yaml:"endMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#start_day ScheduleRotation#start_day}.
	StartDay *string `field:"required" json:"startDay" yaml:"startDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#start_hour ScheduleRotation#start_hour}.
	StartHour *float64 `field:"required" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#start_min ScheduleRotation#start_min}.
	StartMin *float64 `field:"required" json:"startMin" yaml:"startMin"`
}

