package netappvolume


type NetappVolumeExportPolicy struct {
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/5.21.0/docs/resources/netapp_volume#rules NetappVolume#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}
