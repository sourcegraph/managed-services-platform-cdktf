package googlegkeonprembaremetaladmincluster


type GoogleGkeonpremBareMetalAdminClusterNodeAccessConfig struct {
	// LoginUser is the user name used to access node machines. It defaults to "root" if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_gkeonprem_bare_metal_admin_cluster#login_user GoogleGkeonpremBareMetalAdminCluster#login_user}
	LoginUser *string `field:"optional" json:"loginUser" yaml:"loginUser"`
}

