package googlecomputesecuritypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputesecuritypolicy/internal"
)

type GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference interface {
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
	InternalValue() *GoogleComputeSecurityPolicyAdvancedOptionsConfig
	SetInternalValue(val *GoogleComputeSecurityPolicyAdvancedOptionsConfig)
	JsonCustomConfig() GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	JsonCustomConfigInput() *GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	JsonParsing() *string
	SetJsonParsing(val *string)
	JsonParsingInput() *string
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	RequestBodyInspectionSize() *string
	SetRequestBodyInspectionSize(val *string)
	RequestBodyInspectionSizeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UserIpRequestHeaders() *[]*string
	SetUserIpRequestHeaders(val *[]*string)
	UserIpRequestHeadersInput() *[]*string
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
	PutJsonCustomConfig(value *GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig)
	ResetJsonCustomConfig()
	ResetJsonParsing()
	ResetLogLevel()
	ResetRequestBodyInspectionSize()
	ResetUserIpRequestHeaders()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference
type jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) InternalValue() *GoogleComputeSecurityPolicyAdvancedOptionsConfig {
	var returns *GoogleComputeSecurityPolicyAdvancedOptionsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfig() GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference {
	var returns GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfigOutputReference
	_jsii_.Get(
		j,
		"jsonCustomConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) JsonCustomConfigInput() *GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig {
	var returns *GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig
	_jsii_.Get(
		j,
		"jsonCustomConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsing() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) JsonParsingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonParsingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) RequestBodyInspectionSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInspectionSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) RequestBodyInspectionSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestBodyInspectionSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) UserIpRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIpRequestHeadersInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference_Override(g GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeSecurityPolicy.GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetInternalValue(val *GoogleComputeSecurityPolicyAdvancedOptionsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetJsonParsing(val *string) {
	if err := j.validateSetJsonParsingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsonParsing",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetRequestBodyInspectionSize(val *string) {
	if err := j.validateSetRequestBodyInspectionSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBodyInspectionSize",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference)SetUserIpRequestHeaders(val *[]*string) {
	if err := j.validateSetUserIpRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIpRequestHeaders",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) PutJsonCustomConfig(value *GoogleComputeSecurityPolicyAdvancedOptionsConfigJsonCustomConfig) {
	if err := g.validatePutJsonCustomConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putJsonCustomConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonCustomConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonCustomConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ResetJsonParsing() {
	_jsii_.InvokeVoid(
		g,
		"resetJsonParsing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ResetLogLevel() {
	_jsii_.InvokeVoid(
		g,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ResetRequestBodyInspectionSize() {
	_jsii_.InvokeVoid(
		g,
		"resetRequestBodyInspectionSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ResetUserIpRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetUserIpRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeSecurityPolicyAdvancedOptionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

