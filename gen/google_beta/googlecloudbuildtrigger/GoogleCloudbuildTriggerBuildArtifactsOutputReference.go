package googlecloudbuildtrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecloudbuildtrigger/internal"
)

type GoogleCloudbuildTriggerBuildArtifactsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Images() *[]*string
	SetImages(val *[]*string)
	ImagesInput() *[]*string
	InternalValue() *GoogleCloudbuildTriggerBuildArtifacts
	SetInternalValue(val *GoogleCloudbuildTriggerBuildArtifacts)
	MavenArtifacts() GoogleCloudbuildTriggerBuildArtifactsMavenArtifactsList
	MavenArtifactsInput() any
	NpmPackages() GoogleCloudbuildTriggerBuildArtifactsNpmPackagesList
	NpmPackagesInput() any
	Objects() GoogleCloudbuildTriggerBuildArtifactsObjectsOutputReference
	ObjectsInput() *GoogleCloudbuildTriggerBuildArtifactsObjects
	PythonPackages() GoogleCloudbuildTriggerBuildArtifactsPythonPackagesList
	PythonPackagesInput() any
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
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutMavenArtifacts(value any)
	PutNpmPackages(value any)
	PutObjects(value *GoogleCloudbuildTriggerBuildArtifactsObjects)
	PutPythonPackages(value any)
	ResetImages()
	ResetMavenArtifacts()
	ResetNpmPackages()
	ResetObjects()
	ResetPythonPackages()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleCloudbuildTriggerBuildArtifactsOutputReference
type jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) Images() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"images",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ImagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"imagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) InternalValue() *GoogleCloudbuildTriggerBuildArtifacts {
	var returns *GoogleCloudbuildTriggerBuildArtifacts
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) MavenArtifacts() GoogleCloudbuildTriggerBuildArtifactsMavenArtifactsList {
	var returns GoogleCloudbuildTriggerBuildArtifactsMavenArtifactsList
	_jsii_.Get(
		j,
		"mavenArtifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) MavenArtifactsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"mavenArtifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) NpmPackages() GoogleCloudbuildTriggerBuildArtifactsNpmPackagesList {
	var returns GoogleCloudbuildTriggerBuildArtifactsNpmPackagesList
	_jsii_.Get(
		j,
		"npmPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) NpmPackagesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"npmPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) Objects() GoogleCloudbuildTriggerBuildArtifactsObjectsOutputReference {
	var returns GoogleCloudbuildTriggerBuildArtifactsObjectsOutputReference
	_jsii_.Get(
		j,
		"objects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ObjectsInput() *GoogleCloudbuildTriggerBuildArtifactsObjects {
	var returns *GoogleCloudbuildTriggerBuildArtifactsObjects
	_jsii_.Get(
		j,
		"objectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) PythonPackages() GoogleCloudbuildTriggerBuildArtifactsPythonPackagesList {
	var returns GoogleCloudbuildTriggerBuildArtifactsPythonPackagesList
	_jsii_.Get(
		j,
		"pythonPackages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) PythonPackagesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pythonPackagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleCloudbuildTriggerBuildArtifactsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleCloudbuildTriggerBuildArtifactsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleCloudbuildTriggerBuildArtifactsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildArtifactsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleCloudbuildTriggerBuildArtifactsOutputReference_Override(g GoogleCloudbuildTriggerBuildArtifactsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleCloudbuildTrigger.GoogleCloudbuildTriggerBuildArtifactsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) SetImages(val *[]*string) {
	if err := j.validateSetImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"images",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) SetInternalValue(val *GoogleCloudbuildTriggerBuildArtifacts) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) PutMavenArtifacts(value any) {
	if err := g.validatePutMavenArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMavenArtifacts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) PutNpmPackages(value any) {
	if err := g.validatePutNpmPackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNpmPackages",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) PutObjects(value *GoogleCloudbuildTriggerBuildArtifactsObjects) {
	if err := g.validatePutObjectsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putObjects",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) PutPythonPackages(value any) {
	if err := g.validatePutPythonPackagesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPythonPackages",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ResetImages() {
	_jsii_.InvokeVoid(
		g,
		"resetImages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ResetMavenArtifacts() {
	_jsii_.InvokeVoid(
		g,
		"resetMavenArtifacts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ResetNpmPackages() {
	_jsii_.InvokeVoid(
		g,
		"resetNpmPackages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ResetObjects() {
	_jsii_.InvokeVoid(
		g,
		"resetObjects",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ResetPythonPackages() {
	_jsii_.InvokeVoid(
		g,
		"resetPythonPackages",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleCloudbuildTriggerBuildArtifactsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
