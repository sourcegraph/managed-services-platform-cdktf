package accessidentityprovider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProvider",
		reflect.TypeOf((*AccessIdentityProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "configInput", GoGetter: "ConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putConfig", GoMethod: "PutConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putScimConfig", GoMethod: "PutScimConfig"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetConfig", GoMethod: "ResetConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetScimConfig", GoMethod: "ResetScimConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneId", GoMethod: "ResetZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "scimConfig", GoGetter: "ScimConfig"},
			_jsii_.MemberProperty{JsiiProperty: "scimConfigInput", GoGetter: "ScimConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "zoneId", GoGetter: "ZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "zoneIdInput", GoGetter: "ZoneIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_AccessIdentityProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderConfig",
		reflect.TypeOf((*AccessIdentityProviderConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderConfigA",
		reflect.TypeOf((*AccessIdentityProviderConfigA)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderConfigAList",
		reflect.TypeOf((*AccessIdentityProviderConfigAList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AccessIdentityProviderConfigAList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderConfigAOutputReference",
		reflect.TypeOf((*AccessIdentityProviderConfigAOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apiToken", GoGetter: "ApiToken"},
			_jsii_.MemberProperty{JsiiProperty: "apiTokenInput", GoGetter: "ApiTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "appsDomain", GoGetter: "AppsDomain"},
			_jsii_.MemberProperty{JsiiProperty: "appsDomainInput", GoGetter: "AppsDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "attributes", GoGetter: "Attributes"},
			_jsii_.MemberProperty{JsiiProperty: "attributesInput", GoGetter: "AttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationServerId", GoGetter: "AuthorizationServerId"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationServerIdInput", GoGetter: "AuthorizationServerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "authUrl", GoGetter: "AuthUrl"},
			_jsii_.MemberProperty{JsiiProperty: "authUrlInput", GoGetter: "AuthUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "centrifyAccount", GoGetter: "CentrifyAccount"},
			_jsii_.MemberProperty{JsiiProperty: "centrifyAccountInput", GoGetter: "CentrifyAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "centrifyAppId", GoGetter: "CentrifyAppId"},
			_jsii_.MemberProperty{JsiiProperty: "centrifyAppIdInput", GoGetter: "CentrifyAppIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "certsUrl", GoGetter: "CertsUrl"},
			_jsii_.MemberProperty{JsiiProperty: "certsUrlInput", GoGetter: "CertsUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "claims", GoGetter: "Claims"},
			_jsii_.MemberProperty{JsiiProperty: "claimsInput", GoGetter: "ClaimsInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalAccessEnabled", GoGetter: "ConditionalAccessEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalAccessEnabledInput", GoGetter: "ConditionalAccessEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "directoryId", GoGetter: "DirectoryId"},
			_jsii_.MemberProperty{JsiiProperty: "directoryIdInput", GoGetter: "DirectoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailAttributeName", GoGetter: "EmailAttributeName"},
			_jsii_.MemberProperty{JsiiProperty: "emailAttributeNameInput", GoGetter: "EmailAttributeNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailClaimName", GoGetter: "EmailClaimName"},
			_jsii_.MemberProperty{JsiiProperty: "emailClaimNameInput", GoGetter: "EmailClaimNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "idpPublicCert", GoGetter: "IdpPublicCert"},
			_jsii_.MemberProperty{JsiiProperty: "idpPublicCertInput", GoGetter: "IdpPublicCertInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "issuerUrl", GoGetter: "IssuerUrl"},
			_jsii_.MemberProperty{JsiiProperty: "issuerUrlInput", GoGetter: "IssuerUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "oktaAccount", GoGetter: "OktaAccount"},
			_jsii_.MemberProperty{JsiiProperty: "oktaAccountInput", GoGetter: "OktaAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "oneloginAccount", GoGetter: "OneloginAccount"},
			_jsii_.MemberProperty{JsiiProperty: "oneloginAccountInput", GoGetter: "OneloginAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "pingEnvId", GoGetter: "PingEnvId"},
			_jsii_.MemberProperty{JsiiProperty: "pingEnvIdInput", GoGetter: "PingEnvIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "pkceEnabled", GoGetter: "PkceEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "pkceEnabledInput", GoGetter: "PkceEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUrl", GoGetter: "RedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetApiToken", GoMethod: "ResetApiToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppsDomain", GoMethod: "ResetAppsDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetAttributes", GoMethod: "ResetAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationServerId", GoMethod: "ResetAuthorizationServerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthUrl", GoMethod: "ResetAuthUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetCentrifyAccount", GoMethod: "ResetCentrifyAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetCentrifyAppId", GoMethod: "ResetCentrifyAppId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertsUrl", GoMethod: "ResetCertsUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetClaims", GoMethod: "ResetClaims"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecret", GoMethod: "ResetClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetConditionalAccessEnabled", GoMethod: "ResetConditionalAccessEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirectoryId", GoMethod: "ResetDirectoryId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailAttributeName", GoMethod: "ResetEmailAttributeName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailClaimName", GoMethod: "ResetEmailClaimName"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdpPublicCert", GoMethod: "ResetIdpPublicCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssuerUrl", GoMethod: "ResetIssuerUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetOktaAccount", GoMethod: "ResetOktaAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetOneloginAccount", GoMethod: "ResetOneloginAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetPingEnvId", GoMethod: "ResetPingEnvId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPkceEnabled", GoMethod: "ResetPkceEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignRequest", GoMethod: "ResetSignRequest"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsoTargetUrl", GoMethod: "ResetSsoTargetUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetSupportGroups", GoMethod: "ResetSupportGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenUrl", GoMethod: "ResetTokenUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "signRequest", GoGetter: "SignRequest"},
			_jsii_.MemberProperty{JsiiProperty: "signRequestInput", GoGetter: "SignRequestInput"},
			_jsii_.MemberProperty{JsiiProperty: "ssoTargetUrl", GoGetter: "SsoTargetUrl"},
			_jsii_.MemberProperty{JsiiProperty: "ssoTargetUrlInput", GoGetter: "SsoTargetUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "supportGroups", GoGetter: "SupportGroups"},
			_jsii_.MemberProperty{JsiiProperty: "supportGroupsInput", GoGetter: "SupportGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenUrl", GoGetter: "TokenUrl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenUrlInput", GoGetter: "TokenUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_AccessIdentityProviderConfigAOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderScimConfig",
		reflect.TypeOf((*AccessIdentityProviderScimConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderScimConfigList",
		reflect.TypeOf((*AccessIdentityProviderScimConfigList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_AccessIdentityProviderScimConfigList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.accessIdentityProvider.AccessIdentityProviderScimConfigOutputReference",
		reflect.TypeOf((*AccessIdentityProviderScimConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "groupMemberDeprovision", GoGetter: "GroupMemberDeprovision"},
			_jsii_.MemberProperty{JsiiProperty: "groupMemberDeprovisionInput", GoGetter: "GroupMemberDeprovisionInput"},
			_jsii_.MemberProperty{JsiiProperty: "identityUpdateBehavior", GoGetter: "IdentityUpdateBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "identityUpdateBehaviorInput", GoGetter: "IdentityUpdateBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupMemberDeprovision", GoMethod: "ResetGroupMemberDeprovision"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityUpdateBehavior", GoMethod: "ResetIdentityUpdateBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeatDeprovision", GoMethod: "ResetSeatDeprovision"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecret", GoMethod: "ResetSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserDeprovision", GoMethod: "ResetUserDeprovision"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "seatDeprovision", GoGetter: "SeatDeprovision"},
			_jsii_.MemberProperty{JsiiProperty: "seatDeprovisionInput", GoGetter: "SeatDeprovisionInput"},
			_jsii_.MemberProperty{JsiiProperty: "secret", GoGetter: "Secret"},
			_jsii_.MemberProperty{JsiiProperty: "secretInput", GoGetter: "SecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userDeprovision", GoGetter: "UserDeprovision"},
			_jsii_.MemberProperty{JsiiProperty: "userDeprovisionInput", GoGetter: "UserDeprovisionInput"},
		},
		func() interface{} {
			j := jsiiProxy_AccessIdentityProviderScimConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
