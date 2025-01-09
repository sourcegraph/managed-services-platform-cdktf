package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/provider/internal"
)

// Represents a {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs opsgenie}.
type OpsgenieProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	ApiKey() *string
	SetApiKey(val *string)
	ApiKeyInput() *string
	ApiRetryCount() *float64
	SetApiRetryCount(val *float64)
	ApiRetryCountInput() *float64
	ApiRetryWaitMax() *float64
	SetApiRetryWaitMax(val *float64)
	ApiRetryWaitMaxInput() *float64
	ApiRetryWaitMin() *float64
	SetApiRetryWaitMin(val *float64)
	ApiRetryWaitMinInput() *float64
	ApiUrl() *string
	SetApiUrl(val *string)
	ApiUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetApiRetryCount()
	ResetApiRetryWaitMax()
	ResetApiRetryWaitMin()
	ResetApiUrl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for OpsgenieProvider
type jsiiProxy_OpsgenieProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_OpsgenieProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiRetryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiRetryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiRetryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiRetryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiRetryWaitMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiRetryWaitMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiRetryWaitMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiRetryWaitMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiRetryWaitMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiRetryWaitMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiRetryWaitMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"apiRetryWaitMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ApiUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsgenieProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs opsgenie} Resource.
func NewOpsgenieProvider(scope constructs.Construct, id *string, config *OpsgenieProviderConfig) OpsgenieProvider {
	_init_.Initialize()

	if err := validateNewOpsgenieProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsgenieProvider{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.provider.OpsgenieProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs opsgenie} Resource.
func NewOpsgenieProvider_Override(o OpsgenieProvider, scope constructs.Construct, id *string, config *OpsgenieProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.provider.OpsgenieProvider",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsgenieProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_OpsgenieProvider)SetApiKey(val *string) {
	_jsii_.Set(
		j,
		"apiKey",
		val,
	)
}

func (j *jsiiProxy_OpsgenieProvider)SetApiRetryCount(val *float64) {
	_jsii_.Set(
		j,
		"apiRetryCount",
		val,
	)
}

func (j *jsiiProxy_OpsgenieProvider)SetApiRetryWaitMax(val *float64) {
	_jsii_.Set(
		j,
		"apiRetryWaitMax",
		val,
	)
}

func (j *jsiiProxy_OpsgenieProvider)SetApiRetryWaitMin(val *float64) {
	_jsii_.Set(
		j,
		"apiRetryWaitMin",
		val,
	)
}

func (j *jsiiProxy_OpsgenieProvider)SetApiUrl(val *string) {
	_jsii_.Set(
		j,
		"apiUrl",
		val,
	)
}

// Generates CDKTF code for importing a OpsgenieProvider resource upon running "cdktf plan <stack-name>".
func OpsgenieProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOpsgenieProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.provider.OpsgenieProvider",
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
func OpsgenieProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsgenieProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.provider.OpsgenieProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsgenieProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsgenieProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.provider.OpsgenieProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsgenieProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsgenieProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-opsgenie.provider.OpsgenieProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsgenieProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-opsgenie.provider.OpsgenieProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpsgenieProvider) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpsgenieProvider) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsgenieProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		o,
		"resetAlias",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsgenieProvider) ResetApiRetryCount() {
	_jsii_.InvokeVoid(
		o,
		"resetApiRetryCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsgenieProvider) ResetApiRetryWaitMax() {
	_jsii_.InvokeVoid(
		o,
		"resetApiRetryWaitMax",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsgenieProvider) ResetApiRetryWaitMin() {
	_jsii_.InvokeVoid(
		o,
		"resetApiRetryWaitMin",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsgenieProvider) ResetApiUrl() {
	_jsii_.InvokeVoid(
		o,
		"resetApiUrl",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsgenieProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsgenieProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsgenieProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsgenieProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsgenieProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsgenieProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsgenieProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

