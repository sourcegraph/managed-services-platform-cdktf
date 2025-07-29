package googlenetworkconnectivityspoke


type GoogleNetworkConnectivitySpokeLinkedVpnTunnels struct {
	// A value that controls whether site-to-site data transfer is enabled for these resources.
	//
	// Note that data transfer is available only in supported locations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_spoke#site_to_site_data_transfer GoogleNetworkConnectivitySpoke#site_to_site_data_transfer}
	SiteToSiteDataTransfer interface{} `field:"required" json:"siteToSiteDataTransfer" yaml:"siteToSiteDataTransfer"`
	// The URIs of linked VPN tunnel resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_spoke#uris GoogleNetworkConnectivitySpoke#uris}
	Uris *[]*string `field:"required" json:"uris" yaml:"uris"`
	// IP ranges allowed to be included during import from hub (does not control transit connectivity).
	//
	// The only allowed value for now is "ALL_IPV4_RANGES".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_network_connectivity_spoke#include_import_ranges GoogleNetworkConnectivitySpoke#include_import_ranges}
	IncludeImportRanges *[]*string `field:"optional" json:"includeImportRanges" yaml:"includeImportRanges"`
}

