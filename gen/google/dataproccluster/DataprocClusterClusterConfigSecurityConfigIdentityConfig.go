package dataproccluster


type DataprocClusterClusterConfigSecurityConfigIdentityConfig struct {
	// User to service account mappings for multi-tenancy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_cluster#user_service_account_mapping DataprocCluster#user_service_account_mapping}
	UserServiceAccountMapping *map[string]*string `field:"required" json:"userServiceAccountMapping" yaml:"userServiceAccountMapping"`
}

