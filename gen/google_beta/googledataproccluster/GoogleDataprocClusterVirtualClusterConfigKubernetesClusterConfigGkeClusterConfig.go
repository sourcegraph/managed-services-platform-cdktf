package googledataproccluster


type GoogleDataprocClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig struct {
	// A target GKE cluster to deploy to.
	//
	// It must be in the same project and region as the Dataproc cluster (the GKE cluster can be zonal or regional). Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_cluster#gke_cluster_target GoogleDataprocCluster#gke_cluster_target}
	GkeClusterTarget *string `field:"optional" json:"gkeClusterTarget" yaml:"gkeClusterTarget"`
	// node_pool_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataproc_cluster#node_pool_target GoogleDataprocCluster#node_pool_target}
	NodePoolTarget interface{} `field:"optional" json:"nodePoolTarget" yaml:"nodePoolTarget"`
}

