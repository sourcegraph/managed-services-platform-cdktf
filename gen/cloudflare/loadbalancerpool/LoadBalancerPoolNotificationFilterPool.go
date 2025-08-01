package loadbalancerpool


type LoadBalancerPoolNotificationFilterPool struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/load_balancer_pool#disable LoadBalancerPool#disable}
	Disable interface{} `field:"optional" json:"disable" yaml:"disable"`
	// If present, send notifications only for this health status (e.g. false for only DOWN events). Use null to reset (all events).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/load_balancer_pool#healthy LoadBalancerPool#healthy}
	Healthy interface{} `field:"optional" json:"healthy" yaml:"healthy"`
}

