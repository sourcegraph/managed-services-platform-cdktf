package connectivitydirectoryservice


type ConnectivityDirectoryServiceHostResolverNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/connectivity_directory_service#tunnel_id ConnectivityDirectoryService#tunnel_id}.
	TunnelId *string `field:"required" json:"tunnelId" yaml:"tunnelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/connectivity_directory_service#resolver_ips ConnectivityDirectoryService#resolver_ips}.
	ResolverIps *[]*string `field:"optional" json:"resolverIps" yaml:"resolverIps"`
}

