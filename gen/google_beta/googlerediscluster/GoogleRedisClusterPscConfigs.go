package googlerediscluster


type GoogleRedisClusterPscConfigs struct {
	// Required. The consumer network where the network address of the discovery endpoint will be reserved, in the form of projects/{network_project_id_or_number}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_redis_cluster#network GoogleRedisCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}
