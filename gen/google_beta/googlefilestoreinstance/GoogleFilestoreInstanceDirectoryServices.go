package googlefilestoreinstance


type GoogleFilestoreInstanceDirectoryServices struct {
	// ldap block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_filestore_instance#ldap GoogleFilestoreInstance#ldap}
	Ldap *GoogleFilestoreInstanceDirectoryServicesLdap `field:"optional" json:"ldap" yaml:"ldap"`
}

