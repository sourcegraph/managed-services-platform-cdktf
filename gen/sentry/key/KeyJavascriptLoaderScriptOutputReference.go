package key

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/sentry/key/internal"
)

type KeyJavascriptLoaderScriptOutputReference interface {
	cdktf.ComplexObject
	BrowserSdkVersion() *string
	SetBrowserSdkVersion(val *string)
	BrowserSdkVersionInput() *string
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
	DebugEnabled() any
	SetDebugEnabled(val any)
	DebugEnabledInput() any
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	PerformanceMonitoringEnabled() any
	SetPerformanceMonitoringEnabled(val any)
	PerformanceMonitoringEnabledInput() any
	SessionReplayEnabled() any
	SetSessionReplayEnabled(val any)
	SessionReplayEnabledInput() any
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
	ResetBrowserSdkVersion()
	ResetDebugEnabled()
	ResetPerformanceMonitoringEnabled()
	ResetSessionReplayEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KeyJavascriptLoaderScriptOutputReference
type jsiiProxy_KeyJavascriptLoaderScriptOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) BrowserSdkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserSdkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) BrowserSdkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"browserSdkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) DebugEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"debugEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) DebugEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"debugEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) PerformanceMonitoringEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"performanceMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) PerformanceMonitoringEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"performanceMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SessionReplayEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"sessionReplayEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SessionReplayEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sessionReplayEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewKeyJavascriptLoaderScriptOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KeyJavascriptLoaderScriptOutputReference {
	_init_.Initialize()

	if err := validateNewKeyJavascriptLoaderScriptOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyJavascriptLoaderScriptOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-sentry.key.KeyJavascriptLoaderScriptOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKeyJavascriptLoaderScriptOutputReference_Override(k KeyJavascriptLoaderScriptOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-sentry.key.KeyJavascriptLoaderScriptOutputReference",
		[]any{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetBrowserSdkVersion(val *string) {
	if err := j.validateSetBrowserSdkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"browserSdkVersion",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetDebugEnabled(val any) {
	if err := j.validateSetDebugEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"debugEnabled",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetPerformanceMonitoringEnabled(val any) {
	if err := j.validateSetPerformanceMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetSessionReplayEnabled(val any) {
	if err := j.validateSetSessionReplayEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionReplayEnabled",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ResetBrowserSdkVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetBrowserSdkVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ResetDebugEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetDebugEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ResetPerformanceMonitoringEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetPerformanceMonitoringEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ResetSessionReplayEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetSessionReplayEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		k,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyJavascriptLoaderScriptOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
