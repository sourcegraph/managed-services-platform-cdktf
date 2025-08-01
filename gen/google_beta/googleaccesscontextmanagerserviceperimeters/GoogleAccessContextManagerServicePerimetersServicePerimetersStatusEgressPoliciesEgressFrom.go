package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersStatusEgressPoliciesEgressFrom struct {
	// A list of identities that are allowed access through this 'EgressPolicy'.
	//
	// Should be in the format of email address. The email address should
	// represent individual user or service account only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_service_perimeters#identities GoogleAccessContextManagerServicePerimeters#identities}
	Identities *[]*string `field:"optional" json:"identities" yaml:"identities"`
	// Specifies the type of identities that are allowed access to outside the perimeter.
	//
	// If left unspecified, then members of 'identities' field will
	// be allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_service_perimeters#identity_type GoogleAccessContextManagerServicePerimeters#identity_type}
	IdentityType *string `field:"optional" json:"identityType" yaml:"identityType"`
	// Whether to enforce traffic restrictions based on 'sources' field.
	//
	// If the 'sources' field is non-empty, then this field must be set to 'SOURCE_RESTRICTION_ENABLED'. Possible values: ["SOURCE_RESTRICTION_UNSPECIFIED", "SOURCE_RESTRICTION_ENABLED", "SOURCE_RESTRICTION_DISABLED"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_service_perimeters#source_restriction GoogleAccessContextManagerServicePerimeters#source_restriction}
	SourceRestriction *string `field:"optional" json:"sourceRestriction" yaml:"sourceRestriction"`
	// sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_access_context_manager_service_perimeters#sources GoogleAccessContextManagerServicePerimeters#sources}
	Sources interface{} `field:"optional" json:"sources" yaml:"sources"`
}

