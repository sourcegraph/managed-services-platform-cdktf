package gkeonprembaremetalcluster


type GkeonpremBareMetalClusterNodeAccessConfig struct {
	// LoginUser is the user name used to access node machines. It defaults to "root" if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/gkeonprem_bare_metal_cluster#login_user GkeonpremBareMetalCluster#login_user}
	LoginUser *string `field:"optional" json:"loginUser" yaml:"loginUser"`
}

