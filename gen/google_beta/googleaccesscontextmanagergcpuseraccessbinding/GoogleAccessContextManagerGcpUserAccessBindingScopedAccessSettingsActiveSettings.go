package googleaccesscontextmanagergcpuseraccessbinding


type GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettings struct {
	// Optional.
	//
	// Access level that a user must have to be granted access. Only one access level is supported, not multiple. This repeated field must have exactly one element. Example: "accessPolicies/9522/accessLevels/device_trusted"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_gcp_user_access_binding#access_levels GoogleAccessContextManagerGcpUserAccessBinding#access_levels}
	AccessLevels *[]*string `field:"optional" json:"accessLevels" yaml:"accessLevels"`
	// session_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_gcp_user_access_binding#session_settings GoogleAccessContextManagerGcpUserAccessBinding#session_settings}
	SessionSettings *GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsActiveSettingsSessionSettings `field:"optional" json:"sessionSettings" yaml:"sessionSettings"`
}

