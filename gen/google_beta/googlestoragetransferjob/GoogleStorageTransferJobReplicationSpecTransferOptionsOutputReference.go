package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference interface {
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
	DeleteObjectsFromSourceAfterTransfer() interface{}
	SetDeleteObjectsFromSourceAfterTransfer(val interface{})
	DeleteObjectsFromSourceAfterTransferInput() interface{}
	DeleteObjectsUniqueInSink() interface{}
	SetDeleteObjectsUniqueInSink(val interface{})
	DeleteObjectsUniqueInSinkInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleStorageTransferJobReplicationSpecTransferOptions
	SetInternalValue(val *GoogleStorageTransferJobReplicationSpecTransferOptions)
	MetadataOptions() GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptionsOutputReference
	MetadataOptionsInput() *GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptions
	OverwriteObjectsAlreadyExistingInSink() interface{}
	SetOverwriteObjectsAlreadyExistingInSink(val interface{})
	OverwriteObjectsAlreadyExistingInSinkInput() interface{}
	OverwriteWhen() *string
	SetOverwriteWhen(val *string)
	OverwriteWhenInput() *string
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
	PutMetadataOptions(value *GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptions)
	ResetDeleteObjectsFromSourceAfterTransfer()
	ResetDeleteObjectsUniqueInSink()
	ResetMetadataOptions()
	ResetOverwriteObjectsAlreadyExistingInSink()
	ResetOverwriteWhen()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference
type jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) DeleteObjectsFromSourceAfterTransfer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteObjectsFromSourceAfterTransfer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) DeleteObjectsFromSourceAfterTransferInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteObjectsFromSourceAfterTransferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) DeleteObjectsUniqueInSink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteObjectsUniqueInSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) DeleteObjectsUniqueInSinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteObjectsUniqueInSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) InternalValue() *GoogleStorageTransferJobReplicationSpecTransferOptions {
	var returns *GoogleStorageTransferJobReplicationSpecTransferOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) MetadataOptions() GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptionsOutputReference {
	var returns GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptionsOutputReference
	_jsii_.Get(
		j,
		"metadataOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) MetadataOptionsInput() *GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptions {
	var returns *GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptions
	_jsii_.Get(
		j,
		"metadataOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) OverwriteObjectsAlreadyExistingInSink() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteObjectsAlreadyExistingInSink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) OverwriteObjectsAlreadyExistingInSinkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteObjectsAlreadyExistingInSinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) OverwriteWhen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteWhen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) OverwriteWhenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overwriteWhenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobReplicationSpecTransferOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleStorageTransferJob.GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference_Override(g GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleStorageTransferJob.GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetDeleteObjectsFromSourceAfterTransfer(val interface{}) {
	if err := j.validateSetDeleteObjectsFromSourceAfterTransferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteObjectsFromSourceAfterTransfer",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetDeleteObjectsUniqueInSink(val interface{}) {
	if err := j.validateSetDeleteObjectsUniqueInSinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteObjectsUniqueInSink",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetInternalValue(val *GoogleStorageTransferJobReplicationSpecTransferOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetOverwriteObjectsAlreadyExistingInSink(val interface{}) {
	if err := j.validateSetOverwriteObjectsAlreadyExistingInSinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteObjectsAlreadyExistingInSink",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetOverwriteWhen(val *string) {
	if err := j.validateSetOverwriteWhenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteWhen",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) PutMetadataOptions(value *GoogleStorageTransferJobReplicationSpecTransferOptionsMetadataOptions) {
	if err := g.validatePutMetadataOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMetadataOptions",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ResetDeleteObjectsFromSourceAfterTransfer() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteObjectsFromSourceAfterTransfer",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ResetDeleteObjectsUniqueInSink() {
	_jsii_.InvokeVoid(
		g,
		"resetDeleteObjectsUniqueInSink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ResetMetadataOptions() {
	_jsii_.InvokeVoid(
		g,
		"resetMetadataOptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ResetOverwriteObjectsAlreadyExistingInSink() {
	_jsii_.InvokeVoid(
		g,
		"resetOverwriteObjectsAlreadyExistingInSink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ResetOverwriteWhen() {
	_jsii_.InvokeVoid(
		g,
		"resetOverwriteWhen",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobReplicationSpecTransferOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

