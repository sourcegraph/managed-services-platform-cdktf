package alloydbcluster


type AlloydbClusterPscConfig struct {
	// Create an instance that allows connections from Private Service Connect endpoints to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/alloydb_cluster#psc_enabled AlloydbCluster#psc_enabled}
	PscEnabled interface{} `field:"optional" json:"pscEnabled" yaml:"pscEnabled"`
}

