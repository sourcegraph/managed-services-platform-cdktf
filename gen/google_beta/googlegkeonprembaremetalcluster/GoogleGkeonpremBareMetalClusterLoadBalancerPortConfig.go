package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancerPortConfig struct {
	// The port that control plane hosted load balancers will listen on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gkeonprem_bare_metal_cluster#control_plane_load_balancer_port GoogleGkeonpremBareMetalCluster#control_plane_load_balancer_port}
	ControlPlaneLoadBalancerPort *float64 `field:"required" json:"controlPlaneLoadBalancerPort" yaml:"controlPlaneLoadBalancerPort"`
}

