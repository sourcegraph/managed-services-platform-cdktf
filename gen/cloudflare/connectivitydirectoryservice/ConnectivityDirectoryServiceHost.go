package connectivitydirectoryservice


type ConnectivityDirectoryServiceHost struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/connectivity_directory_service#hostname ConnectivityDirectoryService#hostname}.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/connectivity_directory_service#ipv4 ConnectivityDirectoryService#ipv4}.
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/connectivity_directory_service#ipv6 ConnectivityDirectoryService#ipv6}.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/connectivity_directory_service#network ConnectivityDirectoryService#network}.
	Network *ConnectivityDirectoryServiceHostNetwork `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/connectivity_directory_service#resolver_network ConnectivityDirectoryService#resolver_network}.
	ResolverNetwork *ConnectivityDirectoryServiceHostResolverNetwork `field:"optional" json:"resolverNetwork" yaml:"resolverNetwork"`
}

