package osconfigv2policyorchestrator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/osconfigv2policyorchestrator/internal"
)

type OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference interface {
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
	LocationSelector() OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelectorOutputReference
	LocationSelectorInput() *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector
	ResourceHierarchySelector() OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelectorOutputReference
	ResourceHierarchySelectorInput() *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelector
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
	PutLocationSelector(value *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector)
	PutResourceHierarchySelector(value *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelector)
	ResetLocationSelector()
	ResetResourceHierarchySelector()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference
type jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) LocationSelector() OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelectorOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelectorOutputReference
	_jsii_.Get(
		j,
		"locationSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) LocationSelectorInput() *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector {
	var returns *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector
	_jsii_.Get(
		j,
		"locationSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ResourceHierarchySelector() OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelectorOutputReference {
	var returns OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelectorOutputReference
	_jsii_.Get(
		j,
		"resourceHierarchySelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ResourceHierarchySelectorInput() *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelector {
	var returns *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelector
	_jsii_.Get(
		j,
		"resourceHierarchySelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference {
	_init_.Initialize()

	if err := validateNewOsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestrator.OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewOsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference_Override(o OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestrator.OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		o,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) PutLocationSelector(value *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsLocationSelector) {
	if err := o.validatePutLocationSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLocationSelector",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) PutResourceHierarchySelector(value *OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsResourceHierarchySelector) {
	if err := o.validatePutResourceHierarchySelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putResourceHierarchySelector",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ResetLocationSelector() {
	_jsii_.InvokeVoid(
		o,
		"resetLocationSelector",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ResetResourceHierarchySelector() {
	_jsii_.InvokeVoid(
		o,
		"resetResourceHierarchySelector",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorOrchestrationScopeSelectorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

