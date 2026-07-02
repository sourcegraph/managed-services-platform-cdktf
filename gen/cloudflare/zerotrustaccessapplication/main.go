package zerotrustaccessapplication

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		reflect.TypeFor[ZeroTrustAccessApplication](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowAuthenticateViaWarp", GoGetter: "AllowAuthenticateViaWarp"},
			_jsii_.MemberProperty{JsiiProperty: "allowAuthenticateViaWarpInput", GoGetter: "AllowAuthenticateViaWarpInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedIdps", GoGetter: "AllowedIdps"},
			_jsii_.MemberProperty{JsiiProperty: "allowedIdpsInput", GoGetter: "AllowedIdpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowIframe", GoGetter: "AllowIframe"},
			_jsii_.MemberProperty{JsiiProperty: "allowIframeInput", GoGetter: "AllowIframeInput"},
			_jsii_.MemberProperty{JsiiProperty: "appLauncherLogoUrl", GoGetter: "AppLauncherLogoUrl"},
			_jsii_.MemberProperty{JsiiProperty: "appLauncherLogoUrlInput", GoGetter: "AppLauncherLogoUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "appLauncherVisible", GoGetter: "AppLauncherVisible"},
			_jsii_.MemberProperty{JsiiProperty: "appLauncherVisibleInput", GoGetter: "AppLauncherVisibleInput"},
			_jsii_.MemberProperty{JsiiProperty: "aud", GoGetter: "Aud"},
			_jsii_.MemberProperty{JsiiProperty: "autoRedirectToIdentity", GoGetter: "AutoRedirectToIdentity"},
			_jsii_.MemberProperty{JsiiProperty: "autoRedirectToIdentityInput", GoGetter: "AutoRedirectToIdentityInput"},
			_jsii_.MemberProperty{JsiiProperty: "bgColor", GoGetter: "BgColor"},
			_jsii_.MemberProperty{JsiiProperty: "bgColorInput", GoGetter: "BgColorInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "corsHeaders", GoGetter: "CorsHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "corsHeadersInput", GoGetter: "CorsHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customDenyMessage", GoGetter: "CustomDenyMessage"},
			_jsii_.MemberProperty{JsiiProperty: "customDenyMessageInput", GoGetter: "CustomDenyMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "customDenyUrl", GoGetter: "CustomDenyUrl"},
			_jsii_.MemberProperty{JsiiProperty: "customDenyUrlInput", GoGetter: "CustomDenyUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "customNonIdentityDenyUrl", GoGetter: "CustomNonIdentityDenyUrl"},
			_jsii_.MemberProperty{JsiiProperty: "customNonIdentityDenyUrlInput", GoGetter: "CustomNonIdentityDenyUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "customPages", GoGetter: "CustomPages"},
			_jsii_.MemberProperty{JsiiProperty: "customPagesInput", GoGetter: "CustomPagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "destinations", GoGetter: "Destinations"},
			_jsii_.MemberProperty{JsiiProperty: "destinationsInput", GoGetter: "DestinationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableBindingCookie", GoGetter: "EnableBindingCookie"},
			_jsii_.MemberProperty{JsiiProperty: "enableBindingCookieInput", GoGetter: "EnableBindingCookieInput"},
			_jsii_.MemberProperty{JsiiProperty: "footerLinks", GoGetter: "FooterLinks"},
			_jsii_.MemberProperty{JsiiProperty: "footerLinksInput", GoGetter: "FooterLinksInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "headerBgColor", GoGetter: "HeaderBgColor"},
			_jsii_.MemberProperty{JsiiProperty: "headerBgColorInput", GoGetter: "HeaderBgColorInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpOnlyCookieAttribute", GoGetter: "HttpOnlyCookieAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "httpOnlyCookieAttributeInput", GoGetter: "HttpOnlyCookieAttributeInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "landingPageDesign", GoGetter: "LandingPageDesign"},
			_jsii_.MemberProperty{JsiiProperty: "landingPageDesignInput", GoGetter: "LandingPageDesignInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logoUrl", GoGetter: "LogoUrl"},
			_jsii_.MemberProperty{JsiiProperty: "logoUrlInput", GoGetter: "LogoUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "optionsPreflightBypass", GoGetter: "OptionsPreflightBypass"},
			_jsii_.MemberProperty{JsiiProperty: "optionsPreflightBypassInput", GoGetter: "OptionsPreflightBypassInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pathCookieAttribute", GoGetter: "PathCookieAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pathCookieAttributeInput", GoGetter: "PathCookieAttributeInput"},
			_jsii_.MemberProperty{JsiiProperty: "policies", GoGetter: "Policies"},
			_jsii_.MemberProperty{JsiiProperty: "policiesInput", GoGetter: "PoliciesInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCorsHeaders", GoMethod: "PutCorsHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "putDestinations", GoMethod: "PutDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "putFooterLinks", GoMethod: "PutFooterLinks"},
			_jsii_.MemberMethod{JsiiMethod: "putLandingPageDesign", GoMethod: "PutLandingPageDesign"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicies", GoMethod: "PutPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "putSaasApp", GoMethod: "PutSaasApp"},
			_jsii_.MemberMethod{JsiiMethod: "putScimConfig", GoMethod: "PutScimConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putTargetCriteria", GoMethod: "PutTargetCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readServiceTokensFromHeader", GoGetter: "ReadServiceTokensFromHeader"},
			_jsii_.MemberProperty{JsiiProperty: "readServiceTokensFromHeaderInput", GoGetter: "ReadServiceTokensFromHeaderInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccountId", GoMethod: "ResetAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAuthenticateViaWarp", GoMethod: "ResetAllowAuthenticateViaWarp"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedIdps", GoMethod: "ResetAllowedIdps"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowIframe", GoMethod: "ResetAllowIframe"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLauncherLogoUrl", GoMethod: "ResetAppLauncherLogoUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLauncherVisible", GoMethod: "ResetAppLauncherVisible"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoRedirectToIdentity", GoMethod: "ResetAutoRedirectToIdentity"},
			_jsii_.MemberMethod{JsiiMethod: "resetBgColor", GoMethod: "ResetBgColor"},
			_jsii_.MemberMethod{JsiiMethod: "resetCorsHeaders", GoMethod: "ResetCorsHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomDenyMessage", GoMethod: "ResetCustomDenyMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomDenyUrl", GoMethod: "ResetCustomDenyUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomNonIdentityDenyUrl", GoMethod: "ResetCustomNonIdentityDenyUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomPages", GoMethod: "ResetCustomPages"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinations", GoMethod: "ResetDestinations"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomain", GoMethod: "ResetDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableBindingCookie", GoMethod: "ResetEnableBindingCookie"},
			_jsii_.MemberMethod{JsiiMethod: "resetFooterLinks", GoMethod: "ResetFooterLinks"},
			_jsii_.MemberMethod{JsiiMethod: "resetHeaderBgColor", GoMethod: "ResetHeaderBgColor"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpOnlyCookieAttribute", GoMethod: "ResetHttpOnlyCookieAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetLandingPageDesign", GoMethod: "ResetLandingPageDesign"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogoUrl", GoMethod: "ResetLogoUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionsPreflightBypass", GoMethod: "ResetOptionsPreflightBypass"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPathCookieAttribute", GoMethod: "ResetPathCookieAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicies", GoMethod: "ResetPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadServiceTokensFromHeader", GoMethod: "ResetReadServiceTokensFromHeader"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaasApp", GoMethod: "ResetSaasApp"},
			_jsii_.MemberMethod{JsiiMethod: "resetSameSiteCookieAttribute", GoMethod: "ResetSameSiteCookieAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetScimConfig", GoMethod: "ResetScimConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelfHostedDomains", GoMethod: "ResetSelfHostedDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAuth401Redirect", GoMethod: "ResetServiceAuth401Redirect"},
			_jsii_.MemberMethod{JsiiMethod: "resetSessionDuration", GoMethod: "ResetSessionDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipAppLauncherLoginPage", GoMethod: "ResetSkipAppLauncherLoginPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipInterstitial", GoMethod: "ResetSkipInterstitial"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetCriteria", GoMethod: "ResetTargetCriteria"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetZoneId", GoMethod: "ResetZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "saasApp", GoGetter: "SaasApp"},
			_jsii_.MemberProperty{JsiiProperty: "saasAppInput", GoGetter: "SaasAppInput"},
			_jsii_.MemberProperty{JsiiProperty: "sameSiteCookieAttribute", GoGetter: "SameSiteCookieAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "sameSiteCookieAttributeInput", GoGetter: "SameSiteCookieAttributeInput"},
			_jsii_.MemberProperty{JsiiProperty: "scimConfig", GoGetter: "ScimConfig"},
			_jsii_.MemberProperty{JsiiProperty: "scimConfigInput", GoGetter: "ScimConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "selfHostedDomains", GoGetter: "SelfHostedDomains"},
			_jsii_.MemberProperty{JsiiProperty: "selfHostedDomainsInput", GoGetter: "SelfHostedDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAuth401Redirect", GoGetter: "ServiceAuth401Redirect"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAuth401RedirectInput", GoGetter: "ServiceAuth401RedirectInput"},
			_jsii_.MemberProperty{JsiiProperty: "sessionDuration", GoGetter: "SessionDuration"},
			_jsii_.MemberProperty{JsiiProperty: "sessionDurationInput", GoGetter: "SessionDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipAppLauncherLoginPage", GoGetter: "SkipAppLauncherLoginPage"},
			_jsii_.MemberProperty{JsiiProperty: "skipAppLauncherLoginPageInput", GoGetter: "SkipAppLauncherLoginPageInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipInterstitial", GoGetter: "SkipInterstitial"},
			_jsii_.MemberProperty{JsiiProperty: "skipInterstitialInput", GoGetter: "SkipInterstitialInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetCriteria", GoGetter: "TargetCriteria"},
			_jsii_.MemberProperty{JsiiProperty: "targetCriteriaInput", GoGetter: "TargetCriteriaInput"},
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplication{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationConfig",
		reflect.TypeFor[ZeroTrustAccessApplicationConfig](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationCorsHeaders",
		reflect.TypeFor[ZeroTrustAccessApplicationCorsHeaders](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationCorsHeadersOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationCorsHeadersOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowAllHeaders", GoGetter: "AllowAllHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllHeadersInput", GoGetter: "AllowAllHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllMethods", GoGetter: "AllowAllMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllMethodsInput", GoGetter: "AllowAllMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllOrigins", GoGetter: "AllowAllOrigins"},
			_jsii_.MemberProperty{JsiiProperty: "allowAllOriginsInput", GoGetter: "AllowAllOriginsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowCredentials", GoGetter: "AllowCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "allowCredentialsInput", GoGetter: "AllowCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedHeaders", GoGetter: "AllowedHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "allowedHeadersInput", GoGetter: "AllowedHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedMethods", GoGetter: "AllowedMethods"},
			_jsii_.MemberProperty{JsiiProperty: "allowedMethodsInput", GoGetter: "AllowedMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOrigins", GoGetter: "AllowedOrigins"},
			_jsii_.MemberProperty{JsiiProperty: "allowedOriginsInput", GoGetter: "AllowedOriginsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxAge", GoGetter: "MaxAge"},
			_jsii_.MemberProperty{JsiiProperty: "maxAgeInput", GoGetter: "MaxAgeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAllHeaders", GoMethod: "ResetAllowAllHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAllMethods", GoMethod: "ResetAllowAllMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAllOrigins", GoMethod: "ResetAllowAllOrigins"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowCredentials", GoMethod: "ResetAllowCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedHeaders", GoMethod: "ResetAllowedHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedMethods", GoMethod: "ResetAllowedMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowedOrigins", GoMethod: "ResetAllowedOrigins"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxAge", GoMethod: "ResetMaxAge"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationCorsHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationDestinations",
		reflect.TypeFor[ZeroTrustAccessApplicationDestinations](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationDestinationsList",
		reflect.TypeFor[ZeroTrustAccessApplicationDestinationsList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationDestinationsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationDestinationsOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationDestinationsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cidr", GoGetter: "Cidr"},
			_jsii_.MemberProperty{JsiiProperty: "cidrInput", GoGetter: "CidrInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameInput", GoGetter: "HostnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "l4Protocol", GoGetter: "L4Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "l4ProtocolInput", GoGetter: "L4ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "portRange", GoGetter: "PortRange"},
			_jsii_.MemberProperty{JsiiProperty: "portRangeInput", GoGetter: "PortRangeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCidr", GoMethod: "ResetCidr"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostname", GoMethod: "ResetHostname"},
			_jsii_.MemberMethod{JsiiMethod: "resetL4Protocol", GoMethod: "ResetL4Protocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetPortRange", GoMethod: "ResetPortRange"},
			_jsii_.MemberMethod{JsiiMethod: "resetType", GoMethod: "ResetType"},
			_jsii_.MemberMethod{JsiiMethod: "resetUri", GoMethod: "ResetUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetVnetId", GoMethod: "ResetVnetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "uri", GoGetter: "Uri"},
			_jsii_.MemberProperty{JsiiProperty: "uriInput", GoGetter: "UriInput"},
			_jsii_.MemberProperty{JsiiProperty: "vnetId", GoGetter: "VnetId"},
			_jsii_.MemberProperty{JsiiProperty: "vnetIdInput", GoGetter: "VnetIdInput"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationDestinationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationFooterLinks",
		reflect.TypeFor[ZeroTrustAccessApplicationFooterLinks](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationFooterLinksList",
		reflect.TypeFor[ZeroTrustAccessApplicationFooterLinksList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationFooterLinksList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationFooterLinksOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationFooterLinksOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationFooterLinksOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationLandingPageDesign",
		reflect.TypeFor[ZeroTrustAccessApplicationLandingPageDesign](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationLandingPageDesignOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationLandingPageDesignOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "buttonColor", GoGetter: "ButtonColor"},
			_jsii_.MemberProperty{JsiiProperty: "buttonColorInput", GoGetter: "ButtonColorInput"},
			_jsii_.MemberProperty{JsiiProperty: "buttonTextColor", GoGetter: "ButtonTextColor"},
			_jsii_.MemberProperty{JsiiProperty: "buttonTextColorInput", GoGetter: "ButtonTextColorInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "imageUrl", GoGetter: "ImageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "imageUrlInput", GoGetter: "ImageUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "message", GoGetter: "Message"},
			_jsii_.MemberProperty{JsiiProperty: "messageInput", GoGetter: "MessageInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetButtonColor", GoMethod: "ResetButtonColor"},
			_jsii_.MemberMethod{JsiiMethod: "resetButtonTextColor", GoMethod: "ResetButtonTextColor"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageUrl", GoMethod: "ResetImageUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessage", GoMethod: "ResetMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetTitle", GoMethod: "ResetTitle"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberProperty{JsiiProperty: "titleInput", GoGetter: "TitleInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationLandingPageDesignOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPolicies",
		reflect.TypeFor[ZeroTrustAccessApplicationPolicies](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesConnectionRules",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesConnectionRules](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putSsh", GoMethod: "PutSsh"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsh", GoMethod: "ResetSsh"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssh", GoGetter: "Ssh"},
			_jsii_.MemberProperty{JsiiProperty: "sshInput", GoGetter: "SshInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesConnectionRulesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesConnectionRulesSsh",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesConnectionRulesSsh](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesConnectionRulesSshOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesConnectionRulesSshOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowEmailAlias", GoGetter: "AllowEmailAlias"},
			_jsii_.MemberProperty{JsiiProperty: "allowEmailAliasInput", GoGetter: "AllowEmailAliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowEmailAlias", GoMethod: "ResetAllowEmailAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "usernames", GoGetter: "Usernames"},
			_jsii_.MemberProperty{JsiiProperty: "usernamesInput", GoGetter: "UsernamesInput"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesConnectionRulesSshOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExclude",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExclude](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceToken",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceToken](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeAnyValidServiceTokenOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAuthContext",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAuthContext](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acId", GoGetter: "AcId"},
			_jsii_.MemberProperty{JsiiProperty: "acIdInput", GoGetter: "AcIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeAuthContextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAuthMethod",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAuthMethod](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authMethod", GoGetter: "AuthMethod"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodInput", GoGetter: "AuthMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeAuthMethodOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAzureAd",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAzureAd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeAzureAdOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeCertificate",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeCertificate](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeCommonName",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeCommonName](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeCommonNameOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeDevicePosture",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeDevicePosture](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "integrationUid", GoGetter: "IntegrationUid"},
			_jsii_.MemberProperty{JsiiProperty: "integrationUidInput", GoGetter: "IntegrationUidInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeDevicePostureOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEmail",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEmail](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEmailDomain",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEmailDomain](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeEmailDomainOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEmailListStruct",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEmailListStruct](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeEmailListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeEmailOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEveryone",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEveryone](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeEveryoneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluation",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluation](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateUrl", GoGetter: "EvaluateUrl"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateUrlInput", GoGetter: "EvaluateUrlInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keysUrl", GoGetter: "KeysUrl"},
			_jsii_.MemberProperty{JsiiProperty: "keysUrlInput", GoGetter: "KeysUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeExternalEvaluationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGeo",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGeo](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "countryCode", GoGetter: "CountryCode"},
			_jsii_.MemberProperty{JsiiProperty: "countryCodeInput", GoGetter: "CountryCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeGeoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGithubOrganization",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGithubOrganization](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeam", GoMethod: "ResetTeam"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "team", GoGetter: "Team"},
			_jsii_.MemberProperty{JsiiProperty: "teamInput", GoGetter: "TeamInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeGithubOrganizationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGroup",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGroup](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeGroupOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGsuite",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGsuite](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeGsuiteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeIp",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeIp](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeIpListStruct",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeIpListStruct](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeIpListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeIpOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeIpOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "ipInput", GoGetter: "IpInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeIpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeList",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeLoginMethod",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeLoginMethod](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeLoginMethodOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeOidc",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeOidc](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeOidcOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeOidcOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "claimName", GoGetter: "ClaimName"},
			_jsii_.MemberProperty{JsiiProperty: "claimNameInput", GoGetter: "ClaimNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "claimValue", GoGetter: "ClaimValue"},
			_jsii_.MemberProperty{JsiiProperty: "claimValueInput", GoGetter: "ClaimValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOidcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeOkta",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeOkta](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOktaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "anyValidServiceToken", GoGetter: "AnyValidServiceToken"},
			_jsii_.MemberProperty{JsiiProperty: "anyValidServiceTokenInput", GoGetter: "AnyValidServiceTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "authContext", GoGetter: "AuthContext"},
			_jsii_.MemberProperty{JsiiProperty: "authContextInput", GoGetter: "AuthContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "authMethod", GoGetter: "AuthMethod"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodInput", GoGetter: "AuthMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureAd", GoGetter: "AzureAd"},
			_jsii_.MemberProperty{JsiiProperty: "azureAdInput", GoGetter: "AzureAdInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "devicePosture", GoGetter: "DevicePosture"},
			_jsii_.MemberProperty{JsiiProperty: "devicePostureInput", GoGetter: "DevicePostureInput"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailDomain", GoGetter: "EmailDomain"},
			_jsii_.MemberProperty{JsiiProperty: "emailDomainInput", GoGetter: "EmailDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailList", GoGetter: "EmailList"},
			_jsii_.MemberProperty{JsiiProperty: "emailListInput", GoGetter: "EmailListInput"},
			_jsii_.MemberProperty{JsiiProperty: "everyone", GoGetter: "Everyone"},
			_jsii_.MemberProperty{JsiiProperty: "everyoneInput", GoGetter: "EveryoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalEvaluation", GoGetter: "ExternalEvaluation"},
			_jsii_.MemberProperty{JsiiProperty: "externalEvaluationInput", GoGetter: "ExternalEvaluationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "geo", GoGetter: "Geo"},
			_jsii_.MemberProperty{JsiiProperty: "geoInput", GoGetter: "GeoInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "githubOrganization", GoGetter: "GithubOrganization"},
			_jsii_.MemberProperty{JsiiProperty: "githubOrganizationInput", GoGetter: "GithubOrganizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
			_jsii_.MemberProperty{JsiiProperty: "groupInput", GoGetter: "GroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "gsuite", GoGetter: "Gsuite"},
			_jsii_.MemberProperty{JsiiProperty: "gsuiteInput", GoGetter: "GsuiteInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "ipInput", GoGetter: "IpInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipList", GoGetter: "IpList"},
			_jsii_.MemberProperty{JsiiProperty: "ipListInput", GoGetter: "IpListInput"},
			_jsii_.MemberProperty{JsiiProperty: "loginMethod", GoGetter: "LoginMethod"},
			_jsii_.MemberProperty{JsiiProperty: "loginMethodInput", GoGetter: "LoginMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidc", GoGetter: "Oidc"},
			_jsii_.MemberProperty{JsiiProperty: "oidcInput", GoGetter: "OidcInput"},
			_jsii_.MemberProperty{JsiiProperty: "okta", GoGetter: "Okta"},
			_jsii_.MemberProperty{JsiiProperty: "oktaInput", GoGetter: "OktaInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAnyValidServiceToken", GoMethod: "PutAnyValidServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthContext", GoMethod: "PutAuthContext"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthMethod", GoMethod: "PutAuthMethod"},
			_jsii_.MemberMethod{JsiiMethod: "putAzureAd", GoMethod: "PutAzureAd"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putCommonName", GoMethod: "PutCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "putDevicePosture", GoMethod: "PutDevicePosture"},
			_jsii_.MemberMethod{JsiiMethod: "putEmail", GoMethod: "PutEmail"},
			_jsii_.MemberMethod{JsiiMethod: "putEmailDomain", GoMethod: "PutEmailDomain"},
			_jsii_.MemberMethod{JsiiMethod: "putEmailList", GoMethod: "PutEmailList"},
			_jsii_.MemberMethod{JsiiMethod: "putEveryone", GoMethod: "PutEveryone"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalEvaluation", GoMethod: "PutExternalEvaluation"},
			_jsii_.MemberMethod{JsiiMethod: "putGeo", GoMethod: "PutGeo"},
			_jsii_.MemberMethod{JsiiMethod: "putGithubOrganization", GoMethod: "PutGithubOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "putGroup", GoMethod: "PutGroup"},
			_jsii_.MemberMethod{JsiiMethod: "putGsuite", GoMethod: "PutGsuite"},
			_jsii_.MemberMethod{JsiiMethod: "putIp", GoMethod: "PutIp"},
			_jsii_.MemberMethod{JsiiMethod: "putIpList", GoMethod: "PutIpList"},
			_jsii_.MemberMethod{JsiiMethod: "putLoginMethod", GoMethod: "PutLoginMethod"},
			_jsii_.MemberMethod{JsiiMethod: "putOidc", GoMethod: "PutOidc"},
			_jsii_.MemberMethod{JsiiMethod: "putOkta", GoMethod: "PutOkta"},
			_jsii_.MemberMethod{JsiiMethod: "putSaml", GoMethod: "PutSaml"},
			_jsii_.MemberMethod{JsiiMethod: "putServiceToken", GoMethod: "PutServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnyValidServiceToken", GoMethod: "ResetAnyValidServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthContext", GoMethod: "ResetAuthContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthMethod", GoMethod: "ResetAuthMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureAd", GoMethod: "ResetAzureAd"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommonName", GoMethod: "ResetCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDevicePosture", GoMethod: "ResetDevicePosture"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailDomain", GoMethod: "ResetEmailDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailList", GoMethod: "ResetEmailList"},
			_jsii_.MemberMethod{JsiiMethod: "resetEveryone", GoMethod: "ResetEveryone"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalEvaluation", GoMethod: "ResetExternalEvaluation"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeo", GoMethod: "ResetGeo"},
			_jsii_.MemberMethod{JsiiMethod: "resetGithubOrganization", GoMethod: "ResetGithubOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroup", GoMethod: "ResetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetGsuite", GoMethod: "ResetGsuite"},
			_jsii_.MemberMethod{JsiiMethod: "resetIp", GoMethod: "ResetIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpList", GoMethod: "ResetIpList"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginMethod", GoMethod: "ResetLoginMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidc", GoMethod: "ResetOidc"},
			_jsii_.MemberMethod{JsiiMethod: "resetOkta", GoMethod: "ResetOkta"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml", GoMethod: "ResetSaml"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceToken", GoMethod: "ResetServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "saml", GoGetter: "Saml"},
			_jsii_.MemberProperty{JsiiProperty: "samlInput", GoGetter: "SamlInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberProperty{JsiiProperty: "serviceTokenInput", GoGetter: "ServiceTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeSaml",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeSaml](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributeName", GoGetter: "AttributeName"},
			_jsii_.MemberProperty{JsiiProperty: "attributeNameInput", GoGetter: "AttributeNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValue", GoGetter: "AttributeValue"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValueInput", GoGetter: "AttributeValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeSamlOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeServiceToken",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeServiceToken](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenId", GoGetter: "TokenId"},
			_jsii_.MemberProperty{JsiiProperty: "tokenIdInput", GoGetter: "TokenIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesExcludeServiceTokenOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesInclude",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesInclude](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceToken",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceToken](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeAnyValidServiceTokenOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAuthContext",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAuthContext](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAuthContextOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAuthContextOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acId", GoGetter: "AcId"},
			_jsii_.MemberProperty{JsiiProperty: "acIdInput", GoGetter: "AcIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeAuthContextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAuthMethod",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAuthMethod](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAuthMethodOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAuthMethodOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authMethod", GoGetter: "AuthMethod"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodInput", GoGetter: "AuthMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeAuthMethodOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAzureAd",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAzureAd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeAzureAdOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeAzureAdOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeAzureAdOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeCertificate",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeCertificate](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeCertificateOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeCertificateOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeCommonName",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeCommonName](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeCommonNameOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeCommonNameOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeCommonNameOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeDevicePosture",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeDevicePosture](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeDevicePostureOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeDevicePostureOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "integrationUid", GoGetter: "IntegrationUid"},
			_jsii_.MemberProperty{JsiiProperty: "integrationUidInput", GoGetter: "IntegrationUidInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeDevicePostureOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEmail",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEmail](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEmailDomain",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEmailDomain](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEmailDomainOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEmailDomainOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeEmailDomainOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEmailListStruct",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEmailListStruct](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEmailListStructOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEmailListStructOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeEmailListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEmailOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEmailOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeEmailOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEveryone",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEveryone](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeEveryoneOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeEveryoneOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeEveryoneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluation",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluation](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateUrl", GoGetter: "EvaluateUrl"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateUrlInput", GoGetter: "EvaluateUrlInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keysUrl", GoGetter: "KeysUrl"},
			_jsii_.MemberProperty{JsiiProperty: "keysUrlInput", GoGetter: "KeysUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeExternalEvaluationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGeo",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGeo](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGeoOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGeoOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "countryCode", GoGetter: "CountryCode"},
			_jsii_.MemberProperty{JsiiProperty: "countryCodeInput", GoGetter: "CountryCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeGeoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGithubOrganization",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGithubOrganization](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGithubOrganizationOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGithubOrganizationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeam", GoMethod: "ResetTeam"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "team", GoGetter: "Team"},
			_jsii_.MemberProperty{JsiiProperty: "teamInput", GoGetter: "TeamInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeGithubOrganizationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGroup",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGroup](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGroupOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGroupOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeGroupOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGsuite",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGsuite](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeGsuiteOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeGsuiteOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeGsuiteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeIp",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeIp](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeIpListStruct",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeIpListStruct](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeIpListStructOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeIpListStructOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeIpListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeIpOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeIpOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "ipInput", GoGetter: "IpInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeIpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeList",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeLoginMethod",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeLoginMethod](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeLoginMethodOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeLoginMethodOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeLoginMethodOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeOidc",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeOidc](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeOidcOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeOidcOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "claimName", GoGetter: "ClaimName"},
			_jsii_.MemberProperty{JsiiProperty: "claimNameInput", GoGetter: "ClaimNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "claimValue", GoGetter: "ClaimValue"},
			_jsii_.MemberProperty{JsiiProperty: "claimValueInput", GoGetter: "ClaimValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOidcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeOkta",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeOkta](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeOktaOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeOktaOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOktaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "anyValidServiceToken", GoGetter: "AnyValidServiceToken"},
			_jsii_.MemberProperty{JsiiProperty: "anyValidServiceTokenInput", GoGetter: "AnyValidServiceTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "authContext", GoGetter: "AuthContext"},
			_jsii_.MemberProperty{JsiiProperty: "authContextInput", GoGetter: "AuthContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "authMethod", GoGetter: "AuthMethod"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodInput", GoGetter: "AuthMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureAd", GoGetter: "AzureAd"},
			_jsii_.MemberProperty{JsiiProperty: "azureAdInput", GoGetter: "AzureAdInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "devicePosture", GoGetter: "DevicePosture"},
			_jsii_.MemberProperty{JsiiProperty: "devicePostureInput", GoGetter: "DevicePostureInput"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailDomain", GoGetter: "EmailDomain"},
			_jsii_.MemberProperty{JsiiProperty: "emailDomainInput", GoGetter: "EmailDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailList", GoGetter: "EmailList"},
			_jsii_.MemberProperty{JsiiProperty: "emailListInput", GoGetter: "EmailListInput"},
			_jsii_.MemberProperty{JsiiProperty: "everyone", GoGetter: "Everyone"},
			_jsii_.MemberProperty{JsiiProperty: "everyoneInput", GoGetter: "EveryoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalEvaluation", GoGetter: "ExternalEvaluation"},
			_jsii_.MemberProperty{JsiiProperty: "externalEvaluationInput", GoGetter: "ExternalEvaluationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "geo", GoGetter: "Geo"},
			_jsii_.MemberProperty{JsiiProperty: "geoInput", GoGetter: "GeoInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "githubOrganization", GoGetter: "GithubOrganization"},
			_jsii_.MemberProperty{JsiiProperty: "githubOrganizationInput", GoGetter: "GithubOrganizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
			_jsii_.MemberProperty{JsiiProperty: "groupInput", GoGetter: "GroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "gsuite", GoGetter: "Gsuite"},
			_jsii_.MemberProperty{JsiiProperty: "gsuiteInput", GoGetter: "GsuiteInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "ipInput", GoGetter: "IpInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipList", GoGetter: "IpList"},
			_jsii_.MemberProperty{JsiiProperty: "ipListInput", GoGetter: "IpListInput"},
			_jsii_.MemberProperty{JsiiProperty: "loginMethod", GoGetter: "LoginMethod"},
			_jsii_.MemberProperty{JsiiProperty: "loginMethodInput", GoGetter: "LoginMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidc", GoGetter: "Oidc"},
			_jsii_.MemberProperty{JsiiProperty: "oidcInput", GoGetter: "OidcInput"},
			_jsii_.MemberProperty{JsiiProperty: "okta", GoGetter: "Okta"},
			_jsii_.MemberProperty{JsiiProperty: "oktaInput", GoGetter: "OktaInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAnyValidServiceToken", GoMethod: "PutAnyValidServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthContext", GoMethod: "PutAuthContext"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthMethod", GoMethod: "PutAuthMethod"},
			_jsii_.MemberMethod{JsiiMethod: "putAzureAd", GoMethod: "PutAzureAd"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putCommonName", GoMethod: "PutCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "putDevicePosture", GoMethod: "PutDevicePosture"},
			_jsii_.MemberMethod{JsiiMethod: "putEmail", GoMethod: "PutEmail"},
			_jsii_.MemberMethod{JsiiMethod: "putEmailDomain", GoMethod: "PutEmailDomain"},
			_jsii_.MemberMethod{JsiiMethod: "putEmailList", GoMethod: "PutEmailList"},
			_jsii_.MemberMethod{JsiiMethod: "putEveryone", GoMethod: "PutEveryone"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalEvaluation", GoMethod: "PutExternalEvaluation"},
			_jsii_.MemberMethod{JsiiMethod: "putGeo", GoMethod: "PutGeo"},
			_jsii_.MemberMethod{JsiiMethod: "putGithubOrganization", GoMethod: "PutGithubOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "putGroup", GoMethod: "PutGroup"},
			_jsii_.MemberMethod{JsiiMethod: "putGsuite", GoMethod: "PutGsuite"},
			_jsii_.MemberMethod{JsiiMethod: "putIp", GoMethod: "PutIp"},
			_jsii_.MemberMethod{JsiiMethod: "putIpList", GoMethod: "PutIpList"},
			_jsii_.MemberMethod{JsiiMethod: "putLoginMethod", GoMethod: "PutLoginMethod"},
			_jsii_.MemberMethod{JsiiMethod: "putOidc", GoMethod: "PutOidc"},
			_jsii_.MemberMethod{JsiiMethod: "putOkta", GoMethod: "PutOkta"},
			_jsii_.MemberMethod{JsiiMethod: "putSaml", GoMethod: "PutSaml"},
			_jsii_.MemberMethod{JsiiMethod: "putServiceToken", GoMethod: "PutServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnyValidServiceToken", GoMethod: "ResetAnyValidServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthContext", GoMethod: "ResetAuthContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthMethod", GoMethod: "ResetAuthMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureAd", GoMethod: "ResetAzureAd"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommonName", GoMethod: "ResetCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDevicePosture", GoMethod: "ResetDevicePosture"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailDomain", GoMethod: "ResetEmailDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailList", GoMethod: "ResetEmailList"},
			_jsii_.MemberMethod{JsiiMethod: "resetEveryone", GoMethod: "ResetEveryone"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalEvaluation", GoMethod: "ResetExternalEvaluation"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeo", GoMethod: "ResetGeo"},
			_jsii_.MemberMethod{JsiiMethod: "resetGithubOrganization", GoMethod: "ResetGithubOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroup", GoMethod: "ResetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetGsuite", GoMethod: "ResetGsuite"},
			_jsii_.MemberMethod{JsiiMethod: "resetIp", GoMethod: "ResetIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpList", GoMethod: "ResetIpList"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginMethod", GoMethod: "ResetLoginMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidc", GoMethod: "ResetOidc"},
			_jsii_.MemberMethod{JsiiMethod: "resetOkta", GoMethod: "ResetOkta"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml", GoMethod: "ResetSaml"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceToken", GoMethod: "ResetServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "saml", GoGetter: "Saml"},
			_jsii_.MemberProperty{JsiiProperty: "samlInput", GoGetter: "SamlInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberProperty{JsiiProperty: "serviceTokenInput", GoGetter: "ServiceTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeSaml",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeSaml](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeSamlOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeSamlOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributeName", GoGetter: "AttributeName"},
			_jsii_.MemberProperty{JsiiProperty: "attributeNameInput", GoGetter: "AttributeNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValue", GoGetter: "AttributeValue"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValueInput", GoGetter: "AttributeValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeSamlOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeServiceToken",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeServiceToken](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesIncludeServiceTokenOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesIncludeServiceTokenOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenId", GoGetter: "TokenId"},
			_jsii_.MemberProperty{JsiiProperty: "tokenIdInput", GoGetter: "TokenIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesIncludeServiceTokenOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesList",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionRules", GoGetter: "ConnectionRules"},
			_jsii_.MemberProperty{JsiiProperty: "connectionRulesInput", GoGetter: "ConnectionRulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "decision", GoGetter: "Decision"},
			_jsii_.MemberProperty{JsiiProperty: "decisionInput", GoGetter: "DecisionInput"},
			_jsii_.MemberProperty{JsiiProperty: "exclude", GoGetter: "Exclude"},
			_jsii_.MemberProperty{JsiiProperty: "excludeInput", GoGetter: "ExcludeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "include", GoGetter: "Include"},
			_jsii_.MemberProperty{JsiiProperty: "includeInput", GoGetter: "IncludeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "precedence", GoGetter: "Precedence"},
			_jsii_.MemberProperty{JsiiProperty: "precedenceInput", GoGetter: "PrecedenceInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConnectionRules", GoMethod: "PutConnectionRules"},
			_jsii_.MemberMethod{JsiiMethod: "putExclude", GoMethod: "PutExclude"},
			_jsii_.MemberMethod{JsiiMethod: "putInclude", GoMethod: "PutInclude"},
			_jsii_.MemberMethod{JsiiMethod: "putRequire", GoMethod: "PutRequire"},
			_jsii_.MemberProperty{JsiiProperty: "require", GoGetter: "Require"},
			_jsii_.MemberProperty{JsiiProperty: "requireInput", GoGetter: "RequireInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionRules", GoMethod: "ResetConnectionRules"},
			_jsii_.MemberMethod{JsiiMethod: "resetDecision", GoMethod: "ResetDecision"},
			_jsii_.MemberMethod{JsiiMethod: "resetExclude", GoMethod: "ResetExclude"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInclude", GoMethod: "ResetInclude"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrecedence", GoMethod: "ResetPrecedence"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequire", GoMethod: "ResetRequire"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequire",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequire](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceToken",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceToken](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAnyValidServiceTokenOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAuthContext",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAuthContext](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acId", GoGetter: "AcId"},
			_jsii_.MemberProperty{JsiiProperty: "acIdInput", GoGetter: "AcIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthContextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAuthMethod",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAuthMethod](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authMethod", GoGetter: "AuthMethod"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodInput", GoGetter: "AuthMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAuthMethodOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAzureAd",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAzureAd](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireAzureAdOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireCertificate",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireCertificate](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireCertificateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireCommonName",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireCommonName](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireCommonNameOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireDevicePosture",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireDevicePosture](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "integrationUid", GoGetter: "IntegrationUid"},
			_jsii_.MemberProperty{JsiiProperty: "integrationUidInput", GoGetter: "IntegrationUidInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireDevicePostureOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEmail",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEmail](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEmailDomain",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEmailDomain](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireEmailDomainOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEmailListStruct",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEmailListStruct](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireEmailListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEmailOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEmailOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireEmailOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEveryone",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEveryone](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireEveryoneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireExternalEvaluation",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireExternalEvaluation](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateUrl", GoGetter: "EvaluateUrl"},
			_jsii_.MemberProperty{JsiiProperty: "evaluateUrlInput", GoGetter: "EvaluateUrlInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keysUrl", GoGetter: "KeysUrl"},
			_jsii_.MemberProperty{JsiiProperty: "keysUrlInput", GoGetter: "KeysUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireExternalEvaluationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGeo",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGeo](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGeoOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGeoOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "countryCode", GoGetter: "CountryCode"},
			_jsii_.MemberProperty{JsiiProperty: "countryCodeInput", GoGetter: "CountryCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireGeoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGithubOrganization",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGithubOrganization](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeam", GoMethod: "ResetTeam"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "team", GoGetter: "Team"},
			_jsii_.MemberProperty{JsiiProperty: "teamInput", GoGetter: "TeamInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireGithubOrganizationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGroup",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGroup](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGroupOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGroupOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireGroupOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGsuite",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGsuite](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireGsuiteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireIp",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireIp](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireIpListStruct",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireIpListStruct](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireIpListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireIpOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireIpOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "ipInput", GoGetter: "IpInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireIpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireList",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireLoginMethod",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireLoginMethod](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireLoginMethodOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireOidc",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireOidc](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireOidcOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireOidcOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "claimName", GoGetter: "ClaimName"},
			_jsii_.MemberProperty{JsiiProperty: "claimNameInput", GoGetter: "ClaimNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "claimValue", GoGetter: "ClaimValue"},
			_jsii_.MemberProperty{JsiiProperty: "claimValueInput", GoGetter: "ClaimValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOidcOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireOkta",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireOkta](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireOktaOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireOktaOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOktaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "anyValidServiceToken", GoGetter: "AnyValidServiceToken"},
			_jsii_.MemberProperty{JsiiProperty: "anyValidServiceTokenInput", GoGetter: "AnyValidServiceTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "authContext", GoGetter: "AuthContext"},
			_jsii_.MemberProperty{JsiiProperty: "authContextInput", GoGetter: "AuthContextInput"},
			_jsii_.MemberProperty{JsiiProperty: "authMethod", GoGetter: "AuthMethod"},
			_jsii_.MemberProperty{JsiiProperty: "authMethodInput", GoGetter: "AuthMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "azureAd", GoGetter: "AzureAd"},
			_jsii_.MemberProperty{JsiiProperty: "azureAdInput", GoGetter: "AzureAdInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificate", GoGetter: "Certificate"},
			_jsii_.MemberProperty{JsiiProperty: "certificateInput", GoGetter: "CertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "devicePosture", GoGetter: "DevicePosture"},
			_jsii_.MemberProperty{JsiiProperty: "devicePostureInput", GoGetter: "DevicePostureInput"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailDomain", GoGetter: "EmailDomain"},
			_jsii_.MemberProperty{JsiiProperty: "emailDomainInput", GoGetter: "EmailDomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "emailList", GoGetter: "EmailList"},
			_jsii_.MemberProperty{JsiiProperty: "emailListInput", GoGetter: "EmailListInput"},
			_jsii_.MemberProperty{JsiiProperty: "everyone", GoGetter: "Everyone"},
			_jsii_.MemberProperty{JsiiProperty: "everyoneInput", GoGetter: "EveryoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "externalEvaluation", GoGetter: "ExternalEvaluation"},
			_jsii_.MemberProperty{JsiiProperty: "externalEvaluationInput", GoGetter: "ExternalEvaluationInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "geo", GoGetter: "Geo"},
			_jsii_.MemberProperty{JsiiProperty: "geoInput", GoGetter: "GeoInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "githubOrganization", GoGetter: "GithubOrganization"},
			_jsii_.MemberProperty{JsiiProperty: "githubOrganizationInput", GoGetter: "GithubOrganizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "group", GoGetter: "Group"},
			_jsii_.MemberProperty{JsiiProperty: "groupInput", GoGetter: "GroupInput"},
			_jsii_.MemberProperty{JsiiProperty: "gsuite", GoGetter: "Gsuite"},
			_jsii_.MemberProperty{JsiiProperty: "gsuiteInput", GoGetter: "GsuiteInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "ipInput", GoGetter: "IpInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipList", GoGetter: "IpList"},
			_jsii_.MemberProperty{JsiiProperty: "ipListInput", GoGetter: "IpListInput"},
			_jsii_.MemberProperty{JsiiProperty: "loginMethod", GoGetter: "LoginMethod"},
			_jsii_.MemberProperty{JsiiProperty: "loginMethodInput", GoGetter: "LoginMethodInput"},
			_jsii_.MemberProperty{JsiiProperty: "oidc", GoGetter: "Oidc"},
			_jsii_.MemberProperty{JsiiProperty: "oidcInput", GoGetter: "OidcInput"},
			_jsii_.MemberProperty{JsiiProperty: "okta", GoGetter: "Okta"},
			_jsii_.MemberProperty{JsiiProperty: "oktaInput", GoGetter: "OktaInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAnyValidServiceToken", GoMethod: "PutAnyValidServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthContext", GoMethod: "PutAuthContext"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthMethod", GoMethod: "PutAuthMethod"},
			_jsii_.MemberMethod{JsiiMethod: "putAzureAd", GoMethod: "PutAzureAd"},
			_jsii_.MemberMethod{JsiiMethod: "putCertificate", GoMethod: "PutCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "putCommonName", GoMethod: "PutCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "putDevicePosture", GoMethod: "PutDevicePosture"},
			_jsii_.MemberMethod{JsiiMethod: "putEmail", GoMethod: "PutEmail"},
			_jsii_.MemberMethod{JsiiMethod: "putEmailDomain", GoMethod: "PutEmailDomain"},
			_jsii_.MemberMethod{JsiiMethod: "putEmailList", GoMethod: "PutEmailList"},
			_jsii_.MemberMethod{JsiiMethod: "putEveryone", GoMethod: "PutEveryone"},
			_jsii_.MemberMethod{JsiiMethod: "putExternalEvaluation", GoMethod: "PutExternalEvaluation"},
			_jsii_.MemberMethod{JsiiMethod: "putGeo", GoMethod: "PutGeo"},
			_jsii_.MemberMethod{JsiiMethod: "putGithubOrganization", GoMethod: "PutGithubOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "putGroup", GoMethod: "PutGroup"},
			_jsii_.MemberMethod{JsiiMethod: "putGsuite", GoMethod: "PutGsuite"},
			_jsii_.MemberMethod{JsiiMethod: "putIp", GoMethod: "PutIp"},
			_jsii_.MemberMethod{JsiiMethod: "putIpList", GoMethod: "PutIpList"},
			_jsii_.MemberMethod{JsiiMethod: "putLoginMethod", GoMethod: "PutLoginMethod"},
			_jsii_.MemberMethod{JsiiMethod: "putOidc", GoMethod: "PutOidc"},
			_jsii_.MemberMethod{JsiiMethod: "putOkta", GoMethod: "PutOkta"},
			_jsii_.MemberMethod{JsiiMethod: "putSaml", GoMethod: "PutSaml"},
			_jsii_.MemberMethod{JsiiMethod: "putServiceToken", GoMethod: "PutServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAnyValidServiceToken", GoMethod: "ResetAnyValidServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthContext", GoMethod: "ResetAuthContext"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthMethod", GoMethod: "ResetAuthMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetAzureAd", GoMethod: "ResetAzureAd"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificate", GoMethod: "ResetCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommonName", GoMethod: "ResetCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDevicePosture", GoMethod: "ResetDevicePosture"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailDomain", GoMethod: "ResetEmailDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailList", GoMethod: "ResetEmailList"},
			_jsii_.MemberMethod{JsiiMethod: "resetEveryone", GoMethod: "ResetEveryone"},
			_jsii_.MemberMethod{JsiiMethod: "resetExternalEvaluation", GoMethod: "ResetExternalEvaluation"},
			_jsii_.MemberMethod{JsiiMethod: "resetGeo", GoMethod: "ResetGeo"},
			_jsii_.MemberMethod{JsiiMethod: "resetGithubOrganization", GoMethod: "ResetGithubOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroup", GoMethod: "ResetGroup"},
			_jsii_.MemberMethod{JsiiMethod: "resetGsuite", GoMethod: "ResetGsuite"},
			_jsii_.MemberMethod{JsiiMethod: "resetIp", GoMethod: "ResetIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpList", GoMethod: "ResetIpList"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoginMethod", GoMethod: "ResetLoginMethod"},
			_jsii_.MemberMethod{JsiiMethod: "resetOidc", GoMethod: "ResetOidc"},
			_jsii_.MemberMethod{JsiiMethod: "resetOkta", GoMethod: "ResetOkta"},
			_jsii_.MemberMethod{JsiiMethod: "resetSaml", GoMethod: "ResetSaml"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceToken", GoMethod: "ResetServiceToken"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "saml", GoGetter: "Saml"},
			_jsii_.MemberProperty{JsiiProperty: "samlInput", GoGetter: "SamlInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceToken", GoGetter: "ServiceToken"},
			_jsii_.MemberProperty{JsiiProperty: "serviceTokenInput", GoGetter: "ServiceTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireSaml",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireSaml](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireSamlOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireSamlOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributeName", GoGetter: "AttributeName"},
			_jsii_.MemberProperty{JsiiProperty: "attributeNameInput", GoGetter: "AttributeNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValue", GoGetter: "AttributeValue"},
			_jsii_.MemberProperty{JsiiProperty: "attributeValueInput", GoGetter: "AttributeValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityProviderId", GoGetter: "IdentityProviderId"},
			_jsii_.MemberProperty{JsiiProperty: "identityProviderIdInput", GoGetter: "IdentityProviderIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireSamlOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireServiceToken",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireServiceToken](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tokenId", GoGetter: "TokenId"},
			_jsii_.MemberProperty{JsiiProperty: "tokenIdInput", GoGetter: "TokenIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationPoliciesRequireServiceTokenOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasApp",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasApp](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributes",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributes](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributesList",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributesList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomAttributesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributesOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributesOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyName", GoGetter: "FriendlyName"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyNameInput", GoGetter: "FriendlyNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameFormat", GoGetter: "NameFormat"},
			_jsii_.MemberProperty{JsiiProperty: "nameFormatInput", GoGetter: "NameFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSource", GoMethod: "PutSource"},
			_jsii_.MemberProperty{JsiiProperty: "required", GoGetter: "Required"},
			_jsii_.MemberProperty{JsiiProperty: "requiredInput", GoGetter: "RequiredInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetFriendlyName", GoMethod: "ResetFriendlyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNameFormat", GoMethod: "ResetNameFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequired", GoMethod: "ResetRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetSource", GoMethod: "ResetSource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomAttributesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributesSource",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributesSource](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdp",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdp](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdpList",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdpList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdpList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdpOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdpOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "idpId", GoGetter: "IdpId"},
			_jsii_.MemberProperty{JsiiProperty: "idpIdInput", GoGetter: "IdpIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdpId", GoMethod: "ResetIdpId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceName", GoMethod: "ResetSourceName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceName", GoGetter: "SourceName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceNameInput", GoGetter: "SourceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomAttributesSourceNameByIdpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomAttributesSourceOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomAttributesSourceOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameByIdp", GoGetter: "NameByIdp"},
			_jsii_.MemberProperty{JsiiProperty: "nameByIdpInput", GoGetter: "NameByIdpInput"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putNameByIdp", GoMethod: "PutNameByIdp"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNameByIdp", GoMethod: "ResetNameByIdp"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomAttributesSourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomClaims",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomClaims](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomClaimsList",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomClaimsList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomClaimsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomClaimsOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomClaimsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putSource", GoMethod: "PutSource"},
			_jsii_.MemberProperty{JsiiProperty: "required", GoGetter: "Required"},
			_jsii_.MemberProperty{JsiiProperty: "requiredInput", GoGetter: "RequiredInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequired", GoMethod: "ResetRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetScope", GoMethod: "ResetScope"},
			_jsii_.MemberMethod{JsiiMethod: "resetSource", GoMethod: "ResetSource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "scopeInput", GoGetter: "ScopeInput"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomClaimsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomClaimsSource",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomClaimsSource](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppCustomClaimsSourceOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppCustomClaimsSourceOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameByIdp", GoGetter: "NameByIdp"},
			_jsii_.MemberProperty{JsiiProperty: "nameByIdpInput", GoGetter: "NameByIdpInput"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNameByIdp", GoMethod: "ResetNameByIdp"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppCustomClaimsSourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptions",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptions](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetReturnAccessTokenFromAuthorizationEndpoint", GoMethod: "ResetReturnAccessTokenFromAuthorizationEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resetReturnIdTokenFromAuthorizationEndpoint", GoMethod: "ResetReturnIdTokenFromAuthorizationEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "returnAccessTokenFromAuthorizationEndpoint", GoGetter: "ReturnAccessTokenFromAuthorizationEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "returnAccessTokenFromAuthorizationEndpointInput", GoGetter: "ReturnAccessTokenFromAuthorizationEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "returnIdTokenFromAuthorizationEndpoint", GoGetter: "ReturnIdTokenFromAuthorizationEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "returnIdTokenFromAuthorizationEndpointInput", GoGetter: "ReturnIdTokenFromAuthorizationEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppHybridAndImplicitOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessTokenLifetime", GoGetter: "AccessTokenLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "accessTokenLifetimeInput", GoGetter: "AccessTokenLifetimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowPkceWithoutClientSecret", GoGetter: "AllowPkceWithoutClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "allowPkceWithoutClientSecretInput", GoGetter: "AllowPkceWithoutClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "appLauncherUrl", GoGetter: "AppLauncherUrl"},
			_jsii_.MemberProperty{JsiiProperty: "appLauncherUrlInput", GoGetter: "AppLauncherUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "authType", GoGetter: "AuthType"},
			_jsii_.MemberProperty{JsiiProperty: "authTypeInput", GoGetter: "AuthTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "consumerServiceUrl", GoGetter: "ConsumerServiceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "consumerServiceUrlInput", GoGetter: "ConsumerServiceUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customAttributes", GoGetter: "CustomAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "customAttributesInput", GoGetter: "CustomAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "customClaims", GoGetter: "CustomClaims"},
			_jsii_.MemberProperty{JsiiProperty: "customClaimsInput", GoGetter: "CustomClaimsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRelayState", GoGetter: "DefaultRelayState"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRelayStateInput", GoGetter: "DefaultRelayStateInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "grantTypes", GoGetter: "GrantTypes"},
			_jsii_.MemberProperty{JsiiProperty: "grantTypesInput", GoGetter: "GrantTypesInput"},
			_jsii_.MemberProperty{JsiiProperty: "groupFilterRegex", GoGetter: "GroupFilterRegex"},
			_jsii_.MemberProperty{JsiiProperty: "groupFilterRegexInput", GoGetter: "GroupFilterRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "hybridAndImplicitOptions", GoGetter: "HybridAndImplicitOptions"},
			_jsii_.MemberProperty{JsiiProperty: "hybridAndImplicitOptionsInput", GoGetter: "HybridAndImplicitOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "idpEntityId", GoGetter: "IdpEntityId"},
			_jsii_.MemberProperty{JsiiProperty: "idpEntityIdInput", GoGetter: "IdpEntityIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "nameIdFormat", GoGetter: "NameIdFormat"},
			_jsii_.MemberProperty{JsiiProperty: "nameIdFormatInput", GoGetter: "NameIdFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "nameIdTransformJsonata", GoGetter: "NameIdTransformJsonata"},
			_jsii_.MemberProperty{JsiiProperty: "nameIdTransformJsonataInput", GoGetter: "NameIdTransformJsonataInput"},
			_jsii_.MemberProperty{JsiiProperty: "publicKey", GoGetter: "PublicKey"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomAttributes", GoMethod: "PutCustomAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomClaims", GoMethod: "PutCustomClaims"},
			_jsii_.MemberMethod{JsiiMethod: "putHybridAndImplicitOptions", GoMethod: "PutHybridAndImplicitOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putRefreshTokenOptions", GoMethod: "PutRefreshTokenOptions"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUris", GoGetter: "RedirectUris"},
			_jsii_.MemberProperty{JsiiProperty: "redirectUrisInput", GoGetter: "RedirectUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenOptions", GoGetter: "RefreshTokenOptions"},
			_jsii_.MemberProperty{JsiiProperty: "refreshTokenOptionsInput", GoGetter: "RefreshTokenOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessTokenLifetime", GoMethod: "ResetAccessTokenLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowPkceWithoutClientSecret", GoMethod: "ResetAllowPkceWithoutClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLauncherUrl", GoMethod: "ResetAppLauncherUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthType", GoMethod: "ResetAuthType"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsumerServiceUrl", GoMethod: "ResetConsumerServiceUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomAttributes", GoMethod: "ResetCustomAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomClaims", GoMethod: "ResetCustomClaims"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRelayState", GoMethod: "ResetDefaultRelayState"},
			_jsii_.MemberMethod{JsiiMethod: "resetGrantTypes", GoMethod: "ResetGrantTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupFilterRegex", GoMethod: "ResetGroupFilterRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetHybridAndImplicitOptions", GoMethod: "ResetHybridAndImplicitOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdpEntityId", GoMethod: "ResetIdpEntityId"},
			_jsii_.MemberMethod{JsiiMethod: "resetNameIdFormat", GoMethod: "ResetNameIdFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetNameIdTransformJsonata", GoMethod: "ResetNameIdTransformJsonata"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedirectUris", GoMethod: "ResetRedirectUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetRefreshTokenOptions", GoMethod: "ResetRefreshTokenOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetSamlAttributeTransformJsonata", GoMethod: "ResetSamlAttributeTransformJsonata"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpEntityId", GoMethod: "ResetSpEntityId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSsoEndpoint", GoMethod: "ResetSsoEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "samlAttributeTransformJsonata", GoGetter: "SamlAttributeTransformJsonata"},
			_jsii_.MemberProperty{JsiiProperty: "samlAttributeTransformJsonataInput", GoGetter: "SamlAttributeTransformJsonataInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "spEntityId", GoGetter: "SpEntityId"},
			_jsii_.MemberProperty{JsiiProperty: "spEntityIdInput", GoGetter: "SpEntityIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ssoEndpoint", GoGetter: "SsoEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "ssoEndpointInput", GoGetter: "SsoEndpointInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppRefreshTokenOptions",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppRefreshTokenOptions](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifetime", GoGetter: "Lifetime"},
			_jsii_.MemberProperty{JsiiProperty: "lifetimeInput", GoGetter: "LifetimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifetime", GoMethod: "ResetLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationSaasAppRefreshTokenOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfig",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfig](),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigAuthentication",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigAuthentication](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigAuthenticationOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigAuthenticationOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authorizationUrl", GoGetter: "AuthorizationUrl"},
			_jsii_.MemberProperty{JsiiProperty: "authorizationUrlInput", GoGetter: "AuthorizationUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientId", GoGetter: "ClientId"},
			_jsii_.MemberProperty{JsiiProperty: "clientIdInput", GoGetter: "ClientIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecret", GoGetter: "ClientSecret"},
			_jsii_.MemberProperty{JsiiProperty: "clientSecretInput", GoGetter: "ClientSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthorizationUrl", GoMethod: "ResetAuthorizationUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientId", GoMethod: "ResetClientId"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientSecret", GoMethod: "ResetClientSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetScopes", GoMethod: "ResetScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetTokenUrl", GoMethod: "ResetTokenUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetUser", GoMethod: "ResetUser"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scheme", GoGetter: "Scheme"},
			_jsii_.MemberProperty{JsiiProperty: "schemeInput", GoGetter: "SchemeInput"},
			_jsii_.MemberProperty{JsiiProperty: "scopes", GoGetter: "Scopes"},
			_jsii_.MemberProperty{JsiiProperty: "scopesInput", GoGetter: "ScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "tokenUrl", GoGetter: "TokenUrl"},
			_jsii_.MemberProperty{JsiiProperty: "tokenUrlInput", GoGetter: "TokenUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "user", GoGetter: "User"},
			_jsii_.MemberProperty{JsiiProperty: "userInput", GoGetter: "UserInput"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationScimConfigAuthenticationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigMappings",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigMappings](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigMappingsList",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigMappingsList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationScimConfigMappingsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigMappingsOperations",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigMappingsOperations](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigMappingsOperationsOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigMappingsOperationsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationScimConfigMappingsOperationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigMappingsOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigMappingsOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "filter", GoGetter: "Filter"},
			_jsii_.MemberProperty{JsiiProperty: "filterInput", GoGetter: "FilterInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "operations", GoGetter: "Operations"},
			_jsii_.MemberProperty{JsiiProperty: "operationsInput", GoGetter: "OperationsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOperations", GoMethod: "PutOperations"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilter", GoMethod: "ResetFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperations", GoMethod: "ResetOperations"},
			_jsii_.MemberMethod{JsiiMethod: "resetStrictness", GoMethod: "ResetStrictness"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransformJsonata", GoMethod: "ResetTransformJsonata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "strictness", GoGetter: "Strictness"},
			_jsii_.MemberProperty{JsiiProperty: "strictnessInput", GoGetter: "StrictnessInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "transformJsonata", GoGetter: "TransformJsonata"},
			_jsii_.MemberProperty{JsiiProperty: "transformJsonataInput", GoGetter: "TransformJsonataInput"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationScimConfigMappingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationScimConfigOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationScimConfigOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "authentication", GoGetter: "Authentication"},
			_jsii_.MemberProperty{JsiiProperty: "authenticationInput", GoGetter: "AuthenticationInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deactivateOnDelete", GoGetter: "DeactivateOnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "deactivateOnDeleteInput", GoGetter: "DeactivateOnDeleteInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idpUid", GoGetter: "IdpUid"},
			_jsii_.MemberProperty{JsiiProperty: "idpUidInput", GoGetter: "IdpUidInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "mappings", GoGetter: "Mappings"},
			_jsii_.MemberProperty{JsiiProperty: "mappingsInput", GoGetter: "MappingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAuthentication", GoMethod: "PutAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "putMappings", GoMethod: "PutMappings"},
			_jsii_.MemberProperty{JsiiProperty: "remoteUri", GoGetter: "RemoteUri"},
			_jsii_.MemberProperty{JsiiProperty: "remoteUriInput", GoGetter: "RemoteUriInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthentication", GoMethod: "ResetAuthentication"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeactivateOnDelete", GoMethod: "ResetDeactivateOnDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetMappings", GoMethod: "ResetMappings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationScimConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationTargetCriteria",
		reflect.TypeFor[ZeroTrustAccessApplicationTargetCriteria](),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationTargetCriteriaList",
		reflect.TypeFor[ZeroTrustAccessApplicationTargetCriteriaList](),
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
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationTargetCriteriaList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplicationTargetCriteriaOutputReference",
		reflect.TypeFor[ZeroTrustAccessApplicationTargetCriteriaOutputReference](),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "port", GoGetter: "Port"},
			_jsii_.MemberProperty{JsiiProperty: "portInput", GoGetter: "PortInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetAttributes", GoGetter: "TargetAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetAttributesInput", GoGetter: "TargetAttributesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() any {
			j := jsiiProxy_ZeroTrustAccessApplicationTargetCriteriaOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
