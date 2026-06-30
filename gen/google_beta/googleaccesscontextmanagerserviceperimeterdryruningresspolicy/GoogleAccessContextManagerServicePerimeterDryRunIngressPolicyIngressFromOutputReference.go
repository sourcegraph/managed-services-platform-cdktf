package googleaccesscontextmanagerserviceperimeterdryruningresspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleaccesscontextmanagerserviceperimeterdryruningresspolicy/internal"
)

type GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference interface {
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
	Identities() *[]*string
	SetIdentities(val *[]*string)
	IdentitiesInput() *[]*string
	IdentityType() *string
	SetIdentityType(val *string)
	IdentityTypeInput() *string
	InternalValue() *GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFrom
	SetInternalValue(val *GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFrom)
	Sources() GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromSourcesList
	SourcesInput() any
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
	PutSources(value any)
	ResetIdentities()
	ResetIdentityType()
	ResetSources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference
type jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) Identities() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) IdentitiesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"identitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) IdentityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) IdentityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) InternalValue() *GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFrom {
	var returns *GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFrom
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) Sources() GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromSourcesList {
	var returns GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromSourcesList
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SourcesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"sourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerServicePerimeterDryRunIngressPolicy.GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference_Override(g GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerServicePerimeterDryRunIngressPolicy.GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SetIdentities(val *[]*string) {
	if err := j.validateSetIdentitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identities",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SetIdentityType(val *string) {
	if err := j.validateSetIdentityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityType",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SetInternalValue(val *GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFrom) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) PutSources(value any) {
	if err := g.validatePutSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSources",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) ResetIdentities() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentities",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) ResetIdentityType() {
	_jsii_.InvokeVoid(
		g,
		"resetIdentityType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) ResetSources() {
	_jsii_.InvokeVoid(
		g,
		"resetSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressFromOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
