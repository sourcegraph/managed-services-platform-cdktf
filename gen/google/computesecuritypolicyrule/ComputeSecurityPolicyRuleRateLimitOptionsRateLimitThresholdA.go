package computesecuritypolicyrule


type ComputeSecurityPolicyRuleRateLimitOptionsRateLimitThresholdA struct {
	// Number of HTTP(S) requests for calculating the threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_security_policy_rule#count ComputeSecurityPolicyRuleA#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Interval over which the threshold is computed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_security_policy_rule#interval_sec ComputeSecurityPolicyRuleA#interval_sec}
	IntervalSec *float64 `field:"optional" json:"intervalSec" yaml:"intervalSec"`
}

