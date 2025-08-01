package apihubplugininstance


type ApihubPluginInstanceAuthConfigUserPasswordConfig struct {
	// password block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apihub_plugin_instance#password ApihubPluginInstance#password}
	Password *ApihubPluginInstanceAuthConfigUserPasswordConfigPassword `field:"required" json:"password" yaml:"password"`
	// Username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/apihub_plugin_instance#username ApihubPluginInstance#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

