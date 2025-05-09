package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterLoadBalancerBgpLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfigs struct {
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//
	// These will added in addition to any default label(s)
	// that Kubernetes may apply to the node. In case of conflict in
	// label keys, the applied set may differ depending on the Kubernetes
	// version -- it's best to assume the behavior is undefined and
	// conflicts should be avoided. For more information, including usage
	// and the valid values, see:
	//   - http://kubernetes.io/v1.1/docs/user-guide/labels.html
	// An object containing a list of "key": value pairs.
	// For example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gkeonprem_bare_metal_cluster#labels GoogleGkeonpremBareMetalCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The default IPv4 address for SSH access and Kubernetes node. Example: 192.168.0.1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_gkeonprem_bare_metal_cluster#node_ip GoogleGkeonpremBareMetalCluster#node_ip}
	NodeIp *string `field:"optional" json:"nodeIp" yaml:"nodeIp"`
}

