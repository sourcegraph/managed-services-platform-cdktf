package computebackendbucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/computebackendbucket/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket google_compute_backend_bucket}.
type ComputeBackendBucket interface {
	cdktf.TerraformResource
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnPolicy() ComputeBackendBucketCdnPolicyOutputReference
	CdnPolicyInput() *ComputeBackendBucketCdnPolicy
	CompressionMode() *string
	SetCompressionMode(val *string)
	CompressionModeInput() *string
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
	CreationTimestamp() *string
	CustomResponseHeaders() *[]*string
	SetCustomResponseHeaders(val *[]*string)
	CustomResponseHeadersInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeSecurityPolicy() *string
	SetEdgeSecurityPolicy(val *string)
	EdgeSecurityPolicyInput() *string
	EnableCdn() interface{}
	SetEnableCdn(val interface{})
	EnableCdnInput() interface{}
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
	SetName(val *string)
	NameInput() *string
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
	SelfLink() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeBackendBucketTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCdnPolicy(value *ComputeBackendBucketCdnPolicy)
	PutTimeouts(value *ComputeBackendBucketTimeouts)
	ResetCdnPolicy()
	ResetCompressionMode()
	ResetCustomResponseHeaders()
	ResetDescription()
	ResetEdgeSecurityPolicy()
	ResetEnableCdn()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeBackendBucket
type jsiiProxy_ComputeBackendBucket struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeBackendBucket) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CdnPolicy() ComputeBackendBucketCdnPolicyOutputReference {
	var returns ComputeBackendBucketCdnPolicyOutputReference
	_jsii_.Get(
		j,
		"cdnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CdnPolicyInput() *ComputeBackendBucketCdnPolicy {
	var returns *ComputeBackendBucketCdnPolicy
	_jsii_.Get(
		j,
		"cdnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CompressionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CompressionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CustomResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) CustomResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) EdgeSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) EdgeSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) EnableCdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) EnableCdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) Timeouts() ComputeBackendBucketTimeoutsOutputReference {
	var returns ComputeBackendBucketTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeBackendBucket) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket google_compute_backend_bucket} Resource.
func NewComputeBackendBucket(scope constructs.Construct, id *string, config *ComputeBackendBucketConfig) ComputeBackendBucket {
	_init_.Initialize()

	if err := validateNewComputeBackendBucketParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeBackendBucket{}

	_jsii_.Create(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucket",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/4.69.1/docs/resources/compute_backend_bucket google_compute_backend_bucket} Resource.
func NewComputeBackendBucket_Override(c ComputeBackendBucket, scope constructs.Construct, id *string, config *ComputeBackendBucketConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucket",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetCompressionMode(val *string) {
	if err := j.validateSetCompressionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionMode",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetCustomResponseHeaders(val *[]*string) {
	if err := j.validateSetCustomResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetEdgeSecurityPolicy(val *string) {
	if err := j.validateSetEdgeSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetEnableCdn(val interface{}) {
	if err := j.validateSetEnableCdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCdn",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeBackendBucket)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
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
func ComputeBackendBucket_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeBackendBucket_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucket",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeBackendBucket_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeBackendBucket_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucket",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeBackendBucket_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeBackendBucket_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucket",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeBackendBucket_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeBackendBucket.ComputeBackendBucket",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeBackendBucket) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeBackendBucket) PutCdnPolicy(value *ComputeBackendBucketCdnPolicy) {
	if err := c.validatePutCdnPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCdnPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeBackendBucket) PutTimeouts(value *ComputeBackendBucketTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetCdnPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetCdnPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetCompressionMode() {
	_jsii_.InvokeVoid(
		c,
		"resetCompressionMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetCustomResponseHeaders() {
	_jsii_.InvokeVoid(
		c,
		"resetCustomResponseHeaders",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetEdgeSecurityPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetEdgeSecurityPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetEnableCdn() {
	_jsii_.InvokeVoid(
		c,
		"resetEnableCdn",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeBackendBucket) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeBackendBucket) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
