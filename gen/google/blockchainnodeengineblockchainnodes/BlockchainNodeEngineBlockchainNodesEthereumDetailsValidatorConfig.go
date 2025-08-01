package blockchainnodeengineblockchainnodes


type BlockchainNodeEngineBlockchainNodesEthereumDetailsValidatorConfig struct {
	// URLs for MEV-relay services to use for block building.
	//
	// When set, a managed MEV-boost service is configured on the beacon client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/blockchain_node_engine_blockchain_nodes#mev_relay_urls BlockchainNodeEngineBlockchainNodes#mev_relay_urls}
	MevRelayUrls *[]*string `field:"optional" json:"mevRelayUrls" yaml:"mevRelayUrls"`
}

