package managedkafkacluster


type ManagedKafkaClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/managed_kafka_cluster#create ManagedKafkaCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/managed_kafka_cluster#delete ManagedKafkaCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/managed_kafka_cluster#update ManagedKafkaCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

