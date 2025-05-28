package googlebigtableappprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlebigtableappprofile/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_bigtable_app_profile google_bigtable_app_profile}.
type GoogleBigtableAppProfile interface {
	cdktf.TerraformResource
	AppProfileId() *string
	SetAppProfileId(val *string)
	AppProfileIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataBoostIsolationReadOnly() GoogleBigtableAppProfileDataBoostIsolationReadOnlyOutputReference
	DataBoostIsolationReadOnlyInput() *GoogleBigtableAppProfileDataBoostIsolationReadOnly
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	IgnoreWarnings() interface{}
	SetIgnoreWarnings(val interface{})
	IgnoreWarningsInput() interface{}
	Instance() *string
	SetInstance(val *string)
	InstanceInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MultiClusterRoutingClusterIds() *[]*string
	SetMultiClusterRoutingClusterIds(val *[]*string)
	MultiClusterRoutingClusterIdsInput() *[]*string
	MultiClusterRoutingUseAny() interface{}
	SetMultiClusterRoutingUseAny(val interface{})
	MultiClusterRoutingUseAnyInput() interface{}
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
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RowAffinity() interface{}
	SetRowAffinity(val interface{})
	RowAffinityInput() interface{}
	SingleClusterRouting() GoogleBigtableAppProfileSingleClusterRoutingOutputReference
	SingleClusterRoutingInput() *GoogleBigtableAppProfileSingleClusterRouting
	StandardIsolation() GoogleBigtableAppProfileStandardIsolationOutputReference
	StandardIsolationInput() *GoogleBigtableAppProfileStandardIsolation
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleBigtableAppProfileTimeoutsOutputReference
	TimeoutsInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	HasResourceMove() interface{}
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
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDataBoostIsolationReadOnly(value *GoogleBigtableAppProfileDataBoostIsolationReadOnly)
	PutSingleClusterRouting(value *GoogleBigtableAppProfileSingleClusterRouting)
	PutStandardIsolation(value *GoogleBigtableAppProfileStandardIsolation)
	PutTimeouts(value *GoogleBigtableAppProfileTimeouts)
	ResetDataBoostIsolationReadOnly()
	ResetDescription()
	ResetId()
	ResetIgnoreWarnings()
	ResetInstance()
	ResetMultiClusterRoutingClusterIds()
	ResetMultiClusterRoutingUseAny()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRowAffinity()
	ResetSingleClusterRouting()
	ResetStandardIsolation()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GoogleBigtableAppProfile
type jsiiProxy_GoogleBigtableAppProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleBigtableAppProfile) AppProfileId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appProfileId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) AppProfileIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appProfileIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) DataBoostIsolationReadOnly() GoogleBigtableAppProfileDataBoostIsolationReadOnlyOutputReference {
	var returns GoogleBigtableAppProfileDataBoostIsolationReadOnlyOutputReference
	_jsii_.Get(
		j,
		"dataBoostIsolationReadOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) DataBoostIsolationReadOnlyInput() *GoogleBigtableAppProfileDataBoostIsolationReadOnly {
	var returns *GoogleBigtableAppProfileDataBoostIsolationReadOnly
	_jsii_.Get(
		j,
		"dataBoostIsolationReadOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) IgnoreWarnings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreWarnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) IgnoreWarningsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreWarningsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Instance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) InstanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) MultiClusterRoutingClusterIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"multiClusterRoutingClusterIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) MultiClusterRoutingClusterIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"multiClusterRoutingClusterIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) MultiClusterRoutingUseAny() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiClusterRoutingUseAny",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) MultiClusterRoutingUseAnyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiClusterRoutingUseAnyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) RowAffinity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) RowAffinityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rowAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) SingleClusterRouting() GoogleBigtableAppProfileSingleClusterRoutingOutputReference {
	var returns GoogleBigtableAppProfileSingleClusterRoutingOutputReference
	_jsii_.Get(
		j,
		"singleClusterRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) SingleClusterRoutingInput() *GoogleBigtableAppProfileSingleClusterRouting {
	var returns *GoogleBigtableAppProfileSingleClusterRouting
	_jsii_.Get(
		j,
		"singleClusterRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) StandardIsolation() GoogleBigtableAppProfileStandardIsolationOutputReference {
	var returns GoogleBigtableAppProfileStandardIsolationOutputReference
	_jsii_.Get(
		j,
		"standardIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) StandardIsolationInput() *GoogleBigtableAppProfileStandardIsolation {
	var returns *GoogleBigtableAppProfileStandardIsolation
	_jsii_.Get(
		j,
		"standardIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) Timeouts() GoogleBigtableAppProfileTimeoutsOutputReference {
	var returns GoogleBigtableAppProfileTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBigtableAppProfile) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_bigtable_app_profile google_bigtable_app_profile} Resource.
func NewGoogleBigtableAppProfile(scope constructs.Construct, id *string, config *GoogleBigtableAppProfileConfig) GoogleBigtableAppProfile {
	_init_.Initialize()

	if err := validateNewGoogleBigtableAppProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBigtableAppProfile{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleBigtableAppProfile.GoogleBigtableAppProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_bigtable_app_profile google_bigtable_app_profile} Resource.
func NewGoogleBigtableAppProfile_Override(g GoogleBigtableAppProfile, scope constructs.Construct, id *string, config *GoogleBigtableAppProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleBigtableAppProfile.GoogleBigtableAppProfile",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetAppProfileId(val *string) {
	if err := j.validateSetAppProfileIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appProfileId",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetIgnoreWarnings(val interface{}) {
	if err := j.validateSetIgnoreWarningsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreWarnings",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetInstance(val *string) {
	if err := j.validateSetInstanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instance",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetMultiClusterRoutingClusterIds(val *[]*string) {
	if err := j.validateSetMultiClusterRoutingClusterIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiClusterRoutingClusterIds",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetMultiClusterRoutingUseAny(val interface{}) {
	if err := j.validateSetMultiClusterRoutingUseAnyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiClusterRoutingUseAny",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleBigtableAppProfile)SetRowAffinity(val interface{}) {
	if err := j.validateSetRowAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowAffinity",
		val,
	)
}

// Generates CDKTF code for importing a GoogleBigtableAppProfile resource upon running "cdktf plan <stack-name>".
func GoogleBigtableAppProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleBigtableAppProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigtableAppProfile.GoogleBigtableAppProfile",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
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
func GoogleBigtableAppProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigtableAppProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigtableAppProfile.GoogleBigtableAppProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigtableAppProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigtableAppProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigtableAppProfile.GoogleBigtableAppProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleBigtableAppProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleBigtableAppProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleBigtableAppProfile.GoogleBigtableAppProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleBigtableAppProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleBigtableAppProfile.GoogleBigtableAppProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) PutDataBoostIsolationReadOnly(value *GoogleBigtableAppProfileDataBoostIsolationReadOnly) {
	if err := g.validatePutDataBoostIsolationReadOnlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDataBoostIsolationReadOnly",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) PutSingleClusterRouting(value *GoogleBigtableAppProfileSingleClusterRouting) {
	if err := g.validatePutSingleClusterRoutingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSingleClusterRouting",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) PutStandardIsolation(value *GoogleBigtableAppProfileStandardIsolation) {
	if err := g.validatePutStandardIsolationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStandardIsolation",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) PutTimeouts(value *GoogleBigtableAppProfileTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetDataBoostIsolationReadOnly() {
	_jsii_.InvokeVoid(
		g,
		"resetDataBoostIsolationReadOnly",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetIgnoreWarnings() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreWarnings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetInstance() {
	_jsii_.InvokeVoid(
		g,
		"resetInstance",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetMultiClusterRoutingClusterIds() {
	_jsii_.InvokeVoid(
		g,
		"resetMultiClusterRoutingClusterIds",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetMultiClusterRoutingUseAny() {
	_jsii_.InvokeVoid(
		g,
		"resetMultiClusterRoutingUseAny",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetRowAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetRowAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetSingleClusterRouting() {
	_jsii_.InvokeVoid(
		g,
		"resetSingleClusterRouting",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetStandardIsolation() {
	_jsii_.InvokeVoid(
		g,
		"resetStandardIsolation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBigtableAppProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBigtableAppProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

