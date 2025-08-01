package workbenchinstance


type WorkbenchInstanceGceSetupConfidentialInstanceConfig struct {
	// Defines the type of technology used by the confidential instance. Possible values: ["SEV"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/workbench_instance#confidential_instance_type WorkbenchInstance#confidential_instance_type}
	ConfidentialInstanceType *string `field:"optional" json:"confidentialInstanceType" yaml:"confidentialInstanceType"`
}

