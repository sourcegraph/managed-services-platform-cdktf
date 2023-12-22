package googlemonitoringslo


type GoogleMonitoringSloRequestBasedSli struct {
	// distribution_cut block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_monitoring_slo#distribution_cut GoogleMonitoringSlo#distribution_cut}
	DistributionCut *GoogleMonitoringSloRequestBasedSliDistributionCut `field:"optional" json:"distributionCut" yaml:"distributionCut"`
	// good_total_ratio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.10.0/docs/resources/google_monitoring_slo#good_total_ratio GoogleMonitoringSlo#good_total_ratio}
	GoodTotalRatio *GoogleMonitoringSloRequestBasedSliGoodTotalRatio `field:"optional" json:"goodTotalRatio" yaml:"goodTotalRatio"`
}

