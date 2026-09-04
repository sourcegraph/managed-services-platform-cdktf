package aisearchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aisearchinstance/internal"
)

type AiSearchInstanceSourceParamsWebCrawlerOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ParseOptions() AiSearchInstanceSourceParamsWebCrawlerParseOptionsOutputReference
	ParseOptionsInput() interface{}
	ParseType() *string
	SetParseType(val *string)
	ParseTypeInput() *string
	StoreOptions() AiSearchInstanceSourceParamsWebCrawlerStoreOptionsOutputReference
	StoreOptionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutParseOptions(value *AiSearchInstanceSourceParamsWebCrawlerParseOptions)
	PutStoreOptions(value *AiSearchInstanceSourceParamsWebCrawlerStoreOptions)
	ResetParseOptions()
	ResetParseType()
	ResetStoreOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiSearchInstanceSourceParamsWebCrawlerOutputReference
type jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ParseOptions() AiSearchInstanceSourceParamsWebCrawlerParseOptionsOutputReference {
	var returns AiSearchInstanceSourceParamsWebCrawlerParseOptionsOutputReference
	_jsii_.Get(
		j,
		"parseOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ParseOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parseOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ParseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ParseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) StoreOptions() AiSearchInstanceSourceParamsWebCrawlerStoreOptionsOutputReference {
	var returns AiSearchInstanceSourceParamsWebCrawlerStoreOptionsOutputReference
	_jsii_.Get(
		j,
		"storeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) StoreOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchInstanceSourceParamsWebCrawlerOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AiSearchInstanceSourceParamsWebCrawlerOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchInstanceSourceParamsWebCrawlerOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstanceSourceParamsWebCrawlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchInstanceSourceParamsWebCrawlerOutputReference_Override(a AiSearchInstanceSourceParamsWebCrawlerOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstanceSourceParamsWebCrawlerOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference)SetParseType(val *string) {
	if err := j.validateSetParseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parseType",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) PutParseOptions(value *AiSearchInstanceSourceParamsWebCrawlerParseOptions) {
	if err := a.validatePutParseOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParseOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) PutStoreOptions(value *AiSearchInstanceSourceParamsWebCrawlerStoreOptions) {
	if err := a.validatePutStoreOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStoreOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ResetParseOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetParseOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ResetParseType() {
	_jsii_.InvokeVoid(
		a,
		"resetParseType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ResetStoreOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetStoreOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

