package googlealloydbinstance


type GoogleAlloydbInstancePscInstanceConfigPscInterfaceConfigs struct {
	// The network attachment resource created in the consumer project to which the PSC interface will be linked.
	//
	// This is of the format: "projects/${CONSUMER_PROJECT}/regions/${REGION}/networkAttachments/${NETWORK_ATTACHMENT_NAME}".
	// The network attachment must be in the same region as the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_alloydb_instance#network_attachment_resource GoogleAlloydbInstance#network_attachment_resource}
	NetworkAttachmentResource *string `field:"optional" json:"networkAttachmentResource" yaml:"networkAttachmentResource"`
}

