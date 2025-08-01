package alloydbcluster


type AlloydbClusterSecondaryConfig struct {
	// Name of the primary cluster must be in the format 'projects/{project}/locations/{location}/clusters/{cluster_id}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/alloydb_cluster#primary_cluster_name AlloydbCluster#primary_cluster_name}
	PrimaryClusterName *string `field:"required" json:"primaryClusterName" yaml:"primaryClusterName"`
}

