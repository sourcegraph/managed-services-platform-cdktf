package magicwanipsectunnel


type MagicWanIpsecTunnelCustomRemoteIdentities struct {
	// A custom IKE ID of type FQDN that may be used to identity the IPsec tunnel.
	//
	// The
	// generated IKE IDs can still be used even if this custom value is specified.
	//
	// Must be of the form `<custom label>.<account ID>.custom.ipsec.cloudflare.com`.
	//
	// This custom ID does not need to be unique. Two IPsec tunnels may have the same custom
	// fqdn_id. However, if another IPsec tunnel has the same value then the two tunnels
	// cannot have the same cloudflare_endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/magic_wan_ipsec_tunnel#fqdn_id MagicWanIpsecTunnel#fqdn_id}
	FqdnId *string `field:"optional" json:"fqdnId" yaml:"fqdnId"`
}

