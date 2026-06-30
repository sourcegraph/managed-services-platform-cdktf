package zerotrustdevicecustomprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustdevicecustomprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_custom_profile cloudflare_zero_trust_device_custom_profile}.
type ZeroTrustDeviceCustomProfile interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AllowedToLeave() any
	SetAllowedToLeave(val any)
	AllowedToLeaveInput() any
	AllowModeSwitch() any
	SetAllowModeSwitch(val any)
	AllowModeSwitchInput() any
	AllowUpdates() any
	SetAllowUpdates(val any)
	AllowUpdatesInput() any
	AutoConnect() *float64
	SetAutoConnect(val *float64)
	AutoConnectInput() *float64
	CaptivePortal() *float64
	SetCaptivePortal(val *float64)
	CaptivePortalInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	Default() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableAutoFallback() any
	SetDisableAutoFallback(val any)
	DisableAutoFallbackInput() any
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	Exclude() ZeroTrustDeviceCustomProfileExcludeList
	ExcludeInput() any
	ExcludeOfficeIps() any
	SetExcludeOfficeIps(val any)
	ExcludeOfficeIpsInput() any
	FallbackDomains() ZeroTrustDeviceCustomProfileFallbackDomainsList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayUniqueId() *string
	Id() *string
	Include() ZeroTrustDeviceCustomProfileIncludeList
	IncludeInput() any
	LanAllowMinutes() *float64
	SetLanAllowMinutes(val *float64)
	LanAllowMinutesInput() *float64
	LanAllowSubnetSize() *float64
	SetLanAllowSubnetSize(val *float64)
	LanAllowSubnetSizeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Match() *string
	SetMatch(val *string)
	MatchInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PolicyId() *string
	Precedence() *float64
	SetPrecedence(val *float64)
	PrecedenceInput() *float64
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
	RegisterInterfaceIpWithDns() any
	SetRegisterInterfaceIpWithDns(val any)
	RegisterInterfaceIpWithDnsInput() any
	SccmVpnBoundarySupport() any
	SetSccmVpnBoundarySupport(val any)
	SccmVpnBoundarySupportInput() any
	ServiceModeV2() ZeroTrustDeviceCustomProfileServiceModeV2OutputReference
	ServiceModeV2Input() any
	SupportUrl() *string
	SetSupportUrl(val *string)
	SupportUrlInput() *string
	SwitchLocked() any
	SetSwitchLocked(val any)
	SwitchLockedInput() any
	TargetTests() ZeroTrustDeviceCustomProfileTargetTestsList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	TunnelProtocol() *string
	SetTunnelProtocol(val *string)
	TunnelProtocolInput() *string
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
	PutExclude(value any)
	PutInclude(value any)
	PutServiceModeV2(value *ZeroTrustDeviceCustomProfileServiceModeV2)
	ResetAllowedToLeave()
	ResetAllowModeSwitch()
	ResetAllowUpdates()
	ResetAutoConnect()
	ResetCaptivePortal()
	ResetDescription()
	ResetDisableAutoFallback()
	ResetEnabled()
	ResetExclude()
	ResetExcludeOfficeIps()
	ResetInclude()
	ResetLanAllowMinutes()
	ResetLanAllowSubnetSize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegisterInterfaceIpWithDns()
	ResetSccmVpnBoundarySupport()
	ResetServiceModeV2()
	ResetSupportUrl()
	ResetSwitchLocked()
	ResetTunnelProtocol()
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

