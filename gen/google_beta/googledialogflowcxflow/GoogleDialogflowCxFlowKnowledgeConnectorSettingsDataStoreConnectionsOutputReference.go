package googledialogflowcxflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledialogflowcxflow/internal"
)

type GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference interface {
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
	DataStore() *string
	SetDataStore(val *string)
	DataStoreInput() *string
	DataStoreType() *string
	SetDataStoreType(val *string)
	DataStoreTypeInput() *string
	DocumentProcessingMode() *string
	SetDocumentProcessingMode(val *string)
	DocumentProcessingModeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
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
	ResetDataStore()
	ResetDataStoreType()
	ResetDocumentProcessingMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference
type jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) DataStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) DataStoreInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) DataStoreType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStoreType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) DataStoreTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStoreTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) DocumentProcessingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentProcessingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) DocumentProcessingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"documentProcessingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference_Override(g GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDialogflowCxFlow.GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetDataStore(val *string) {
	if err := j.validateSetDataStoreParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStore",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetDataStoreType(val *string) {
	if err := j.validateSetDataStoreTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStoreType",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetDocumentProcessingMode(val *string) {
	if err := j.validateSetDocumentProcessingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"documentProcessingMode",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) ResetDataStore() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStore",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) ResetDataStoreType() {
	_jsii_.InvokeVoid(
		g,
		"resetDataStoreType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) ResetDocumentProcessingMode() {
	_jsii_.InvokeVoid(
		g,
		"resetDocumentProcessingMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDialogflowCxFlowKnowledgeConnectorSettingsDataStoreConnectionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

