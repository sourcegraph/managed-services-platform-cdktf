package workbenchinstance


type WorkbenchInstanceGceSetupServiceAccounts struct {
	// Optional. Email address of the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workbench_instance#email WorkbenchInstance#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
}

