package googlecomposerenvironment


type GoogleComposerEnvironmentConfigWebServerConfig struct {
	// Optional.
	//
	// Machine type on which Airflow web server is running. It has to be one of: composer-n1-webserver-2, composer-n1-webserver-4 or composer-n1-webserver-8. If not specified, composer-n1-webserver-2 will be used. Value custom is returned only in response, if Airflow web server parameters were manually changed to a non-standard values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_composer_environment#machine_type GoogleComposerEnvironment#machine_type}
	MachineType *string `field:"required" json:"machineType" yaml:"machineType"`
}

