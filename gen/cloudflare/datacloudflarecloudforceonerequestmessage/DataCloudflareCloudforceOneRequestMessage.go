package datacloudflarecloudforceonerequestmessage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflarecloudforceonerequestmessage/internal"
)

// Represents a {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/cloudforce_one_request_message cloudflare_cloudforce_one_request_message}.
type DataCloudflareCloudforceOneRequestMessage interface {
	cdktf.TerraformDataSource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	After() *string
	SetAfter(val *string)
	AfterInput() *string
	Author() *string
	Before() *string
	SetBefore(val *string)
	BeforeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	Created() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *float64
	IsFollowOnRequest() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Page() *float64
	SetPage(val *float64)
	PageInput() *float64
	PerPage() *float64
	SetPerPage(val *float64)
	PerPageInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	RequestId() *string
	SetRequestId(val *string)
	RequestIdInput() *string
	SortBy() *string
	SetSortBy(val *string)
	SortByInput() *string
	SortOrder() *string
	SetSortOrder(val *string)
	SortOrderInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Updated() *string
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
	ResetAfter()
	ResetBefore()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSortBy()
	ResetSortOrder()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataCloudflareCloudforceOneRequestMessage
type jsiiProxy_DataCloudflareCloudforceOneRequestMessage struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) After() *string {
	var returns *string
	_jsii_.Get(
		j,
		"after",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) AfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"afterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Author() *string {
	var returns *string
	_jsii_.Get(
		j,
		"author",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Before() *string {
	var returns *string
	_jsii_.Get(
		j,
		"before",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) BeforeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beforeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Created() *string {
	var returns *string
	_jsii_.Get(
		j,
		"created",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) IsFollowOnRequest() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isFollowOnRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Page() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"page",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) PageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"pageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) PerPage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) PerPageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"perPageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) RequestId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) RequestIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) SortBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) SortByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) SortOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) SortOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) Updated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updated",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/cloudforce_one_request_message cloudflare_cloudforce_one_request_message} Data Source.
func NewDataCloudflareCloudforceOneRequestMessage(scope constructs.Construct, id *string, config *DataCloudflareCloudforceOneRequestMessageConfig) DataCloudflareCloudforceOneRequestMessage {
	_init_.Initialize()

	if err := validateNewDataCloudflareCloudforceOneRequestMessageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareCloudforceOneRequestMessage{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequestMessage.DataCloudflareCloudforceOneRequestMessage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/data-sources/cloudforce_one_request_message cloudflare_cloudforce_one_request_message} Data Source.
func NewDataCloudflareCloudforceOneRequestMessage_Override(d DataCloudflareCloudforceOneRequestMessage, scope constructs.Construct, id *string, config *DataCloudflareCloudforceOneRequestMessageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequestMessage.DataCloudflareCloudforceOneRequestMessage",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetAfter(val *string) {
	if err := j.validateSetAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"after",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetBefore(val *string) {
	if err := j.validateSetBeforeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"before",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetPage(val *float64) {
	if err := j.validateSetPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"page",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetPerPage(val *float64) {
	if err := j.validateSetPerPageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"perPage",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetRequestId(val *string) {
	if err := j.validateSetRequestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestId",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetSortBy(val *string) {
	if err := j.validateSetSortByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortBy",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCloudforceOneRequestMessage)SetSortOrder(val *string) {
	if err := j.validateSetSortOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortOrder",
		val,
	)
}

// Generates CDKTF code for importing a DataCloudflareCloudforceOneRequestMessage resource upon running "cdktf plan <stack-name>".
func DataCloudflareCloudforceOneRequestMessage_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataCloudflareCloudforceOneRequestMessage_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequestMessage.DataCloudflareCloudforceOneRequestMessage",
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
func DataCloudflareCloudforceOneRequestMessage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareCloudforceOneRequestMessage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequestMessage.DataCloudflareCloudforceOneRequestMessage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareCloudforceOneRequestMessage_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareCloudforceOneRequestMessage_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequestMessage.DataCloudflareCloudforceOneRequestMessage",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataCloudflareCloudforceOneRequestMessage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataCloudflareCloudforceOneRequestMessage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequestMessage.DataCloudflareCloudforceOneRequestMessage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataCloudflareCloudforceOneRequestMessage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-cloudflare.dataCloudflareCloudforceOneRequestMessage.DataCloudflareCloudforceOneRequestMessage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ResetAfter() {
	_jsii_.InvokeVoid(
		d,
		"resetAfter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ResetBefore() {
	_jsii_.InvokeVoid(
		d,
		"resetBefore",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ResetSortBy() {
	_jsii_.InvokeVoid(
		d,
		"resetSortBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ResetSortOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetSortOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCloudforceOneRequestMessage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

