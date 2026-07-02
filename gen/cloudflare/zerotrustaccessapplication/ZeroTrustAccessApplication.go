package zerotrustaccessapplication

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustaccessapplication/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application cloudflare_zero_trust_access_application}.
type ZeroTrustAccessApplication interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowAuthenticateViaWarp() any
	SetAllowAuthenticateViaWarp(val any)
	AllowAuthenticateViaWarpInput() any
	AllowedIdps() *[]*string
	SetAllowedIdps(val *[]*string)
	AllowedIdpsInput() *[]*string
	AllowIframe() any
	SetAllowIframe(val any)
	AllowIframeInput() any
	AppLauncherLogoUrl() *string
	SetAppLauncherLogoUrl(val *string)
	AppLauncherLogoUrlInput() *string
	AppLauncherVisible() any
	SetAppLauncherVisible(val any)
	AppLauncherVisibleInput() any
	Aud() *string
	AutoRedirectToIdentity() any
	SetAutoRedirectToIdentity(val any)
	AutoRedirectToIdentityInput() any
	BgColor() *string
	SetBgColor(val *string)
	BgColorInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	CorsHeaders() ZeroTrustAccessApplicationCorsHeadersOutputReference
	CorsHeadersInput() any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CustomDenyMessage() *string
	SetCustomDenyMessage(val *string)
	CustomDenyMessageInput() *string
	CustomDenyUrl() *string
	SetCustomDenyUrl(val *string)
	CustomDenyUrlInput() *string
	CustomNonIdentityDenyUrl() *string
	SetCustomNonIdentityDenyUrl(val *string)
	CustomNonIdentityDenyUrlInput() *string
	CustomPages() *[]*string
	SetCustomPages(val *[]*string)
	CustomPagesInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Destinations() ZeroTrustAccessApplicationDestinationsList
	DestinationsInput() any
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	EnableBindingCookie() any
	SetEnableBindingCookie(val any)
	EnableBindingCookieInput() any
	FooterLinks() ZeroTrustAccessApplicationFooterLinksList
	FooterLinksInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HeaderBgColor() *string
	SetHeaderBgColor(val *string)
	HeaderBgColorInput() *string
	HttpOnlyCookieAttribute() any
	SetHttpOnlyCookieAttribute(val any)
	HttpOnlyCookieAttributeInput() any
	Id() *string
	LandingPageDesign() ZeroTrustAccessApplicationLandingPageDesignOutputReference
	LandingPageDesignInput() any
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogoUrl() *string
	SetLogoUrl(val *string)
	LogoUrlInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OptionsPreflightBypass() any
	SetOptionsPreflightBypass(val any)
	OptionsPreflightBypassInput() any
	PathCookieAttribute() any
	SetPathCookieAttribute(val any)
	PathCookieAttributeInput() any
	Policies() ZeroTrustAccessApplicationPoliciesList
	PoliciesInput() any
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	// Experimental.
	RawOverrides() any
	ReadServiceTokensFromHeader() *string
	SetReadServiceTokensFromHeader(val *string)
	ReadServiceTokensFromHeaderInput() *string
	SaasApp() ZeroTrustAccessApplicationSaasAppOutputReference
	SaasAppInput() any
	SameSiteCookieAttribute() *string
	SetSameSiteCookieAttribute(val *string)
	SameSiteCookieAttributeInput() *string
	ScimConfig() ZeroTrustAccessApplicationScimConfigOutputReference
	ScimConfigInput() any
	SelfHostedDomains() *[]*string
	SetSelfHostedDomains(val *[]*string)
	SelfHostedDomainsInput() *[]*string
	ServiceAuth401Redirect() any
	SetServiceAuth401Redirect(val any)
	ServiceAuth401RedirectInput() any
	SessionDuration() *string
	SetSessionDuration(val *string)
	SessionDurationInput() *string
	SkipAppLauncherLoginPage() any
	SetSkipAppLauncherLoginPage(val any)
	SkipAppLauncherLoginPageInput() any
	SkipInterstitial() any
	SetSkipInterstitial(val any)
	SkipInterstitialInput() any
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	TargetCriteria() ZeroTrustAccessApplicationTargetCriteriaList
	TargetCriteriaInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ZoneId() *string
	SetZoneId(val *string)
	ZoneIdInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() any
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCorsHeaders(value *ZeroTrustAccessApplicationCorsHeaders)
	PutDestinations(value any)
	PutFooterLinks(value any)
	PutLandingPageDesign(value *ZeroTrustAccessApplicationLandingPageDesign)
	PutPolicies(value any)
	PutSaasApp(value *ZeroTrustAccessApplicationSaasApp)
	PutScimConfig(value *ZeroTrustAccessApplicationScimConfig)
	PutTargetCriteria(value any)
	ResetAccountId()
	ResetAllowAuthenticateViaWarp()
	ResetAllowedIdps()
	ResetAllowIframe()
	ResetAppLauncherLogoUrl()
	ResetAppLauncherVisible()
	ResetAutoRedirectToIdentity()
	ResetBgColor()
	ResetCorsHeaders()
	ResetCustomDenyMessage()
	ResetCustomDenyUrl()
	ResetCustomNonIdentityDenyUrl()
	ResetCustomPages()
	ResetDestinations()
	ResetDomain()
	ResetEnableBindingCookie()
	ResetFooterLinks()
	ResetHeaderBgColor()
	ResetHttpOnlyCookieAttribute()
	ResetLandingPageDesign()
	ResetLogoUrl()
	ResetName()
	ResetOptionsPreflightBypass()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathCookieAttribute()
	ResetPolicies()
	ResetReadServiceTokensFromHeader()
	ResetSaasApp()
	ResetSameSiteCookieAttribute()
	ResetScimConfig()
	ResetSelfHostedDomains()
	ResetServiceAuth401Redirect()
	ResetSessionDuration()
	ResetSkipAppLauncherLoginPage()
	ResetSkipInterstitial()
	ResetTags()
	ResetTargetCriteria()
	ResetType()
	ResetZoneId()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for ZeroTrustAccessApplication
