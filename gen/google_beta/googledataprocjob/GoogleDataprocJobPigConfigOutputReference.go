package googledataprocjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledataprocjob/internal"
)

type GoogleDataprocJobPigConfigOutputReference interface {
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
	ContinueOnFailure() interface{}
	SetContinueOnFailure(val interface{})
	ContinueOnFailureInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleDataprocJobPigConfig
	SetInternalValue(val *GoogleDataprocJobPigConfig)
	JarFileUris() *[]*string
	SetJarFileUris(val *[]*string)
	JarFileUrisInput() *[]*string
	LoggingConfig() GoogleDataprocJobPigConfigLoggingConfigOutputReference
	LoggingConfigInput() *GoogleDataprocJobPigConfigLoggingConfig
	Properties() *map[string]*string
	SetProperties(val *map[string]*string)
	PropertiesInput() *map[string]*string
	QueryFileUri() *string
	SetQueryFileUri(val *string)
	QueryFileUriInput() *string
	QueryList() *[]*string
	SetQueryList(val *[]*string)
	QueryListInput() *[]*string
	ScriptVariables() *map[string]*string
	SetScriptVariables(val *map[string]*string)
	ScriptVariablesInput() *map[string]*string
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
	PutLoggingConfig(value *GoogleDataprocJobPigConfigLoggingConfig)
	ResetContinueOnFailure()
	ResetJarFileUris()
	ResetLoggingConfig()
	ResetProperties()
	ResetQueryFileUri()
	ResetQueryList()
	ResetScriptVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocJobPigConfigOutputReference
type jsiiProxy_GoogleDataprocJobPigConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ContinueOnFailure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ContinueOnFailureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continueOnFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) InternalValue() *GoogleDataprocJobPigConfig {
	var returns *GoogleDataprocJobPigConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) JarFileUris() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUris",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) JarFileUrisInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jarFileUrisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) LoggingConfig() GoogleDataprocJobPigConfigLoggingConfigOutputReference {
	var returns GoogleDataprocJobPigConfigLoggingConfigOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) LoggingConfigInput() *GoogleDataprocJobPigConfigLoggingConfig {
	var returns *GoogleDataprocJobPigConfigLoggingConfig
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) Properties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) PropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"propertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) QueryFileUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) QueryFileUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryFileUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) QueryList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) QueryListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ScriptVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ScriptVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"scriptVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDataprocJobPigConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataprocJobPigConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataprocJobPigConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocJobPigConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataprocJob.GoogleDataprocJobPigConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataprocJobPigConfigOutputReference_Override(g GoogleDataprocJobPigConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataprocJob.GoogleDataprocJobPigConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetContinueOnFailure(val interface{}) {
	if err := j.validateSetContinueOnFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"continueOnFailure",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetInternalValue(val *GoogleDataprocJobPigConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetJarFileUris(val *[]*string) {
	if err := j.validateSetJarFileUrisParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jarFileUris",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetProperties(val *map[string]*string) {
	if err := j.validateSetPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"properties",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetQueryFileUri(val *string) {
	if err := j.validateSetQueryFileUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryFileUri",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetQueryList(val *[]*string) {
	if err := j.validateSetQueryListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryList",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetScriptVariables(val *map[string]*string) {
	if err := j.validateSetScriptVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scriptVariables",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocJobPigConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) PutLoggingConfig(value *GoogleDataprocJobPigConfigLoggingConfig) {
	if err := g.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ResetContinueOnFailure() {
	_jsii_.InvokeVoid(
		g,
		"resetContinueOnFailure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ResetJarFileUris() {
	_jsii_.InvokeVoid(
		g,
		"resetJarFileUris",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ResetProperties() {
	_jsii_.InvokeVoid(
		g,
		"resetProperties",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ResetQueryFileUri() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryFileUri",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ResetQueryList() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ResetScriptVariables() {
	_jsii_.InvokeVoid(
		g,
		"resetScriptVariables",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocJobPigConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

