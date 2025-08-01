package rediscluster


type RedisClusterCrossClusterReplicationConfigSecondaryClusters struct {
	// The full resource path of the secondary cluster in the format: projects/{project}/locations/{region}/clusters/{cluster-id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/redis_cluster#cluster RedisCluster#cluster}
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
}

