package magicwanipsectunnel


type MagicWanIpsecTunnelBgp struct {
	// ASN used on the customer end of the BGP session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_wan_ipsec_tunnel#customer_asn MagicWanIpsecTunnel#customer_asn}
	CustomerAsn *float64 `field:"required" json:"customerAsn" yaml:"customerAsn"`
	// Prefixes in this list will be advertised to the customer device, in addition to the routes in the Magic routing table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_wan_ipsec_tunnel#extra_prefixes MagicWanIpsecTunnel#extra_prefixes}
	ExtraPrefixes *[]*string `field:"optional" json:"extraPrefixes" yaml:"extraPrefixes"`
	// MD5 key to use for session authentication.
	//
	// Note that *this is not a security measure*. MD5 is not a valid security mechanism, and the
	// key is not treated as a secret value. This is *only* supported for preventing
	// misconfiguration, not for defending against malicious attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the following types of
	// character:
	//
	// * ASCII alphanumerics: `[a-zA-Z0-9]`
	// * Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from newline (0x0A),
	// quotation mark (`"`), vertical tab (0x0B), carriage return (0x0D), tab (0x09), form feed
	// (0x0C), and the question mark (`?`). Requests specifying an MD5 key with one or more of
	// these disallowed characters will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/magic_wan_ipsec_tunnel#md5_key MagicWanIpsecTunnel#md5_key}
	Md5Key *string `field:"optional" json:"md5Key" yaml:"md5Key"`
}

