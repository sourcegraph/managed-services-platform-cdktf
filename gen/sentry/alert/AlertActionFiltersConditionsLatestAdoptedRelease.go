package alert


type AlertActionFiltersConditionsLatestAdoptedRelease struct {
	// The age comparison to use. Valid values are: `older`, and `newer`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#age_comparison Alert#age_comparison}
	AgeComparison *string `field:"required" json:"ageComparison" yaml:"ageComparison"`
	// The environment to compare against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#environment Alert#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// The release age type to use. Valid values are: `oldest`, and `newest`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#release_age_type Alert#release_age_type}
	ReleaseAgeType *string `field:"required" json:"releaseAgeType" yaml:"releaseAgeType"`
}

