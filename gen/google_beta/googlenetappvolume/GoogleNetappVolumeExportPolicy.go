package googlenetappvolume


type GoogleNetappVolumeExportPolicy struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_netapp_volume#rules GoogleNetappVolume#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

