package googlecomputeregioninstancetemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeregioninstancetemplate/internal"
)

type GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference interface {
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
	InternalValue() *GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKey
	SetInternalValue(val *GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKey)
	KmsKeySelfLink() *string
	SetKmsKeySelfLink(val *string)
	KmsKeySelfLinkInput() *string
	KmsKeyServiceAccount() *string
	SetKmsKeyServiceAccount(val *string)
	KmsKeyServiceAccountInput() *string
	RawKey() *string
	SetRawKey(val *string)
	RawKeyInput() *string
	RsaEncryptedKey() *string
	SetRsaEncryptedKey(val *string)
	RsaEncryptedKeyInput() *string
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
	ResetKmsKeySelfLink()
	ResetKmsKeyServiceAccount()
	ResetRawKey()
	ResetRsaEncryptedKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference
type jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) InternalValue() *GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKey {
	var returns *GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKey
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) KmsKeySelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) KmsKeySelfLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeySelfLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) KmsKeyServiceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyServiceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) KmsKeyServiceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyServiceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) RawKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) RawKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rawKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) RsaEncryptedKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaEncryptedKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) RsaEncryptedKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rsaEncryptedKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionInstanceTemplate.GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference_Override(g GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeRegionInstanceTemplate.GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetInternalValue(val *GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKey) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetKmsKeySelfLink(val *string) {
	if err := j.validateSetKmsKeySelfLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeySelfLink",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetKmsKeyServiceAccount(val *string) {
	if err := j.validateSetKmsKeyServiceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyServiceAccount",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetRawKey(val *string) {
	if err := j.validateSetRawKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rawKey",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetRsaEncryptedKey(val *string) {
	if err := j.validateSetRsaEncryptedKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rsaEncryptedKey",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ResetKmsKeySelfLink() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeySelfLink",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ResetKmsKeyServiceAccount() {
	_jsii_.InvokeVoid(
		g,
		"resetKmsKeyServiceAccount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ResetRawKey() {
	_jsii_.InvokeVoid(
		g,
		"resetRawKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ResetRsaEncryptedKey() {
	_jsii_.InvokeVoid(
		g,
		"resetRsaEncryptedKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionInstanceTemplateDiskSourceSnapshotEncryptionKeyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