// The jsii proxy struct for ZeroTrustDeviceCustomProfile
type jsiiProxy_ZeroTrustDeviceCustomProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AllowedToLeave() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedToLeave",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AllowedToLeaveInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowedToLeaveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AllowModeSwitch() any {
	var returns any
	_jsii_.Get(
		j,
		"allowModeSwitch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AllowModeSwitchInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowModeSwitchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AllowUpdates() any {
	var returns any
	_jsii_.Get(
		j,
		"allowUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AllowUpdatesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allowUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AutoConnect() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) AutoConnectInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoConnectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) CaptivePortal() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) CaptivePortalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"captivePortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Default() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) DisableAutoFallback() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAutoFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) DisableAutoFallbackInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableAutoFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Exclude() ZeroTrustDeviceCustomProfileExcludeList {
	var returns ZeroTrustDeviceCustomProfileExcludeList
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) ExcludeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) ExcludeOfficeIps() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeOfficeIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) ExcludeOfficeIpsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"excludeOfficeIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) FallbackDomains() ZeroTrustDeviceCustomProfileFallbackDomainsList {
	var returns ZeroTrustDeviceCustomProfileFallbackDomainsList
	_jsii_.Get(
		j,
		"fallbackDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) GatewayUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Include() ZeroTrustDeviceCustomProfileIncludeList {
	var returns ZeroTrustDeviceCustomProfileIncludeList
	_jsii_.Get(
		j,
		"include",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) IncludeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) LanAllowMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) LanAllowMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) LanAllowSubnetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowSubnetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) LanAllowSubnetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lanAllowSubnetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Match() *string {
	var returns *string
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) MatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) PolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Precedence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) PrecedenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"precedenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) RegisterInterfaceIpWithDns() any {
	var returns any
	_jsii_.Get(
		j,
		"registerInterfaceIpWithDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) RegisterInterfaceIpWithDnsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"registerInterfaceIpWithDnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SccmVpnBoundarySupport() any {
	var returns any
	_jsii_.Get(
		j,
		"sccmVpnBoundarySupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SccmVpnBoundarySupportInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sccmVpnBoundarySupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) ServiceModeV2() ZeroTrustDeviceCustomProfileServiceModeV2OutputReference {
	var returns ZeroTrustDeviceCustomProfileServiceModeV2OutputReference
	_jsii_.Get(
		j,
		"serviceModeV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) ServiceModeV2Input() any {
	var returns any
	_jsii_.Get(
		j,
		"serviceModeV2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SupportUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SupportUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SwitchLocked() any {
	var returns any
	_jsii_.Get(
		j,
		"switchLocked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SwitchLockedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"switchLockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) TargetTests() ZeroTrustDeviceCustomProfileTargetTestsList {
	var returns ZeroTrustDeviceCustomProfileTargetTestsList
	_jsii_.Get(
		j,
		"targetTests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) TunnelProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) TunnelProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tunnelProtocolInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_custom_profile cloudflare_zero_trust_device_custom_profile} Resource.
func NewZeroTrustDeviceCustomProfile(scope constructs.Construct, id *string, config *ZeroTrustDeviceCustomProfileConfig) ZeroTrustDeviceCustomProfile {
	_init_.Initialize()

	if err := validateNewZeroTrustDeviceCustomProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDeviceCustomProfile{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDeviceCustomProfile.ZeroTrustDeviceCustomProfile",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_custom_profile cloudflare_zero_trust_device_custom_profile} Resource.
func NewZeroTrustDeviceCustomProfile_Override(z ZeroTrustDeviceCustomProfile, scope constructs.Construct, id *string, config *ZeroTrustDeviceCustomProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDeviceCustomProfile.ZeroTrustDeviceCustomProfile",
		[]any{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetAllowedToLeave(val any) {
	if err := j.validateSetAllowedToLeaveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedToLeave",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetAllowModeSwitch(val any) {
	if err := j.validateSetAllowModeSwitchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowModeSwitch",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetAllowUpdates(val any) {
	if err := j.validateSetAllowUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUpdates",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetAutoConnect(val *float64) {
	if err := j.validateSetAutoConnectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoConnect",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetCaptivePortal(val *float64) {
	if err := j.validateSetCaptivePortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captivePortal",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetDisableAutoFallback(val any) {
	if err := j.validateSetDisableAutoFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableAutoFallback",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetExcludeOfficeIps(val any) {
	if err := j.validateSetExcludeOfficeIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeOfficeIps",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetLanAllowMinutes(val *float64) {
	if err := j.validateSetLanAllowMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lanAllowMinutes",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetLanAllowSubnetSize(val *float64) {
	if err := j.validateSetLanAllowSubnetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lanAllowSubnetSize",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetMatch(val *string) {
	if err := j.validateSetMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"match",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetPrecedence(val *float64) {
	if err := j.validateSetPrecedenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"precedence",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetRegisterInterfaceIpWithDns(val any) {
	if err := j.validateSetRegisterInterfaceIpWithDnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registerInterfaceIpWithDns",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetSccmVpnBoundarySupport(val any) {
	if err := j.validateSetSccmVpnBoundarySupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sccmVpnBoundarySupport",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetSupportUrl(val *string) {
	if err := j.validateSetSupportUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportUrl",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetSwitchLocked(val any) {
	if err := j.validateSetSwitchLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"switchLocked",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDeviceCustomProfile) SetTunnelProtocol(val *string) {
	if err := j.validateSetTunnelProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tunnelProtocol",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustDeviceCustomProfile resource upon running "cdktf plan <stack-name>".
func ZeroTrustDeviceCustomProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustDeviceCustomProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceCustomProfile.ZeroTrustDeviceCustomProfile",
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
func ZeroTrustDeviceCustomProfile_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceCustomProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceCustomProfile.ZeroTrustDeviceCustomProfile",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDeviceCustomProfile_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceCustomProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceCustomProfile.ZeroTrustDeviceCustomProfile",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDeviceCustomProfile_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDeviceCustomProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDeviceCustomProfile.ZeroTrustDeviceCustomProfile",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ZeroTrustDeviceCustomProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustDeviceCustomProfile.ZeroTrustDeviceCustomProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) AddOverride(path *string, value any) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]any{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]any{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) MoveTo(moveTarget *string, index any) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) PutExclude(value any) {
	if err := z.validatePutExcludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putExclude",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) PutInclude(value any) {
	if err := z.validatePutIncludeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putInclude",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) PutServiceModeV2(value *ZeroTrustDeviceCustomProfileServiceModeV2) {
	if err := z.validatePutServiceModeV2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putServiceModeV2",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetAllowedToLeave() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedToLeave",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetAllowModeSwitch() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowModeSwitch",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetAllowUpdates() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowUpdates",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetAutoConnect() {
	_jsii_.InvokeVoid(
		z,
		"resetAutoConnect",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetCaptivePortal() {
	_jsii_.InvokeVoid(
		z,
		"resetCaptivePortal",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		z,
		"resetDescription",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetDisableAutoFallback() {
	_jsii_.InvokeVoid(
		z,
		"resetDisableAutoFallback",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetExclude() {
	_jsii_.InvokeVoid(
		z,
		"resetExclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetExcludeOfficeIps() {
	_jsii_.InvokeVoid(
		z,
		"resetExcludeOfficeIps",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetInclude() {
	_jsii_.InvokeVoid(
		z,
		"resetInclude",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetLanAllowMinutes() {
	_jsii_.InvokeVoid(
		z,
		"resetLanAllowMinutes",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetLanAllowSubnetSize() {
	_jsii_.InvokeVoid(
		z,
		"resetLanAllowSubnetSize",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetRegisterInterfaceIpWithDns() {
	_jsii_.InvokeVoid(
		z,
		"resetRegisterInterfaceIpWithDns",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetSccmVpnBoundarySupport() {
	_jsii_.InvokeVoid(
		z,
		"resetSccmVpnBoundarySupport",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetServiceModeV2() {
	_jsii_.InvokeVoid(
		z,
		"resetServiceModeV2",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetSupportUrl() {
	_jsii_.InvokeVoid(
		z,
		"resetSupportUrl",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetSwitchLocked() {
	_jsii_.InvokeVoid(
		z,
		"resetSwitchLocked",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ResetTunnelProtocol() {
	_jsii_.InvokeVoid(
		z,
		"resetTunnelProtocol",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDeviceCustomProfile) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
