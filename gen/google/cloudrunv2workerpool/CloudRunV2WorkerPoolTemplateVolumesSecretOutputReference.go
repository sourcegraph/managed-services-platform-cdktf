package cloudrunv2workerpool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/cloudrunv2workerpool/internal"
)

type CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference interface {
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
	DefaultMode() *float64
	SetDefaultMode(val *float64)
	DefaultModeInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *CloudRunV2WorkerPoolTemplateVolumesSecret
	SetInternalValue(val *CloudRunV2WorkerPoolTemplateVolumesSecret)
	Items() CloudRunV2WorkerPoolTemplateVolumesSecretItemsList
	ItemsInput() interface{}
	Secret() *string
	SetSecret(val *string)
	SecretInput() *string
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
	PutItems(value interface{})
	ResetDefaultMode()
	ResetItems()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference
type jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) DefaultMode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) DefaultModeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) InternalValue() *CloudRunV2WorkerPoolTemplateVolumesSecret {
	var returns *CloudRunV2WorkerPoolTemplateVolumesSecret
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) Items() CloudRunV2WorkerPoolTemplateVolumesSecretItemsList {
	var returns CloudRunV2WorkerPoolTemplateVolumesSecretItemsList
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) ItemsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) Secret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) SecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCloudRunV2WorkerPoolTemplateVolumesSecretOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference {
	_init_.Initialize()

	if err := validateNewCloudRunV2WorkerPoolTemplateVolumesSecretOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudRunV2WorkerPoolTemplateVolumesSecretOutputReference_Override(c CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cloudRunV2WorkerPool.CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference)SetDefaultMode(val *float64) {
	if err := j.validateSetDefaultModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMode",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference)SetInternalValue(val *CloudRunV2WorkerPoolTemplateVolumesSecret) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference)SetSecret(val *string) {
	if err := j.validateSetSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secret",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) PutItems(value interface{}) {
	if err := c.validatePutItemsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putItems",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) ResetDefaultMode() {
	_jsii_.InvokeVoid(
		c,
		"resetDefaultMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) ResetItems() {
	_jsii_.InvokeVoid(
		c,
		"resetItems",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudRunV2WorkerPoolTemplateVolumesSecretOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

