package googlevmwareenginenetworkpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlevmwareenginenetworkpolicy/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_vmwareengine_network_policy google_vmwareengine_network_policy}.
type GoogleVmwareengineNetworkPolicy interface {
	cdktf.TerraformResource
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
	CreateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeServicesCidr() *string
	SetEdgeServicesCidr(val *string)
	EdgeServicesCidrInput() *string
	ExternalIp() GoogleVmwareengineNetworkPolicyExternalIpOutputReference
	ExternalIpInput() *GoogleVmwareengineNetworkPolicyExternalIp
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
	InternetAccess() GoogleVmwareengineNetworkPolicyInternetAccessOutputReference
	InternetAccessInput() *GoogleVmwareengineNetworkPolicyInternetAccess
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleVmwareengineNetworkPolicyTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
	VmwareEngineNetwork() *string
	SetVmwareEngineNetwork(val *string)
	VmwareEngineNetworkCanonical() *string
	VmwareEngineNetworkInput() *string
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
	PutExternalIp(value *GoogleVmwareengineNetworkPolicyExternalIp)
	PutInternetAccess(value *GoogleVmwareengineNetworkPolicyInternetAccess)
	PutTimeouts(value *GoogleVmwareengineNetworkPolicyTimeouts)
	ResetDescription()
	ResetExternalIp()
	ResetId()
	ResetInternetAccess()
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

// The jsii proxy struct for GoogleVmwareengineNetworkPolicy
type jsiiProxy_GoogleVmwareengineNetworkPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) EdgeServicesCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeServicesCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) EdgeServicesCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeServicesCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) ExternalIp() GoogleVmwareengineNetworkPolicyExternalIpOutputReference {
	var returns GoogleVmwareengineNetworkPolicyExternalIpOutputReference
	_jsii_.Get(
		j,
		"externalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) ExternalIpInput() *GoogleVmwareengineNetworkPolicyExternalIp {
	var returns *GoogleVmwareengineNetworkPolicyExternalIp
	_jsii_.Get(
		j,
		"externalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) InternetAccess() GoogleVmwareengineNetworkPolicyInternetAccessOutputReference {
	var returns GoogleVmwareengineNetworkPolicyInternetAccessOutputReference
	_jsii_.Get(
		j,
		"internetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) InternetAccessInput() *GoogleVmwareengineNetworkPolicyInternetAccess {
	var returns *GoogleVmwareengineNetworkPolicyInternetAccess
	_jsii_.Get(
		j,
		"internetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Timeouts() GoogleVmwareengineNetworkPolicyTimeoutsOutputReference {
	var returns GoogleVmwareengineNetworkPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) VmwareEngineNetwork() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareEngineNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) VmwareEngineNetworkCanonical() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareEngineNetworkCanonical",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy) VmwareEngineNetworkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmwareEngineNetworkInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_vmwareengine_network_policy google_vmwareengine_network_policy} Resource.
func NewGoogleVmwareengineNetworkPolicy(scope constructs.Construct, id *string, config *GoogleVmwareengineNetworkPolicyConfig) GoogleVmwareengineNetworkPolicy {
	_init_.Initialize()

	if err := validateNewGoogleVmwareengineNetworkPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleVmwareengineNetworkPolicy{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleVmwareengineNetworkPolicy.GoogleVmwareengineNetworkPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_vmwareengine_network_policy google_vmwareengine_network_policy} Resource.
func NewGoogleVmwareengineNetworkPolicy_Override(g GoogleVmwareengineNetworkPolicy, scope constructs.Construct, id *string, config *GoogleVmwareengineNetworkPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleVmwareengineNetworkPolicy.GoogleVmwareengineNetworkPolicy",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetEdgeServicesCidr(val *string) {
	if err := j.validateSetEdgeServicesCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeServicesCidr",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleVmwareengineNetworkPolicy)SetVmwareEngineNetwork(val *string) {
	if err := j.validateSetVmwareEngineNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmwareEngineNetwork",
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
func GoogleVmwareengineNetworkPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVmwareengineNetworkPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleVmwareengineNetworkPolicy.GoogleVmwareengineNetworkPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVmwareengineNetworkPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVmwareengineNetworkPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleVmwareengineNetworkPolicy.GoogleVmwareengineNetworkPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleVmwareengineNetworkPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleVmwareengineNetworkPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleVmwareengineNetworkPolicy.GoogleVmwareengineNetworkPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleVmwareengineNetworkPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleVmwareengineNetworkPolicy.GoogleVmwareengineNetworkPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) PutExternalIp(value *GoogleVmwareengineNetworkPolicyExternalIp) {
	if err := g.validatePutExternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExternalIp",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) PutInternetAccess(value *GoogleVmwareengineNetworkPolicyInternetAccess) {
	if err := g.validatePutInternetAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInternetAccess",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) PutTimeouts(value *GoogleVmwareengineNetworkPolicyTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ResetExternalIp() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalIp",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ResetInternetAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetInternetAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleVmwareengineNetworkPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

