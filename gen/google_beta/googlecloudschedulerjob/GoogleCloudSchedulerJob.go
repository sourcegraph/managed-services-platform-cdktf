package googlecloudschedulerjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecloudschedulerjob/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job google_cloud_scheduler_job}.
type GoogleCloudSchedulerJob interface {
	cdktf.TerraformResource
	AppEngineHttpTarget() GoogleCloudSchedulerJobAppEngineHttpTargetOutputReference
	AppEngineHttpTargetInput() *GoogleCloudSchedulerJobAppEngineHttpTarget
	AttemptDeadline() *string
	SetAttemptDeadline(val *string)
	AttemptDeadlineInput() *string
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
	HttpTarget() GoogleCloudSchedulerJobHttpTargetOutputReference
	HttpTargetInput() *GoogleCloudSchedulerJobHttpTarget
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Paused() any
	SetPaused(val any)
	PausedInput() any
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
	PubsubTarget() GoogleCloudSchedulerJobPubsubTargetOutputReference
	PubsubTargetInput() *GoogleCloudSchedulerJobPubsubTarget
	// Experimental.
	RawOverrides() any
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RetryConfig() GoogleCloudSchedulerJobRetryConfigOutputReference
	RetryConfigInput() *GoogleCloudSchedulerJobRetryConfig
	Schedule() *string
	SetSchedule(val *string)
	ScheduleInput() *string
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleCloudSchedulerJobTimeoutsOutputReference
	TimeoutsInput() any
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutAppEngineHttpTarget(value *GoogleCloudSchedulerJobAppEngineHttpTarget)
	PutHttpTarget(value *GoogleCloudSchedulerJobHttpTarget)
	PutPubsubTarget(value *GoogleCloudSchedulerJobPubsubTarget)
	PutRetryConfig(value *GoogleCloudSchedulerJobRetryConfig)
	PutTimeouts(value *GoogleCloudSchedulerJobTimeouts)
	ResetAppEngineHttpTarget()
	ResetAttemptDeadline()
	ResetDescription()
	ResetHttpTarget()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPaused()
	ResetProject()
	ResetPubsubTarget()
	ResetRegion()
	ResetRetryConfig()
	ResetSchedule()
	ResetTimeouts()
	ResetTimeZone()
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

// The jsii proxy struct for GoogleCloudSchedulerJob
type jsiiProxy_GoogleCloudSchedulerJob struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) AppEngineHttpTarget() GoogleCloudSchedulerJobAppEngineHttpTargetOutputReference {
	var returns GoogleCloudSchedulerJobAppEngineHttpTargetOutputReference
	_jsii_.Get(
		j,
		"appEngineHttpTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) AppEngineHttpTargetInput() *GoogleCloudSchedulerJobAppEngineHttpTarget {
	var returns *GoogleCloudSchedulerJobAppEngineHttpTarget
	_jsii_.Get(
		j,
		"appEngineHttpTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) AttemptDeadline() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attemptDeadline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) AttemptDeadlineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attemptDeadlineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) HttpTarget() GoogleCloudSchedulerJobHttpTargetOutputReference {
	var returns GoogleCloudSchedulerJobHttpTargetOutputReference
	_jsii_.Get(
		j,
		"httpTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) HttpTargetInput() *GoogleCloudSchedulerJobHttpTarget {
	var returns *GoogleCloudSchedulerJobHttpTarget
	_jsii_.Get(
		j,
		"httpTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Paused() any {
	var returns any
	_jsii_.Get(
		j,
		"paused",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) PausedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pausedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) PubsubTarget() GoogleCloudSchedulerJobPubsubTargetOutputReference {
	var returns GoogleCloudSchedulerJobPubsubTargetOutputReference
	_jsii_.Get(
		j,
		"pubsubTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) PubsubTargetInput() *GoogleCloudSchedulerJobPubsubTarget {
	var returns *GoogleCloudSchedulerJobPubsubTarget
	_jsii_.Get(
		j,
		"pubsubTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) RetryConfig() GoogleCloudSchedulerJobRetryConfigOutputReference {
	var returns GoogleCloudSchedulerJobRetryConfigOutputReference
	_jsii_.Get(
		j,
		"retryConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) RetryConfigInput() *GoogleCloudSchedulerJobRetryConfig {
	var returns *GoogleCloudSchedulerJobRetryConfig
	_jsii_.Get(
		j,
		"retryConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Schedule() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) ScheduleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) Timeouts() GoogleCloudSchedulerJobTimeoutsOutputReference {
	var returns GoogleCloudSchedulerJobTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job google_cloud_scheduler_job} Resource.
func NewGoogleCloudSchedulerJob(scope constructs.Construct, id *string, config *GoogleCloudSchedulerJobConfig) GoogleCloudSchedulerJob {
	_init_.Initialize()

	if err := validateNewGoogleCloudSchedulerJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudSchedulerJob{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudSchedulerJob.GoogleCloudSchedulerJob",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_scheduler_job google_cloud_scheduler_job} Resource.
func NewGoogleCloudSchedulerJob_Override(g GoogleCloudSchedulerJob, scope constructs.Construct, id *string, config *GoogleCloudSchedulerJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudSchedulerJob.GoogleCloudSchedulerJob",
		[]any{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetAttemptDeadline(val *string) {
	if err := j.validateSetAttemptDeadlineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attemptDeadline",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetPaused(val any) {
	if err := j.validateSetPausedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"paused",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetSchedule(val *string) {
	if err := j.validateSetScheduleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"schedule",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudSchedulerJob) SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

// Generates CDKTF code for importing a GoogleCloudSchedulerJob resource upon running "cdktf plan <stack-name>".
func GoogleCloudSchedulerJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleCloudSchedulerJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudSchedulerJob.GoogleCloudSchedulerJob",
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
func GoogleCloudSchedulerJob_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudSchedulerJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudSchedulerJob.GoogleCloudSchedulerJob",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudSchedulerJob_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudSchedulerJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudSchedulerJob.GoogleCloudSchedulerJob",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleCloudSchedulerJob_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleCloudSchedulerJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleCloudSchedulerJob.GoogleCloudSchedulerJob",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func GoogleCloudSchedulerJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleCloudSchedulerJob.GoogleCloudSchedulerJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) AddOverride(path *string, value any) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]any{path, value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]any{id, provider},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleCloudSchedulerJob) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) MoveTo(moveTarget *string, index any) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) PutAppEngineHttpTarget(value *GoogleCloudSchedulerJobAppEngineHttpTarget) {
	if err := g.validatePutAppEngineHttpTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAppEngineHttpTarget",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) PutHttpTarget(value *GoogleCloudSchedulerJobHttpTarget) {
	if err := g.validatePutHttpTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHttpTarget",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) PutPubsubTarget(value *GoogleCloudSchedulerJobPubsubTarget) {
	if err := g.validatePutPubsubTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubTarget",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) PutRetryConfig(value *GoogleCloudSchedulerJobRetryConfig) {
	if err := g.validatePutRetryConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRetryConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) PutTimeouts(value *GoogleCloudSchedulerJobTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetAppEngineHttpTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetAppEngineHttpTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetAttemptDeadline() {
	_jsii_.InvokeVoid(
		g,
		"resetAttemptDeadline",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetHttpTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetHttpTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetPaused() {
	_jsii_.InvokeVoid(
		g,
		"resetPaused",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetPubsubTarget() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubTarget",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetRegion() {
	_jsii_.InvokeVoid(
		g,
		"resetRegion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetRetryConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetRetryConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetSchedule() {
	_jsii_.InvokeVoid(
		g,
		"resetSchedule",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ResetTimeZone() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudSchedulerJob) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
