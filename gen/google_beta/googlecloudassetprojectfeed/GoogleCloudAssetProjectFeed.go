package googlecloudassetprojectfeed

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecloudassetprojectfeed/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_asset_project_feed google_cloud_asset_project_feed}.
type GoogleCloudAssetProjectFeed interface {
	cdktf.TerraformResource
	AssetNames() *[]*string
	SetAssetNames(val *[]*string)
	AssetNamesInput() *[]*string
	AssetTypes() *[]*string
	SetAssetTypes(val *[]*string)
	AssetTypesInput() *[]*string
	BillingProject() *string
	SetBillingProject(val *string)
	BillingProjectInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Condition() GoogleCloudAssetProjectFeedConditionOutputReference
	ConditionInput() *GoogleCloudAssetProjectFeedCondition
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	ContentType() *string
	SetContentType(val *string)
	ContentTypeInput() *string
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FeedId() *string
	SetFeedId(val *string)
	FeedIdInput() *string
	FeedOutputConfig() GoogleCloudAssetProjectFeedFeedOutputConfigOutputReference
	FeedOutputConfigInput() *GoogleCloudAssetProjectFeedFeedOutputConfig
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudAssetProjectFeedTimeoutsOutputReference
	TimeoutsInput() any
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
	PutCondition(value *GoogleCloudAssetProjectFeedCondition)
	PutFeedOutputConfig(value *GoogleCloudAssetProjectFeedFeedOutputConfig)
	PutTimeouts(value *GoogleCloudAssetProjectFeedTimeouts)
	ResetAssetNames()
	ResetAssetTypes()
	ResetBillingProject()
	ResetCondition()
	ResetContentType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleCloudAssetProjectFeed
type jsiiProxy_GoogleCloudAssetProjectFeed struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) AssetNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) AssetNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) AssetTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) AssetTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assetTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) BillingProject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) BillingProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingProjectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Condition() GoogleCloudAssetProjectFeedConditionOutputReference {
	var returns GoogleCloudAssetProjectFeedConditionOutputReference
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) ConditionInput() *GoogleCloudAssetProjectFeedCondition {
	var returns *GoogleCloudAssetProjectFeedCondition
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) FeedId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) FeedIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"feedIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) FeedOutputConfig() GoogleCloudAssetProjectFeedFeedOutputConfigOutputReference {
	var returns GoogleCloudAssetProjectFeedFeedOutputConfigOutputReference
	_jsii_.Get(
		j,
		"feedOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) FeedOutputConfigInput() *GoogleCloudAssetProjectFeedFeedOutputConfig {
	var returns *GoogleCloudAssetProjectFeedFeedOutputConfig
	_jsii_.Get(
		j,
		"feedOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) Timeouts() GoogleCloudAssetProjectFeedTimeoutsOutputReference {
	var returns GoogleCloudAssetProjectFeedTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_asset_project_feed google_cloud_asset_project_feed} Resource.
func NewGoogleCloudAssetProjectFeed(scope constructs.Construct, id *string, config *GoogleCloudAssetProjectFeedConfig) GoogleCloudAssetProjectFeed {
	_init_.Initialize()

	if err := validateNewGoogleCloudAssetProjectFeedParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudAssetProjectFeed{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudAssetProjectFeed.GoogleCloudAssetProjectFeed",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_asset_project_feed google_cloud_asset_project_feed} Resource.
func NewGoogleCloudAssetProjectFeed_Override(g GoogleCloudAssetProjectFeed, scope constructs.Construct, id *string, config *GoogleCloudAssetProjectFeedConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudAssetProjectFeed.GoogleCloudAssetProjectFeed",
		[]any{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetAssetNames(val *[]*string) {
	if err := j.validateSetAssetNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetNames",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetAssetTypes(val *[]*string) {
	if err := j.validateSetAssetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetTypes",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetBillingProject(val *string) {
	if err := j.validateSetBillingProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingProject",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetFeedId(val *string) {
	if err := j.validateSetFeedIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feedId",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudAssetProjectFeed) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleCloudAssetProjectFeed resource upon running "cdktf plan <stack-name>".
func GoogleCloudAssetProjectFeed_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCloudAssetProjectFeed_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudAssetProjectFeed.GoogleCloudAssetProjectFeed",
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
func GoogleCloudAssetProjectFeed_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudAssetProjectFeed_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudAssetProjectFeed.GoogleCloudAssetProjectFeed",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudAssetProjectFeed_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudAssetProjectFeed_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudAssetProjectFeed.GoogleCloudAssetProjectFeed",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudAssetProjectFeed_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudAssetProjectFeed_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudAssetProjectFeed.GoogleCloudAssetProjectFeed",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func GoogleCloudAssetProjectFeed_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleCloudAssetProjectFeed.GoogleCloudAssetProjectFeed",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) AddOverride(path *string, value any) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]any{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]any{id, provider},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) MoveTo(moveTarget *string, index any) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) PutCondition(value *GoogleCloudAssetProjectFeedCondition) {
	if err := g.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCondition",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) PutFeedOutputConfig(value *GoogleCloudAssetProjectFeedFeedOutputConfig) {
	if err := g.validatePutFeedOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFeedOutputConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) PutTimeouts(value *GoogleCloudAssetProjectFeedTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetAssetNames() {
	_jsii_.InvokeVoid(
		g,
		"resetAssetNames",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetAssetTypes() {
	_jsii_.InvokeVoid(
		g,
		"resetAssetTypes",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetBillingProject() {
	_jsii_.InvokeVoid(
		g,
		"resetBillingProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetCondition() {
	_jsii_.InvokeVoid(
		g,
		"resetCondition",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetContentType() {
	_jsii_.InvokeVoid(
		g,
		"resetContentType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudAssetProjectFeed) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
