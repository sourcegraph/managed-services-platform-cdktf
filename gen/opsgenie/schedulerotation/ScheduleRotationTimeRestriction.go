package schedulerotation


type ScheduleRotationTimeRestriction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#type ScheduleRotation#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#restriction ScheduleRotation#restriction}
	Restriction *ScheduleRotationTimeRestrictionRestriction `field:"optional" json:"restriction" yaml:"restriction"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs/resources/schedule_rotation#restrictions ScheduleRotation#restrictions}
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
}

