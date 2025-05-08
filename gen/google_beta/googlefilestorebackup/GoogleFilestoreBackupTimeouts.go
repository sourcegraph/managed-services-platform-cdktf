package googlefilestorebackup


type GoogleFilestoreBackupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_filestore_backup#create GoogleFilestoreBackup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_filestore_backup#delete GoogleFilestoreBackup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_filestore_backup#update GoogleFilestoreBackup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

