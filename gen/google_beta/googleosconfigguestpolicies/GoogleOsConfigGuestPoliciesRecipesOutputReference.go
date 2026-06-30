package googleosconfigguestpolicies

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleosconfigguestpolicies/internal"
)

type GoogleOsConfigGuestPoliciesRecipesOutputReference interface {
	cdktf.ComplexObject
	Artifacts() GoogleOsConfigGuestPoliciesRecipesArtifactsList
	ArtifactsInput() any
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
	DesiredState() *string
	SetDesiredState(val *string)
	DesiredStateInput() *string
	// Experimental.
	Fqn() *string
	InstallSteps() GoogleOsConfigGuestPoliciesRecipesInstallStepsList
	InstallStepsInput() any
	InternalValue() any
	SetInternalValue(val any)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UpdateSteps() GoogleOsConfigGuestPoliciesRecipesUpdateStepsList
	UpdateStepsInput() any
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutArtifacts(value any)
	PutInstallSteps(value any)
	PutUpdateSteps(value any)
	ResetArtifacts()
	ResetDesiredState()
	ResetInstallSteps()
	ResetUpdateSteps()
	ResetVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOsConfigGuestPoliciesRecipesOutputReference
type jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) Artifacts() GoogleOsConfigGuestPoliciesRecipesArtifactsList {
	var returns GoogleOsConfigGuestPoliciesRecipesArtifactsList
	_jsii_.Get(
		j,
		"artifacts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ArtifactsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"artifactsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) InstallSteps() GoogleOsConfigGuestPoliciesRecipesInstallStepsList {
	var returns GoogleOsConfigGuestPoliciesRecipesInstallStepsList
	_jsii_.Get(
		j,
		"installSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) InstallStepsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"installStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) UpdateSteps() GoogleOsConfigGuestPoliciesRecipesUpdateStepsList {
	var returns GoogleOsConfigGuestPoliciesRecipesUpdateStepsList
	_jsii_.Get(
		j,
		"updateSteps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) UpdateStepsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"updateStepsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func NewGoogleOsConfigGuestPoliciesRecipesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleOsConfigGuestPoliciesRecipesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOsConfigGuestPoliciesRecipesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleOsConfigGuestPoliciesRecipesOutputReference_Override(g GoogleOsConfigGuestPoliciesRecipesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleOsConfigGuestPolicies.GoogleOsConfigGuestPoliciesRecipesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) PutArtifacts(value any) {
	if err := g.validatePutArtifactsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putArtifacts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) PutInstallSteps(value any) {
	if err := g.validatePutInstallStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInstallSteps",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) PutUpdateSteps(value any) {
	if err := g.validatePutUpdateStepsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpdateSteps",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ResetArtifacts() {
	_jsii_.InvokeVoid(
		g,
		"resetArtifacts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ResetDesiredState() {
	_jsii_.InvokeVoid(
		g,
		"resetDesiredState",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ResetInstallSteps() {
	_jsii_.InvokeVoid(
		g,
		"resetInstallSteps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ResetUpdateSteps() {
	_jsii_.InvokeVoid(
		g,
		"resetUpdateSteps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleOsConfigGuestPoliciesRecipesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
