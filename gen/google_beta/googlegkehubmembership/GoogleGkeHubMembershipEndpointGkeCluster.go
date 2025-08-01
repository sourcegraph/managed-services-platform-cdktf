package googlegkehubmembership


type GoogleGkeHubMembershipEndpointGkeCluster struct {
	// Self-link of the GCP resource for the GKE cluster.
	//
	// For example: '//container.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster'.
	// It can be at the most 1000 characters in length. If the cluster is provisioned with Terraform,
	// this can be '"//container.googleapis.com/${google_container_cluster.my-cluster.id}"' or
	// 'google_container_cluster.my-cluster.id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_gke_hub_membership#resource_link GoogleGkeHubMembership#resource_link}
	ResourceLink *string `field:"required" json:"resourceLink" yaml:"resourceLink"`
}

