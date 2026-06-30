package googlenetworkserviceshttproute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlenetworkserviceshttproute/internal"
)

type GoogleNetworkServicesHttpRouteRulesMatchesOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	FullPathMatch() *string
	SetFullPathMatch(val *string)
	FullPathMatchInput() *string
	Headers() GoogleNetworkServicesHttpRouteRulesMatchesHeadersList
	HeadersInput() any
	IgnoreCase() any
	SetIgnoreCase(val any)
	IgnoreCaseInput() any
	InternalValue() any
	SetInternalValue(val any)
	PrefixMatch() *string
	SetPrefixMatch(val *string)
	PrefixMatchInput() *string
	QueryParameters() GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersList
	QueryParametersInput() any
	RegexMatch() *string
	SetRegexMatch(val *string)
	RegexMatchInput() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutHeaders(value any)
	PutQueryParameters(value any)
	ResetFullPathMatch()
	ResetHeaders()
	ResetIgnoreCase()
	ResetPrefixMatch()
	ResetQueryParameters()
	ResetRegexMatch()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleNetworkServicesHttpRouteRulesMatchesOutputReference
type jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) FullPathMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullPathMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) FullPathMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullPathMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) Headers() GoogleNetworkServicesHttpRouteRulesMatchesHeadersList {
	var returns GoogleNetworkServicesHttpRouteRulesMatchesHeadersList
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) HeadersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) IgnoreCase() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreCase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) IgnoreCaseInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ignoreCaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) PrefixMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) PrefixMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) QueryParameters() GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersList {
	var returns GoogleNetworkServicesHttpRouteRulesMatchesQueryParametersList
	_jsii_.Get(
		j,
		"queryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) QueryParametersInput() any {
	var returns any
	_jsii_.Get(
		j,
		"queryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) RegexMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) RegexMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regexMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleNetworkServicesHttpRouteRulesMatchesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleNetworkServicesHttpRouteRulesMatchesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleNetworkServicesHttpRouteRulesMatchesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleNetworkServicesHttpRouteRulesMatchesOutputReference_Override(g GoogleNetworkServicesHttpRouteRulesMatchesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetworkServicesHttpRoute.GoogleNetworkServicesHttpRouteRulesMatchesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetFullPathMatch(val *string) {
	if err := j.validateSetFullPathMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullPathMatch",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetIgnoreCase(val any) {
	if err := j.validateSetIgnoreCaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreCase",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetPrefixMatch(val *string) {
	if err := j.validateSetPrefixMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefixMatch",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetRegexMatch(val *string) {
	if err := j.validateSetRegexMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regexMatch",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) PutHeaders(value any) {
	if err := g.validatePutHeadersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHeaders",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) PutQueryParameters(value any) {
	if err := g.validatePutQueryParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueryParameters",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ResetFullPathMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetFullPathMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ResetHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ResetIgnoreCase() {
	_jsii_.InvokeVoid(
		g,
		"resetIgnoreCase",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ResetPrefixMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetPrefixMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ResetQueryParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ResetRegexMatch() {
	_jsii_.InvokeVoid(
		g,
		"resetRegexMatch",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetworkServicesHttpRouteRulesMatchesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
