package pipelinesink

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/pipelinesink/internal"
)

type PipelineSinkConfigAOutputReference interface {
	cdktf.ComplexObject
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	Bucket() *string
	SetBucket(val *string)
	BucketInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Credentials() PipelineSinkConfigCredentialsOutputReference
	CredentialsInput() interface{}
	FileNaming() PipelineSinkConfigFileNamingOutputReference
	FileNamingInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Jurisdiction() *string
	SetJurisdiction(val *string)
	JurisdictionInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	Partitioning() PipelineSinkConfigPartitioningOutputReference
	PartitioningInput() interface{}
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RollingPolicy() PipelineSinkConfigRollingPolicyOutputReference
	RollingPolicyInput() interface{}
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutCredentials(value *PipelineSinkConfigCredentials)
	PutFileNaming(value *PipelineSinkConfigFileNaming)
	PutPartitioning(value *PipelineSinkConfigPartitioning)
	PutRollingPolicy(value *PipelineSinkConfigRollingPolicy)
	ResetCredentials()
	ResetFileNaming()
	ResetJurisdiction()
	ResetNamespace()
	ResetPartitioning()
	ResetPath()
	ResetRollingPolicy()
	ResetTableName()
	ResetToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PipelineSinkConfigAOutputReference
type jsiiProxy_PipelineSinkConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Credentials() PipelineSinkConfigCredentialsOutputReference {
	var returns PipelineSinkConfigCredentialsOutputReference
	_jsii_.Get(
		j,
		"credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) CredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) FileNaming() PipelineSinkConfigFileNamingOutputReference {
	var returns PipelineSinkConfigFileNamingOutputReference
	_jsii_.Get(
		j,
		"fileNaming",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) FileNamingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileNamingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Jurisdiction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jurisdiction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) JurisdictionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jurisdictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Partitioning() PipelineSinkConfigPartitioningOutputReference {
	var returns PipelineSinkConfigPartitioningOutputReference
	_jsii_.Get(
		j,
		"partitioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) PartitioningInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) RollingPolicy() PipelineSinkConfigRollingPolicyOutputReference {
	var returns PipelineSinkConfigRollingPolicyOutputReference
	_jsii_.Get(
		j,
		"rollingPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) RollingPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rollingPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


func NewPipelineSinkConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PipelineSinkConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewPipelineSinkConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PipelineSinkConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pipelineSink.PipelineSinkConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPipelineSinkConfigAOutputReference_Override(p PipelineSinkConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.pipelineSink.PipelineSinkConfigAOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetJurisdiction(val *string) {
	if err := j.validateSetJurisdictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jurisdiction",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PipelineSinkConfigAOutputReference)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) PutCredentials(value *PipelineSinkConfigCredentials) {
	if err := p.validatePutCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCredentials",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) PutFileNaming(value *PipelineSinkConfigFileNaming) {
	if err := p.validatePutFileNamingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFileNaming",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) PutPartitioning(value *PipelineSinkConfigPartitioning) {
	if err := p.validatePutPartitioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putPartitioning",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) PutRollingPolicy(value *PipelineSinkConfigRollingPolicy) {
	if err := p.validatePutRollingPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putRollingPolicy",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetCredentials() {
	_jsii_.InvokeVoid(
		p,
		"resetCredentials",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetFileNaming() {
	_jsii_.InvokeVoid(
		p,
		"resetFileNaming",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetJurisdiction() {
	_jsii_.InvokeVoid(
		p,
		"resetJurisdiction",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetPartitioning() {
	_jsii_.InvokeVoid(
		p,
		"resetPartitioning",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		p,
		"resetPath",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetRollingPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetRollingPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetTableName() {
	_jsii_.InvokeVoid(
		p,
		"resetTableName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ResetToken() {
	_jsii_.InvokeVoid(
		p,
		"resetToken",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PipelineSinkConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

