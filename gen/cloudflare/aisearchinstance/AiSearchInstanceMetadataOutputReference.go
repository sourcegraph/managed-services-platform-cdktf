package aisearchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aisearchinstance/internal"
)

type AiSearchInstanceMetadataOutputReference interface {
	cdktf.ComplexObject
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
	CreatedFromAisearchWizard() interface{}
	SetCreatedFromAisearchWizard(val interface{})
	CreatedFromAisearchWizardInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	SearchForAgents() AiSearchInstanceMetadataSearchForAgentsOutputReference
	SearchForAgentsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerDomain() *string
	SetWorkerDomain(val *string)
	WorkerDomainInput() *string
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
	PutSearchForAgents(value *AiSearchInstanceMetadataSearchForAgents)
	ResetCreatedFromAisearchWizard()
	ResetSearchForAgents()
	ResetWorkerDomain()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiSearchInstanceMetadataOutputReference
type jsiiProxy_AiSearchInstanceMetadataOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) CreatedFromAisearchWizard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdFromAisearchWizard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) CreatedFromAisearchWizardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createdFromAisearchWizardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) SearchForAgents() AiSearchInstanceMetadataSearchForAgentsOutputReference {
	var returns AiSearchInstanceMetadataSearchForAgentsOutputReference
	_jsii_.Get(
		j,
		"searchForAgents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) SearchForAgentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchForAgentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) WorkerDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference) WorkerDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerDomainInput",
		&returns,
	)
	return returns
}


func NewAiSearchInstanceMetadataOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AiSearchInstanceMetadataOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchInstanceMetadataOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchInstanceMetadataOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstanceMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchInstanceMetadataOutputReference_Override(a AiSearchInstanceMetadataOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstanceMetadataOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference)SetCreatedFromAisearchWizard(val interface{}) {
	if err := j.validateSetCreatedFromAisearchWizardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdFromAisearchWizard",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceMetadataOutputReference)SetWorkerDomain(val *string) {
	if err := j.validateSetWorkerDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerDomain",
		val,
	)
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) PutSearchForAgents(value *AiSearchInstanceMetadataSearchForAgents) {
	if err := a.validatePutSearchForAgentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSearchForAgents",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) ResetCreatedFromAisearchWizard() {
	_jsii_.InvokeVoid(
		a,
		"resetCreatedFromAisearchWizard",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) ResetSearchForAgents() {
	_jsii_.InvokeVoid(
		a,
		"resetSearchForAgents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) ResetWorkerDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkerDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceMetadataOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