type jsiiProxy_ZeroTrustAccessApplication struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowAuthenticateViaWarp() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAuthenticateViaWarp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowAuthenticateViaWarpInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowAuthenticateViaWarpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowedIdps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowedIdpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedIdpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowIframe() any {
	var returns any
	_jsii_.Get(
		j,
		"allowIframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AllowIframeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowIframeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherLogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherLogoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherLogoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appLauncherLogoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherVisible() any {
	var returns any
	_jsii_.Get(
		j,
		"appLauncherVisible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AppLauncherVisibleInput() any {
	var returns any
	_jsii_.Get(
		j,
		"appLauncherVisibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Aud() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aud",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AutoRedirectToIdentity() any {
	var returns any
	_jsii_.Get(
		j,
		"autoRedirectToIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) AutoRedirectToIdentityInput() any {
	var returns any
	_jsii_.Get(
		j,
		"autoRedirectToIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) BgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) BgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CorsHeaders() ZeroTrustAccessApplicationCorsHeadersOutputReference {
	var returns ZeroTrustAccessApplicationCorsHeadersOutputReference
	_jsii_.Get(
		j,
		"corsHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CorsHeadersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"corsHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomDenyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDenyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomNonIdentityDenyUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customNonIdentityDenyUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomNonIdentityDenyUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customNonIdentityDenyUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomPages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) CustomPagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Destinations() ZeroTrustAccessApplicationDestinationsList {
	var returns ZeroTrustAccessApplicationDestinationsList
	_jsii_.Get(
		j,
		"destinations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DestinationsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"destinationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) EnableBindingCookie() any {
	var returns any
	_jsii_.Get(
		j,
		"enableBindingCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) EnableBindingCookieInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enableBindingCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) FooterLinks() ZeroTrustAccessApplicationFooterLinksList {
	var returns ZeroTrustAccessApplicationFooterLinksList
	_jsii_.Get(
		j,
		"footerLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) FooterLinksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"footerLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HeaderBgColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBgColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HeaderBgColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headerBgColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HttpOnlyCookieAttribute() any {
	var returns any
	_jsii_.Get(
		j,
		"httpOnlyCookieAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) HttpOnlyCookieAttributeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"httpOnlyCookieAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LandingPageDesign() ZeroTrustAccessApplicationLandingPageDesignOutputReference {
	var returns ZeroTrustAccessApplicationLandingPageDesignOutputReference
	_jsii_.Get(
		j,
		"landingPageDesign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LandingPageDesignInput() any {
	var returns any
	_jsii_.Get(
		j,
		"landingPageDesignInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LogoUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) LogoUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) OptionsPreflightBypass() any {
	var returns any
	_jsii_.Get(
		j,
		"optionsPreflightBypass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) OptionsPreflightBypassInput() any {
	var returns any
	_jsii_.Get(
		j,
		"optionsPreflightBypassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) PathCookieAttribute() any {
	var returns any
	_jsii_.Get(
		j,
		"pathCookieAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) PathCookieAttributeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pathCookieAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Policies() ZeroTrustAccessApplicationPoliciesList {
	var returns ZeroTrustAccessApplicationPoliciesList
	_jsii_.Get(
		j,
		"policies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) PoliciesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"policiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ReadServiceTokensFromHeader() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readServiceTokensFromHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ReadServiceTokensFromHeaderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readServiceTokensFromHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SaasApp() ZeroTrustAccessApplicationSaasAppOutputReference {
	var returns ZeroTrustAccessApplicationSaasAppOutputReference
	_jsii_.Get(
		j,
		"saasApp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SaasAppInput() any {
	var returns any
	_jsii_.Get(
		j,
		"saasAppInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SameSiteCookieAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sameSiteCookieAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SameSiteCookieAttributeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sameSiteCookieAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ScimConfig() ZeroTrustAccessApplicationScimConfigOutputReference {
	var returns ZeroTrustAccessApplicationScimConfigOutputReference
	_jsii_.Get(
		j,
		"scimConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ScimConfigInput() any {
	var returns any
	_jsii_.Get(
		j,
		"scimConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SelfHostedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selfHostedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SelfHostedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selfHostedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ServiceAuth401Redirect() any {
	var returns any
	_jsii_.Get(
		j,
		"serviceAuth401Redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ServiceAuth401RedirectInput() any {
	var returns any
	_jsii_.Get(
		j,
		"serviceAuth401RedirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SessionDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SessionDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipAppLauncherLoginPage() any {
	var returns any
	_jsii_.Get(
		j,
		"skipAppLauncherLoginPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipAppLauncherLoginPageInput() any {
	var returns any
	_jsii_.Get(
		j,
		"skipAppLauncherLoginPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipInterstitial() any {
	var returns any
	_jsii_.Get(
		j,
		"skipInterstitial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SkipInterstitialInput() any {
	var returns any
	_jsii_.Get(
		j,
		"skipInterstitialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TargetCriteria() ZeroTrustAccessApplicationTargetCriteriaList {
	var returns ZeroTrustAccessApplicationTargetCriteriaList
	_jsii_.Get(
		j,
		"targetCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TargetCriteriaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"targetCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustAccessApplication) ZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneIdInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application cloudflare_zero_trust_access_application} Resource.
func NewZeroTrustAccessApplication(scope constructs.Construct, id *string, config *ZeroTrustAccessApplicationConfig) ZeroTrustAccessApplication {
	_init_.Initialize()

	if err := validateNewZeroTrustAccessApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustAccessApplication{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_access_application cloudflare_zero_trust_access_application} Resource.
func NewZeroTrustAccessApplication_Override(z ZeroTrustAccessApplication, scope constructs.Construct, id *string, config *ZeroTrustAccessApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		[]any{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetAllowAuthenticateViaWarp(val any) {
	if err := j.validateSetAllowAuthenticateViaWarpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAuthenticateViaWarp",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetAllowedIdps(val *[]*string) {
	if err := j.validateSetAllowedIdpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedIdps",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetAllowIframe(val any) {
	if err := j.validateSetAllowIframeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowIframe",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetAppLauncherLogoUrl(val *string) {
	if err := j.validateSetAppLauncherLogoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLauncherLogoUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetAppLauncherVisible(val any) {
	if err := j.validateSetAppLauncherVisibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appLauncherVisible",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetAutoRedirectToIdentity(val any) {
	if err := j.validateSetAutoRedirectToIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRedirectToIdentity",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetBgColor(val *string) {
	if err := j.validateSetBgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bgColor",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetCustomDenyMessage(val *string) {
	if err := j.validateSetCustomDenyMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDenyMessage",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetCustomDenyUrl(val *string) {
	if err := j.validateSetCustomDenyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDenyUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetCustomNonIdentityDenyUrl(val *string) {
	if err := j.validateSetCustomNonIdentityDenyUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customNonIdentityDenyUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetCustomPages(val *[]*string) {
	if err := j.validateSetCustomPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPages",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetEnableBindingCookie(val any) {
	if err := j.validateSetEnableBindingCookieParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableBindingCookie",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetHeaderBgColor(val *string) {
	if err := j.validateSetHeaderBgColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerBgColor",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetHttpOnlyCookieAttribute(val any) {
	if err := j.validateSetHttpOnlyCookieAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpOnlyCookieAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetLogoUrl(val *string) {
	if err := j.validateSetLogoUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetOptionsPreflightBypass(val any) {
	if err := j.validateSetOptionsPreflightBypassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionsPreflightBypass",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetPathCookieAttribute(val any) {
	if err := j.validateSetPathCookieAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathCookieAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetReadServiceTokensFromHeader(val *string) {
	if err := j.validateSetReadServiceTokensFromHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readServiceTokensFromHeader",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetSameSiteCookieAttribute(val *string) {
	if err := j.validateSetSameSiteCookieAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sameSiteCookieAttribute",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetSelfHostedDomains(val *[]*string) {
	if err := j.validateSetSelfHostedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selfHostedDomains",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetServiceAuth401Redirect(val any) {
	if err := j.validateSetServiceAuth401RedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAuth401Redirect",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetSessionDuration(val *string) {
	if err := j.validateSetSessionDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionDuration",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetSkipAppLauncherLoginPage(val any) {
	if err := j.validateSetSkipAppLauncherLoginPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipAppLauncherLoginPage",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetSkipInterstitial(val any) {
	if err := j.validateSetSkipInterstitialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipInterstitial",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustAccessApplication) SetZoneId(val *string) {
	if err := j.validateSetZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneId",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustAccessApplication resource upon running "cdktf plan <stack-name>".
func ZeroTrustAccessApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ZeroTrustAccessApplication_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessApplication_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustAccessApplication_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustAccessApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ZeroTrustAccessApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustAccessApplication.ZeroTrustAccessApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) AddOverride(path *string, value any) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]any{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := z.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := z.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		z,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := z.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		z,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := z.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		z,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := z.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		z,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := z.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		z,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetStringAttribute(terraformAttribute *string) *string {
	if err := z.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		z,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := z.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		z,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]any{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := z.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		z,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) MoveTo(moveTarget *string, index any) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutCorsHeaders(value *ZeroTrustAccessApplicationCorsHeaders) {
	if err := z.validatePutCorsHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putCorsHeaders",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutDestinations(value any) {
	if err := z.validatePutDestinationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putDestinations",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutFooterLinks(value any) {
	if err := z.validatePutFooterLinksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putFooterLinks",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutLandingPageDesign(value *ZeroTrustAccessApplicationLandingPageDesign) {
	if err := z.validatePutLandingPageDesignParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putLandingPageDesign",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutPolicies(value any) {
	if err := z.validatePutPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putPolicies",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutSaasApp(value *ZeroTrustAccessApplicationSaasApp) {
	if err := z.validatePutSaasAppParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSaasApp",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutScimConfig(value *ZeroTrustAccessApplicationScimConfig) {
	if err := z.validatePutScimConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putScimConfig",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) PutTargetCriteria(value any) {
	if err := z.validatePutTargetCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putTargetCriteria",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAccountId() {
	_jsii_.InvokeVoid(
		z,
		"resetAccountId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAllowAuthenticateViaWarp() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowAuthenticateViaWarp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAllowedIdps() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedIdps",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAllowIframe() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowIframe",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAppLauncherLogoUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetAppLauncherLogoUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAppLauncherVisible() {
	_jsii_.InvokeVoid(
		z,
		"resetAppLauncherVisible",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetAutoRedirectToIdentity() {
	_jsii_.InvokeVoid(
		z,
		"resetAutoRedirectToIdentity",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetBgColor() {
	_jsii_.InvokeVoid(
		z,
		"resetBgColor",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCorsHeaders() {
	_jsii_.InvokeVoid(
		z,
		"resetCorsHeaders",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomDenyMessage() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomDenyMessage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomDenyUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomDenyUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomNonIdentityDenyUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomNonIdentityDenyUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetCustomPages() {
	_jsii_.InvokeVoid(
		z,
		"resetCustomPages",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetDestinations() {
	_jsii_.InvokeVoid(
		z,
		"resetDestinations",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetDomain() {
	_jsii_.InvokeVoid(
		z,
		"resetDomain",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetEnableBindingCookie() {
	_jsii_.InvokeVoid(
		z,
		"resetEnableBindingCookie",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetFooterLinks() {
	_jsii_.InvokeVoid(
		z,
		"resetFooterLinks",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetHeaderBgColor() {
	_jsii_.InvokeVoid(
		z,
		"resetHeaderBgColor",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetHttpOnlyCookieAttribute() {
	_jsii_.InvokeVoid(
		z,
		"resetHttpOnlyCookieAttribute",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetLandingPageDesign() {
	_jsii_.InvokeVoid(
		z,
		"resetLandingPageDesign",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetLogoUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetLogoUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetName() {
	_jsii_.InvokeVoid(
		z,
		"resetName",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetOptionsPreflightBypass() {
	_jsii_.InvokeVoid(
		z,
		"resetOptionsPreflightBypass",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetPathCookieAttribute() {
	_jsii_.InvokeVoid(
		z,
		"resetPathCookieAttribute",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetPolicies() {
	_jsii_.InvokeVoid(
		z,
		"resetPolicies",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetReadServiceTokensFromHeader() {
	_jsii_.InvokeVoid(
		z,
		"resetReadServiceTokensFromHeader",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSaasApp() {
	_jsii_.InvokeVoid(
		z,
		"resetSaasApp",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSameSiteCookieAttribute() {
	_jsii_.InvokeVoid(
		z,
		"resetSameSiteCookieAttribute",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetScimConfig() {
	_jsii_.InvokeVoid(
		z,
		"resetScimConfig",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSelfHostedDomains() {
	_jsii_.InvokeVoid(
		z,
		"resetSelfHostedDomains",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetServiceAuth401Redirect() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceAuth401Redirect",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSessionDuration() {
	_jsii_.InvokeVoid(
		z,
		"resetSessionDuration",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSkipAppLauncherLoginPage() {
	_jsii_.InvokeVoid(
		z,
		"resetSkipAppLauncherLoginPage",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetSkipInterstitial() {
	_jsii_.InvokeVoid(
		z,
		"resetSkipInterstitial",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetTags() {
	_jsii_.InvokeVoid(
		z,
		"resetTags",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetTargetCriteria() {
	_jsii_.InvokeVoid(
		z,
		"resetTargetCriteria",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetType() {
	_jsii_.InvokeVoid(
		z,
		"resetType",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ResetZoneId() {
	_jsii_.InvokeVoid(
		z,
		"resetZoneId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustAccessApplication) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustAccessApplication) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
