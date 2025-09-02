package magictransitsitelan


type MagicTransitSiteLanStaticAddressingDhcpRelay struct {
	// List of DHCP server IPs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/magic_transit_site_lan#server_addresses MagicTransitSiteLan#server_addresses}
	ServerAddresses *[]*string `field:"optional" json:"serverAddresses" yaml:"serverAddresses"`
}

