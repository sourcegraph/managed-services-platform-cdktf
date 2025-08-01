package filestoreinstance


type FilestoreInstanceInitialReplicationReplicas struct {
	// The peer instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/filestore_instance#peer_instance FilestoreInstance#peer_instance}
	PeerInstance *string `field:"required" json:"peerInstance" yaml:"peerInstance"`
}

