package accesscontextmanagergcpuseraccessbinding


type AccessContextManagerGcpUserAccessBindingScopedAccessSettingsScopeClientScopeRestrictedClientApplication struct {
	// The OAuth client ID of the application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_gcp_user_access_binding#client_id AccessContextManagerGcpUserAccessBinding#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The name of the application. Example: "Cloud Console".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/access_context_manager_gcp_user_access_binding#name AccessContextManagerGcpUserAccessBinding#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

