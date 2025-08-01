package datacloudflareloadbalancerpool


type DataCloudflareLoadBalancerPoolFilter struct {
	// The ID of the Monitor to use for checking the health of origins within this pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/load_balancer_pool#monitor DataCloudflareLoadBalancerPool#monitor}
	Monitor *string `field:"optional" json:"monitor" yaml:"monitor"`
}

