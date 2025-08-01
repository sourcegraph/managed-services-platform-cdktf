package netappbackuppolicy


type NetappBackupPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/netapp_backup_policy#create NetappBackupPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/netapp_backup_policy#delete NetappBackupPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/netapp_backup_policy#update NetappBackupPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

