package dialogflowcxgenerativesettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/dialogflowcxgenerativesettings/internal"
)

type DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference interface {
	cdktf.ComplexObject
	Agent() *string
	SetAgent(val *string)
	AgentIdentity() *string
	SetAgentIdentity(val *string)
	AgentIdentityInput() *string
	AgentInput() *string
	AgentScope() *string
	SetAgentScope(val *string)
	AgentScopeInput() *string
	Business() *string
	SetBusiness(val *string)
	BusinessDescription() *string
	SetBusinessDescription(val *string)
	BusinessDescriptionInput() *string
	BusinessInput() *string
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
	DisableDataStoreFallback() interface{}
	SetDisableDataStoreFallback(val interface{})
	DisableDataStoreFallbackInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *DialogflowCxGenerativeSettingsKnowledgeConnectorSettings
	SetInternalValue(val *DialogflowCxGenerativeSettingsKnowledgeConnectorSettings)
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
	ResetAgent()
	ResetAgentIdentity()
	ResetAgentScope()
	ResetBusiness()
	ResetBusinessDescription()
	ResetDisableDataStoreFallback()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference
type jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) Agent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) AgentIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) AgentIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) AgentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) AgentScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) AgentScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) Business() *string {
	var returns *string
	_jsii_.Get(
		j,
		"business",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) BusinessDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) BusinessDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) BusinessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) DisableDataStoreFallback() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDataStoreFallback",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) DisableDataStoreFallbackInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableDataStoreFallbackInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) InternalValue() *DialogflowCxGenerativeSettingsKnowledgeConnectorSettings {
	var returns *DialogflowCxGenerativeSettingsKnowledgeConnectorSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxGenerativeSettings.DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference_Override(d DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxGenerativeSettings.DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetAgent(val *string) {
	if err := j.validateSetAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agent",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetAgentIdentity(val *string) {
	if err := j.validateSetAgentIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentIdentity",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetAgentScope(val *string) {
	if err := j.validateSetAgentScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentScope",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetBusiness(val *string) {
	if err := j.validateSetBusinessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"business",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetBusinessDescription(val *string) {
	if err := j.validateSetBusinessDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessDescription",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetDisableDataStoreFallback(val interface{}) {
	if err := j.validateSetDisableDataStoreFallbackParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableDataStoreFallback",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetInternalValue(val *DialogflowCxGenerativeSettingsKnowledgeConnectorSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ResetAgent() {
	_jsii_.InvokeVoid(
		d,
		"resetAgent",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ResetAgentIdentity() {
	_jsii_.InvokeVoid(
		d,
		"resetAgentIdentity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ResetAgentScope() {
	_jsii_.InvokeVoid(
		d,
		"resetAgentScope",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ResetBusiness() {
	_jsii_.InvokeVoid(
		d,
		"resetBusiness",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ResetBusinessDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetBusinessDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ResetDisableDataStoreFallback() {
	_jsii_.InvokeVoid(
		d,
		"resetDisableDataStoreFallback",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxGenerativeSettingsKnowledgeConnectorSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

