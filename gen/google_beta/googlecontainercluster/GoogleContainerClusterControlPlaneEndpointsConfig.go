package googlecontainercluster


type GoogleContainerClusterControlPlaneEndpointsConfig struct {
	// dns_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_container_cluster#dns_endpoint_config GoogleContainerCluster#dns_endpoint_config}
	DnsEndpointConfig *GoogleContainerClusterControlPlaneEndpointsConfigDnsEndpointConfig `field:"optional" json:"dnsEndpointConfig" yaml:"dnsEndpointConfig"`
}

