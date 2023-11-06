package googlehealthcarefhirstore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlehealthcarefhirstore/internal"
)

type GoogleHealthcareFhirStoreNotificationConfigsOutputReference interface {
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
	PubsubTopic() *string
	SetPubsubTopic(val *string)
	PubsubTopicInput() *string
	SendFullResource() interface{}
	SetSendFullResource(val interface{})
	SendFullResourceInput() interface{}
	SendPreviousResourceOnDelete() interface{}
	SetSendPreviousResourceOnDelete(val interface{})
	SendPreviousResourceOnDeleteInput() interface{}
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
	ResetSendFullResource()
	ResetSendPreviousResourceOnDelete()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHealthcareFhirStoreNotificationConfigsOutputReference
type jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) PubsubTopic() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) PubsubTopicInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pubsubTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) SendFullResource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendFullResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) SendFullResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendFullResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) SendPreviousResourceOnDelete() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendPreviousResourceOnDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) SendPreviousResourceOnDeleteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sendPreviousResourceOnDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHealthcareFhirStoreNotificationConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleHealthcareFhirStoreNotificationConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHealthcareFhirStoreNotificationConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleHealthcareFhirStore.GoogleHealthcareFhirStoreNotificationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleHealthcareFhirStoreNotificationConfigsOutputReference_Override(g GoogleHealthcareFhirStoreNotificationConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleHealthcareFhirStore.GoogleHealthcareFhirStoreNotificationConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetPubsubTopic(val *string) {
	if err := j.validateSetPubsubTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pubsubTopic",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetSendFullResource(val interface{}) {
	if err := j.validateSetSendFullResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendFullResource",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetSendPreviousResourceOnDelete(val interface{}) {
	if err := j.validateSetSendPreviousResourceOnDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendPreviousResourceOnDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) ResetSendFullResource() {
	_jsii_.InvokeVoid(
		g,
		"resetSendFullResource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) ResetSendPreviousResourceOnDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetSendPreviousResourceOnDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHealthcareFhirStoreNotificationConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

