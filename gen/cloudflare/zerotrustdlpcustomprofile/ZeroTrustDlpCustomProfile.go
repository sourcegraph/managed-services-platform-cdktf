package zerotrustdlpcustomprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/zerotrustdlpcustomprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_custom_profile cloudflare_zero_trust_dlp_custom_profile}.
type ZeroTrustDlpCustomProfile interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AiContextEnabled() any
	SetAiContextEnabled(val any)
	AiContextEnabledInput() any
	AllowedMatchCount() *float64
	SetAllowedMatchCount(val *float64)
	AllowedMatchCountInput() *float64
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfidenceThreshold() *string
	SetConfidenceThreshold(val *string)
	ConfidenceThresholdInput() *string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	ContextAwareness() ZeroTrustDlpCustomProfileContextAwarenessOutputReference
	ContextAwarenessInput() any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	CreatedAt() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Entries() ZeroTrustDlpCustomProfileEntriesList
	EntriesInput() any
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OcrEnabled() any
	SetOcrEnabled(val any)
	OcrEnabledInput() any
	OpenAccess() cdktf.IResolvable
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
	SharedEntries() ZeroTrustDlpCustomProfileSharedEntriesList
	SharedEntriesInput() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	UpdatedAt() *string
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
	PutContextAwareness(value *ZeroTrustDlpCustomProfileContextAwareness)
	PutEntries(value any)
	PutSharedEntries(value any)
	ResetAiContextEnabled()
	ResetAllowedMatchCount()
	ResetConfidenceThreshold()
	ResetContextAwareness()
	ResetDescription()
	ResetOcrEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSharedEntries()
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

// The jsii proxy struct for ZeroTrustDlpCustomProfile
type jsiiProxy_ZeroTrustDlpCustomProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AiContextEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"aiContextEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AiContextEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"aiContextEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AllowedMatchCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedMatchCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) AllowedMatchCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allowedMatchCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ConfidenceThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ConfidenceThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"confidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ContextAwareness() ZeroTrustDlpCustomProfileContextAwarenessOutputReference {
	var returns ZeroTrustDlpCustomProfileContextAwarenessOutputReference
	_jsii_.Get(
		j,
		"contextAwareness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ContextAwarenessInput() any {
	var returns any
	_jsii_.Get(
		j,
		"contextAwarenessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Entries() ZeroTrustDlpCustomProfileEntriesList {
	var returns ZeroTrustDlpCustomProfileEntriesList
	_jsii_.Get(
		j,
		"entries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) EntriesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"entriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) OcrEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"ocrEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) OcrEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ocrEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) OpenAccess() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"openAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SharedEntries() ZeroTrustDlpCustomProfileSharedEntriesList {
	var returns ZeroTrustDlpCustomProfileSharedEntriesList
	_jsii_.Get(
		j,
		"sharedEntries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SharedEntriesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sharedEntriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) UpdatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_custom_profile cloudflare_zero_trust_dlp_custom_profile} Resource.
func NewZeroTrustDlpCustomProfile(scope constructs.Construct, id *string, config *ZeroTrustDlpCustomProfileConfig) ZeroTrustDlpCustomProfile {
	_init_.Initialize()

	if err := validateNewZeroTrustDlpCustomProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ZeroTrustDlpCustomProfile{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_dlp_custom_profile cloudflare_zero_trust_dlp_custom_profile} Resource.
func NewZeroTrustDlpCustomProfile_Override(z ZeroTrustDlpCustomProfile, scope constructs.Construct, id *string, config *ZeroTrustDlpCustomProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		[]any{scope, id, config},
		z,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetAiContextEnabled(val any) {
	if err := j.validateSetAiContextEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aiContextEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetAllowedMatchCount(val *float64) {
	if err := j.validateSetAllowedMatchCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMatchCount",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetConfidenceThreshold(val *string) {
	if err := j.validateSetConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetOcrEnabled(val any) {
	if err := j.validateSetOcrEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocrEnabled",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ZeroTrustDlpCustomProfile) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ZeroTrustDlpCustomProfile resource upon running "cdktf plan <stack-name>".
func ZeroTrustDlpCustomProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
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
func ZeroTrustDlpCustomProfile_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDlpCustomProfile_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func ZeroTrustDlpCustomProfile_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateZeroTrustDlpCustomProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func ZeroTrustDlpCustomProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.zeroTrustDlpCustomProfile.ZeroTrustDlpCustomProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) AddMoveTarget(moveTarget *string) {
	if err := z.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) AddOverride(path *string, value any) {
	if err := z.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"addOverride",
		[]any{path, value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		z,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := z.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"importFrom",
		[]any{id, provider},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) MoveFromId(id *string) {
	if err := z.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveFromId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) MoveTo(moveTarget *string, index any) {
	if err := z.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) MoveToId(id *string) {
	if err := z.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"moveToId",
		[]any{id},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) OverrideLogicalId(newLogicalId *string) {
	if err := z.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) PutContextAwareness(value *ZeroTrustDlpCustomProfileContextAwareness) {
	if err := z.validatePutContextAwarenessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putContextAwareness",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) PutEntries(value any) {
	if err := z.validatePutEntriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putEntries",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) PutSharedEntries(value any) {
	if err := z.validatePutSharedEntriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		z,
		"putSharedEntries",
		[]any{value},
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetAiContextEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetAiContextEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetAllowedMatchCount() {
	_jsii_.InvokeVoid(
		z,
		"resetAllowedMatchCount",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetConfidenceThreshold() {
	_jsii_.InvokeVoid(
		z,
		"resetConfidenceThreshold",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetContextAwareness() {
	_jsii_.InvokeVoid(
		z,
		"resetContextAwareness",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		z,
		"resetDescription",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetOcrEnabled() {
	_jsii_.InvokeVoid(
		z,
		"resetOcrEnabled",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		z,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ResetSharedEntries() {
	_jsii_.InvokeVoid(
		z,
		"resetSharedEntries",
		nil, // no parameters
	)
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		z,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		z,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (z *jsiiProxy_ZeroTrustDlpCustomProfile) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		z,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
