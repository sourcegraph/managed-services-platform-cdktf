package googlecomputeinstancefromtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputeinstancefromtemplate/internal"
)

type GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference interface {
	cdktf.ComplexObject
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
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
	EnableConfidentialCompute() interface{}
	SetEnableConfidentialCompute(val interface{})
	EnableConfidentialComputeInput() interface{}
	// Experimental.
	Fqn() *string
	Image() *string
	SetImage(val *string)
	ImageInput() *string
	InternalValue() *GoogleComputeInstanceFromTemplateBootDiskInitializeParams
	SetInternalValue(val *GoogleComputeInstanceFromTemplateBootDiskInitializeParams)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	ProvisionedIops() *float64
	SetProvisionedIops(val *float64)
	ProvisionedIopsInput() *float64
	ProvisionedThroughput() *float64
	SetProvisionedThroughput(val *float64)
	ProvisionedThroughputInput() *float64
	ResourceManagerTags() *map[string]*string
	SetResourceManagerTags(val *map[string]*string)
	ResourceManagerTagsInput() *map[string]*string
	ResourcePolicies() *[]*string
	SetResourcePolicies(val *[]*string)
	ResourcePoliciesInput() *[]*string
	Size() *float64
	SetSize(val *float64)
	SizeInput() *float64
	Snapshot() *string
	SetSnapshot(val *string)
	SnapshotInput() *string
	SourceImageEncryptionKey() GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference
	SourceImageEncryptionKeyInput() *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey
	SourceSnapshotEncryptionKey() GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKeyOutputReference
	SourceSnapshotEncryptionKeyInput() *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKey
	StoragePool() *string
	SetStoragePool(val *string)
	StoragePoolInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutSourceImageEncryptionKey(value *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey)
	PutSourceSnapshotEncryptionKey(value *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKey)
	ResetArchitecture()
	ResetEnableConfidentialCompute()
	ResetImage()
	ResetLabels()
	ResetProvisionedIops()
	ResetProvisionedThroughput()
	ResetResourceManagerTags()
	ResetResourcePolicies()
	ResetSize()
	ResetSnapshot()
	ResetSourceImageEncryptionKey()
	ResetSourceSnapshotEncryptionKey()
	ResetStoragePool()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference
type jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) EnableConfidentialCompute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialCompute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) EnableConfidentialComputeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableConfidentialComputeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Image() *string {
	var returns *string
	_jsii_.Get(
		j,
		"image",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) InternalValue() *GoogleComputeInstanceFromTemplateBootDiskInitializeParams {
	var returns *GoogleComputeInstanceFromTemplateBootDiskInitializeParams
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ProvisionedIops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ProvisionedIopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedIopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ProvisionedThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ProvisionedThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"provisionedThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResourceManagerTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResourceManagerTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"resourceManagerTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResourcePolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResourcePoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Size() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) SizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Snapshot() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) SnapshotInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) SourceImageEncryptionKey() GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference {
	var returns GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceImageEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) SourceImageEncryptionKeyInput() *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey {
	var returns *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey
	_jsii_.Get(
		j,
		"sourceImageEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) SourceSnapshotEncryptionKey() GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKeyOutputReference {
	var returns GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) SourceSnapshotEncryptionKeyInput() *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKey {
	var returns *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKey
	_jsii_.Get(
		j,
		"sourceSnapshotEncryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) StoragePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) StoragePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewGoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference_Override(g GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeInstanceFromTemplate.GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetEnableConfidentialCompute(val interface{}) {
	if err := j.validateSetEnableConfidentialComputeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableConfidentialCompute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetImage(val *string) {
	if err := j.validateSetImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"image",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetInternalValue(val *GoogleComputeInstanceFromTemplateBootDiskInitializeParams) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetProvisionedIops(val *float64) {
	if err := j.validateSetProvisionedIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedIops",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetProvisionedThroughput(val *float64) {
	if err := j.validateSetProvisionedThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionedThroughput",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetResourceManagerTags(val *map[string]*string) {
	if err := j.validateSetResourceManagerTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceManagerTags",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetResourcePolicies(val *[]*string) {
	if err := j.validateSetResourcePoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourcePolicies",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetSize(val *float64) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetSnapshot(val *string) {
	if err := j.validateSetSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshot",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetStoragePool(val *string) {
	if err := j.validateSetStoragePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePool",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) PutSourceImageEncryptionKey(value *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceImageEncryptionKey) {
	if err := g.validatePutSourceImageEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceImageEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) PutSourceSnapshotEncryptionKey(value *GoogleComputeInstanceFromTemplateBootDiskInitializeParamsSourceSnapshotEncryptionKey) {
	if err := g.validatePutSourceSnapshotEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSourceSnapshotEncryptionKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetArchitecture() {
	_jsii_.InvokeVoid(
		g,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetEnableConfidentialCompute() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableConfidentialCompute",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetImage() {
	_jsii_.InvokeVoid(
		g,
		"resetImage",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetProvisionedIops() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedIops",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetProvisionedThroughput() {
	_jsii_.InvokeVoid(
		g,
		"resetProvisionedThroughput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetResourceManagerTags() {
	_jsii_.InvokeVoid(
		g,
		"resetResourceManagerTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetSize() {
	_jsii_.InvokeVoid(
		g,
		"resetSize",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetSnapshot() {
	_jsii_.InvokeVoid(
		g,
		"resetSnapshot",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetSourceImageEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceImageEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetSourceSnapshotEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceSnapshotEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetStoragePool() {
	_jsii_.InvokeVoid(
		g,
		"resetStoragePool",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		g,
		"resetType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeInstanceFromTemplateBootDiskInitializeParamsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

