package googleartifactregistryrepository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleartifactregistryrepository/internal"
)

type GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference interface {
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
	CustomRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepositoryOutputReference
	CustomRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository
	SetInternalValue(val *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository)
	PublicRepository() *string
	SetPublicRepository(val *string)
	PublicRepositoryInput() *string
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
	PutCustomRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository)
	ResetCustomRepository()
	ResetPublicRepository()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference
type jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) CustomRepository() GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepositoryOutputReference {
	var returns GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepositoryOutputReference
	_jsii_.Get(
		j,
		"customRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) CustomRepositoryInput() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository
	_jsii_.Get(
		j,
		"customRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) InternalValue() *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository {
	var returns *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) PublicRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) PublicRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference_Override(g GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleArtifactRegistryRepository.GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference)SetInternalValue(val *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference)SetPublicRepository(val *string) {
	if err := j.validateSetPublicRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicRepository",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) PutCustomRepository(value *GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryCustomRepository) {
	if err := g.validatePutCustomRepositoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCustomRepository",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) ResetCustomRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) ResetPublicRepository() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicRepository",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleArtifactRegistryRepositoryRemoteRepositoryConfigPythonRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

