package googleaccesscontextmanagergcpuseraccessbinding


type GoogleAccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeClientScopeRestrictedClientApplication struct {
	// The OAuth client ID of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_gcp_user_access_binding#client_id GoogleAccessContextManagerGcpUserAccessBinding#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The name of the application. Example: "Cloud Console".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_gcp_user_access_binding#name GoogleAccessContextManagerGcpUserAccessBinding#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

