package googlestoragetransferjob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlestoragetransferjob/internal"
)

type GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference interface {
	cdktf.ComplexObject
	AwsAccessKey() GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutputReference
	AwsAccessKeyInput() *GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey
	BucketName() *string
	SetBucketName(val *string)
	BucketNameInput() *string
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
	InternalValue() *GoogleStorageTransferJobTransferSpecAwsS3DataSource
	SetInternalValue(val *GoogleStorageTransferJobTransferSpecAwsS3DataSource)
	ManagedPrivateNetwork() interface{}
	SetManagedPrivateNetwork(val interface{})
	ManagedPrivateNetworkInput() interface{}
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
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
	PutAwsAccessKey(value *GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey)
	ResetAwsAccessKey()
	ResetManagedPrivateNetwork()
	ResetPath()
	ResetRoleArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference
type jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) AwsAccessKey() GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutputReference {
	var returns GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKeyOutputReference
	_jsii_.Get(
		j,
		"awsAccessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) AwsAccessKeyInput() *GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey {
	var returns *GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey
	_jsii_.Get(
		j,
		"awsAccessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) InternalValue() *GoogleStorageTransferJobTransferSpecAwsS3DataSource {
	var returns *GoogleStorageTransferJobTransferSpecAwsS3DataSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ManagedPrivateNetwork() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedPrivateNetwork",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ManagedPrivateNetworkInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"managedPrivateNetworkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference_Override(g GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleStorageTransferJob.GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetBucketName(val *string) {
	if err := j.validateSetBucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketName",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetInternalValue(val *GoogleStorageTransferJobTransferSpecAwsS3DataSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetManagedPrivateNetwork(val interface{}) {
	if err := j.validateSetManagedPrivateNetworkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPrivateNetwork",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) PutAwsAccessKey(value *GoogleStorageTransferJobTransferSpecAwsS3DataSourceAwsAccessKey) {
	if err := g.validatePutAwsAccessKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAwsAccessKey",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ResetAwsAccessKey() {
	_jsii_.InvokeVoid(
		g,
		"resetAwsAccessKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ResetManagedPrivateNetwork() {
	_jsii_.InvokeVoid(
		g,
		"resetManagedPrivateNetwork",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		g,
		"resetPath",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		g,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageTransferJobTransferSpecAwsS3DataSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

