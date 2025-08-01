package rediscluster


type RedisClusterPscConfigs struct {
	// Required. The consumer network where the network address of the discovery endpoint will be reserved, in the form of projects/{network_project_id_or_number}/global/networks/{network_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/redis_cluster#network RedisCluster#network}
	Network *string `field:"required" json:"network" yaml:"network"`
}

