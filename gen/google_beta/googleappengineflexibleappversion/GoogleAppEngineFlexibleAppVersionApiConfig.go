package googleappengineflexibleappversion


type GoogleAppEngineFlexibleAppVersionApiConfig struct {
	// Path to the script from the application root directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_app_engine_flexible_app_version#script GoogleAppEngineFlexibleAppVersion#script}
	Script *string `field:"required" json:"script" yaml:"script"`
	// Action to take when users access resources that require authentication. Default value: "AUTH_FAIL_ACTION_REDIRECT" Possible values: ["AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_app_engine_flexible_app_version#auth_fail_action GoogleAppEngineFlexibleAppVersion#auth_fail_action}
	AuthFailAction *string `field:"optional" json:"authFailAction" yaml:"authFailAction"`
	// Level of login required to access this resource. Default value: "LOGIN_OPTIONAL" Possible values: ["LOGIN_OPTIONAL", "LOGIN_ADMIN", "LOGIN_REQUIRED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_app_engine_flexible_app_version#login GoogleAppEngineFlexibleAppVersion#login}
	Login *string `field:"optional" json:"login" yaml:"login"`
	// Security (HTTPS) enforcement for this URL. Possible values: ["SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_app_engine_flexible_app_version#security_level GoogleAppEngineFlexibleAppVersion#security_level}
	SecurityLevel *string `field:"optional" json:"securityLevel" yaml:"securityLevel"`
	// URL to serve the endpoint at.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_app_engine_flexible_app_version#url GoogleAppEngineFlexibleAppVersion#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

