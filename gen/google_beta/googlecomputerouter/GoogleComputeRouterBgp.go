package googlecomputerouter


type GoogleComputeRouterBgp struct {
	// Local BGP Autonomous System Number (ASN).
	//
	// Must be an RFC6996
	// private ASN, either 16-bit or 32-bit. The value will be fixed for
	// this router resource. All VPN tunnels that link to this router
	// will have the same local ASN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router#asn GoogleComputeRouter#asn}
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// User-specified list of prefix groups to advertise in custom mode.
	//
	// This field can only be populated if advertiseMode is CUSTOM and
	// is advertised to all peers of the router. These groups will be
	// advertised in addition to any specified prefixes. Leave this field
	// blank to advertise no custom groups.
	//
	// This enum field has the one valid value: ALL_SUBNETS
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router#advertised_groups GoogleComputeRouter#advertised_groups}
	AdvertisedGroups *[]*string `field:"optional" json:"advertisedGroups" yaml:"advertisedGroups"`
	// advertised_ip_ranges block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router#advertised_ip_ranges GoogleComputeRouter#advertised_ip_ranges}
	AdvertisedIpRanges interface{} `field:"optional" json:"advertisedIpRanges" yaml:"advertisedIpRanges"`
	// User-specified flag to indicate which mode to use for advertisement. Default value: "DEFAULT" Possible values: ["DEFAULT", "CUSTOM"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router#advertise_mode GoogleComputeRouter#advertise_mode}
	AdvertiseMode *string `field:"optional" json:"advertiseMode" yaml:"advertiseMode"`
	// Explicitly specifies a range of valid BGP Identifiers for this Router.
	//
	// It is provided as a link-local IPv4 range (from 169.254.0.0/16), of
	// size at least /30, even if the BGP sessions are over IPv6. It must
	// not overlap with any IPv4 BGP session ranges. Other vendors commonly
	// call this router ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router#identifier_range GoogleComputeRouter#identifier_range}
	IdentifierRange *string `field:"optional" json:"identifierRange" yaml:"identifierRange"`
	// The interval in seconds between BGP keepalive messages that are sent to the peer.
	//
	// Hold time is three times the interval at which keepalive
	// messages are sent, and the hold time is the maximum number of seconds
	// allowed to elapse between successive keepalive messages that BGP
	// receives from a peer.
	//
	// BGP will use the smaller of either the local hold time value or the
	// peer's hold time value as the hold time for the BGP connection
	// between the two peers. If set, this value must be between 20 and 60.
	// The default is 20.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_router#keepalive_interval GoogleComputeRouter#keepalive_interval}
	KeepaliveInterval *float64 `field:"optional" json:"keepaliveInterval" yaml:"keepaliveInterval"`
}

