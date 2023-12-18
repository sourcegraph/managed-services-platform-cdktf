package registrymodule


type RegistryModuleTestConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.51.0/docs/resources/registry_module#tests_enabled RegistryModule#tests_enabled}.
	TestsEnabled interface{} `field:"optional" json:"testsEnabled" yaml:"testsEnabled"`
}

