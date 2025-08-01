package assuredworkloadsworkload


type AssuredWorkloadsWorkloadWorkloadOptions struct {
	// Indicates type of KAJ enrollment for the workload.
	//
	// Currently, only specifiying KEY_ACCESS_TRANSPARENCY_OFF is implemented to not enroll in KAT-level KAJ enrollment for Regional Controls workloads. Possible values: KAJ_ENROLLMENT_TYPE_UNSPECIFIED, FULL_KAJ, EKM_ONLY, KEY_ACCESS_TRANSPARENCY_OFF
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/assured_workloads_workload#kaj_enrollment_type AssuredWorkloadsWorkload#kaj_enrollment_type}
	KajEnrollmentType *string `field:"optional" json:"kajEnrollmentType" yaml:"kajEnrollmentType"`
}

