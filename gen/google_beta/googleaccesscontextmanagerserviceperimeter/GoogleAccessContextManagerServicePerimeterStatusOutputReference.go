package googleaccesscontextmanagerserviceperimeter

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleaccesscontextmanagerserviceperimeter/internal"
)

type GoogleAccessContextManagerServicePerimeterStatusOutputReference interface {
	cdktf.ComplexObject
	AccessLevels() *[]*string
	SetAccessLevels(val *[]*string)
	AccessLevelsInput() *[]*string
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
	EgressPolicies() GoogleAccessContextManagerServicePerimeterStatusEgressPoliciesList
	EgressPoliciesInput() any
	// Experimental.
	Fqn() *string
	IngressPolicies() GoogleAccessContextManagerServicePerimeterStatusIngressPoliciesList
	IngressPoliciesInput() any
	InternalValue() *GoogleAccessContextManagerServicePerimeterStatus
	SetInternalValue(val *GoogleAccessContextManagerServicePerimeterStatus)
	Resources() *[]*string
	SetResources(val *[]*string)
	ResourcesInput() *[]*string
	RestrictedServices() *[]*string
	SetRestrictedServices(val *[]*string)
	RestrictedServicesInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcAccessibleServices() GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServicesOutputReference
	VpcAccessibleServicesInput() *GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServices
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
	PutEgressPolicies(value any)
	PutIngressPolicies(value any)
	PutVpcAccessibleServices(value *GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServices)
	ResetAccessLevels()
	ResetEgressPolicies()
	ResetIngressPolicies()
	ResetResources()
	ResetRestrictedServices()
	ResetVpcAccessibleServices()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerServicePerimeterStatusOutputReference
type jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) AccessLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) AccessLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accessLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) EgressPolicies() GoogleAccessContextManagerServicePerimeterStatusEgressPoliciesList {
	var returns GoogleAccessContextManagerServicePerimeterStatusEgressPoliciesList
	_jsii_.Get(
		j,
		"egressPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) EgressPoliciesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"egressPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) IngressPolicies() GoogleAccessContextManagerServicePerimeterStatusIngressPoliciesList {
	var returns GoogleAccessContextManagerServicePerimeterStatusIngressPoliciesList
	_jsii_.Get(
		j,
		"ingressPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) IngressPoliciesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"ingressPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) InternalValue() *GoogleAccessContextManagerServicePerimeterStatus {
	var returns *GoogleAccessContextManagerServicePerimeterStatus
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) RestrictedServices() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) RestrictedServicesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedServicesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) VpcAccessibleServices() GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServicesOutputReference {
	var returns GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServicesOutputReference
	_jsii_.Get(
		j,
		"vpcAccessibleServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) VpcAccessibleServicesInput() *GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServices {
	var returns *GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServices
	_jsii_.Get(
		j,
		"vpcAccessibleServicesInput",
		&returns,
	)
	return returns
}

func NewGoogleAccessContextManagerServicePerimeterStatusOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAccessContextManagerServicePerimeterStatusOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerServicePerimeterStatusOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerServicePerimeter.GoogleAccessContextManagerServicePerimeterStatusOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerServicePerimeterStatusOutputReference_Override(g GoogleAccessContextManagerServicePerimeterStatusOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerServicePerimeter.GoogleAccessContextManagerServicePerimeterStatusOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetAccessLevels(val *[]*string) {
	if err := j.validateSetAccessLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessLevels",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetInternalValue(val *GoogleAccessContextManagerServicePerimeterStatus) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetRestrictedServices(val *[]*string) {
	if err := j.validateSetRestrictedServicesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedServices",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) PutEgressPolicies(value any) {
	if err := g.validatePutEgressPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEgressPolicies",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) PutIngressPolicies(value any) {
	if err := g.validatePutIngressPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIngressPolicies",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) PutVpcAccessibleServices(value *GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServices) {
	if err := g.validatePutVpcAccessibleServicesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcAccessibleServices",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ResetAccessLevels() {
	_jsii_.InvokeVoid(
		g,
		"resetAccessLevels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ResetEgressPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetEgressPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ResetIngressPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetIngressPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		g,
		"resetResources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ResetRestrictedServices() {
	_jsii_.InvokeVoid(
		g,
		"resetRestrictedServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ResetVpcAccessibleServices() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcAccessibleServices",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleAccessContextManagerServicePerimeterStatusOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
