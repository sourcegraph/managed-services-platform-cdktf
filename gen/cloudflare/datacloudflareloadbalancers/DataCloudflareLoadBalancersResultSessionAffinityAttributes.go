package datacloudflareloadbalancers


type DataCloudflareLoadBalancersResultSessionAffinityAttributes struct {
	// Configures the drain duration in seconds.
	//
	// This field is only used when session affinity is enabled on the load balancer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/load_balancers#drain_duration DataCloudflareLoadBalancers#drain_duration}
	DrainDuration *float64 `field:"optional" json:"drainDuration" yaml:"drainDuration"`
}

