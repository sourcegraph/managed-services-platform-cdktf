package aisearchinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/aisearchinstance/internal"
)

type AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference interface {
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
	Depth() *float64
	SetDepth(val *float64)
	DepthInput() *float64
	// Experimental.
	Fqn() *string
	IncludeExternalLinks() interface{}
	SetIncludeExternalLinks(val interface{})
	IncludeExternalLinksInput() interface{}
	IncludeSubdomains() interface{}
	SetIncludeSubdomains(val interface{})
	IncludeSubdomainsInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxAge() *float64
	SetMaxAge(val *float64)
	MaxAgeInput() *float64
	Source() *string
	SetSource(val *string)
	SourceInput() *string
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
	ResetDepth()
	ResetIncludeExternalLinks()
	ResetIncludeSubdomains()
	ResetMaxAge()
	ResetSource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference
type jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) Depth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"depth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) DepthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"depthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) IncludeExternalLinks() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExternalLinks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) IncludeExternalLinksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExternalLinksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) IncludeSubdomains() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) IncludeSubdomainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) MaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) MaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewAiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference_Override(a AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.aiSearchInstance.AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetDepth(val *float64) {
	if err := j.validateSetDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"depth",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetIncludeExternalLinks(val interface{}) {
	if err := j.validateSetIncludeExternalLinksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExternalLinks",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetIncludeSubdomains(val interface{}) {
	if err := j.validateSetIncludeSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubdomains",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetMaxAge(val *float64) {
	if err := j.validateSetMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAge",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ResetDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ResetIncludeExternalLinks() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeExternalLinks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ResetIncludeSubdomains() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeSubdomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ResetMaxAge() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAge",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AiSearchInstanceSourceParamsWebCrawlerCrawlOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

