package alloydbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/alloydbinstance/internal"
)

type AlloydbInstancePscInstanceConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedConsumerProjects() *[]*string
	SetAllowedConsumerProjects(val *[]*string)
	AllowedConsumerProjectsInput() *[]*string
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
	InternalValue() *AlloydbInstancePscInstanceConfig
	SetInternalValue(val *AlloydbInstancePscInstanceConfig)
	PscAutoConnections() AlloydbInstancePscInstanceConfigPscAutoConnectionsList
	PscAutoConnectionsInput() interface{}
	PscDnsName() *string
	PscInterfaceConfigs() AlloydbInstancePscInstanceConfigPscInterfaceConfigsList
	PscInterfaceConfigsInput() interface{}
	ServiceAttachmentLink() *string
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
	PutPscAutoConnections(value interface{})
	PutPscInterfaceConfigs(value interface{})
	ResetAllowedConsumerProjects()
	ResetPscAutoConnections()
	ResetPscInterfaceConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlloydbInstancePscInstanceConfigOutputReference
type jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) AllowedConsumerProjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedConsumerProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) AllowedConsumerProjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedConsumerProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) InternalValue() *AlloydbInstancePscInstanceConfig {
	var returns *AlloydbInstancePscInstanceConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) PscAutoConnections() AlloydbInstancePscInstanceConfigPscAutoConnectionsList {
	var returns AlloydbInstancePscInstanceConfigPscAutoConnectionsList
	_jsii_.Get(
		j,
		"pscAutoConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) PscAutoConnectionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pscAutoConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) PscDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pscDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) PscInterfaceConfigs() AlloydbInstancePscInstanceConfigPscInterfaceConfigsList {
	var returns AlloydbInstancePscInstanceConfigPscInterfaceConfigsList
	_jsii_.Get(
		j,
		"pscInterfaceConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) PscInterfaceConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pscInterfaceConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ServiceAttachmentLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAttachmentLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAlloydbInstancePscInstanceConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AlloydbInstancePscInstanceConfigOutputReference {
	_init_.Initialize()

	if err := validateNewAlloydbInstancePscInstanceConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.alloydbInstance.AlloydbInstancePscInstanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAlloydbInstancePscInstanceConfigOutputReference_Override(a AlloydbInstancePscInstanceConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.alloydbInstance.AlloydbInstancePscInstanceConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference)SetAllowedConsumerProjects(val *[]*string) {
	if err := j.validateSetAllowedConsumerProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedConsumerProjects",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference)SetInternalValue(val *AlloydbInstancePscInstanceConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) PutPscAutoConnections(value interface{}) {
	if err := a.validatePutPscAutoConnectionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPscAutoConnections",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) PutPscInterfaceConfigs(value interface{}) {
	if err := a.validatePutPscInterfaceConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPscInterfaceConfigs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ResetAllowedConsumerProjects() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedConsumerProjects",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ResetPscAutoConnections() {
	_jsii_.InvokeVoid(
		a,
		"resetPscAutoConnections",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ResetPscInterfaceConfigs() {
	_jsii_.InvokeVoid(
		a,
		"resetPscInterfaceConfigs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AlloydbInstancePscInstanceConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

