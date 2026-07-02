package datacloudflareloadbalancers

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflareloadbalancers/internal"
)

type DataCloudflareLoadBalancersResultRulesOverridesOutputReference interface {
	cdktf.ComplexObject
	AdaptiveRouting() DataCloudflareLoadBalancersResultRulesOverridesAdaptiveRoutingOutputReference
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
	CountryPools() any
	SetCountryPools(val any)
	CountryPoolsInput() any
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultPools() *[]*string
	FallbackPool() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataCloudflareLoadBalancersResultRulesOverrides
	SetInternalValue(val *DataCloudflareLoadBalancersResultRulesOverrides)
	LocationStrategy() DataCloudflareLoadBalancersResultRulesOverridesLocationStrategyOutputReference
	PopPools() any
	SetPopPools(val any)
	PopPoolsInput() any
	RandomSteering() DataCloudflareLoadBalancersResultRulesOverridesRandomSteeringOutputReference
	RegionPools() any
	SetRegionPools(val any)
	RegionPoolsInput() any
	SessionAffinity() *string
	SessionAffinityAttributes() DataCloudflareLoadBalancersResultRulesOverridesSessionAffinityAttributesOutputReference
	SessionAffinityTtl() *float64
	SteeringPolicy() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *float64
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
	ResetCountryPools()
	ResetPopPools()
	ResetRegionPools()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareLoadBalancersResultRulesOverridesOutputReference
type jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) AdaptiveRouting() DataCloudflareLoadBalancersResultRulesOverridesAdaptiveRoutingOutputReference {
	var returns DataCloudflareLoadBalancersResultRulesOverridesAdaptiveRoutingOutputReference
	_jsii_.Get(
		j,
		"adaptiveRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) CountryPools() any {
	var returns any
	_jsii_.Get(
		j,
		"countryPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) CountryPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"countryPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) DefaultPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) FallbackPool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fallbackPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) InternalValue() *DataCloudflareLoadBalancersResultRulesOverrides {
	var returns *DataCloudflareLoadBalancersResultRulesOverrides
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) LocationStrategy() DataCloudflareLoadBalancersResultRulesOverridesLocationStrategyOutputReference {
	var returns DataCloudflareLoadBalancersResultRulesOverridesLocationStrategyOutputReference
	_jsii_.Get(
		j,
		"locationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) PopPools() any {
	var returns any
	_jsii_.Get(
		j,
		"popPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) PopPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"popPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) RandomSteering() DataCloudflareLoadBalancersResultRulesOverridesRandomSteeringOutputReference {
	var returns DataCloudflareLoadBalancersResultRulesOverridesRandomSteeringOutputReference
	_jsii_.Get(
		j,
		"randomSteering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) RegionPools() any {
	var returns any
	_jsii_.Get(
		j,
		"regionPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) RegionPoolsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"regionPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SessionAffinityAttributes() DataCloudflareLoadBalancersResultRulesOverridesSessionAffinityAttributesOutputReference {
	var returns DataCloudflareLoadBalancersResultRulesOverridesSessionAffinityAttributesOutputReference
	_jsii_.Get(
		j,
		"sessionAffinityAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SessionAffinityTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionAffinityTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SteeringPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"steeringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func NewDataCloudflareLoadBalancersResultRulesOverridesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareLoadBalancersResultRulesOverridesOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareLoadBalancersResultRulesOverridesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareLoadBalancers.DataCloudflareLoadBalancersResultRulesOverridesOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareLoadBalancersResultRulesOverridesOutputReference_Override(d DataCloudflareLoadBalancersResultRulesOverridesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareLoadBalancers.DataCloudflareLoadBalancersResultRulesOverridesOutputReference",
		[]any{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetCountryPools(val any) {
	if err := j.validateSetCountryPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryPools",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetInternalValue(val *DataCloudflareLoadBalancersResultRulesOverrides) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetPopPools(val any) {
	if err := j.validateSetPopPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"popPools",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetRegionPools(val any) {
	if err := j.validateSetRegionPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionPools",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) ResetCountryPools() {
	_jsii_.InvokeVoid(
		d,
		"resetCountryPools",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) ResetPopPools() {
	_jsii_.InvokeVoid(
		d,
		"resetPopPools",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) ResetRegionPools() {
	_jsii_.InvokeVoid(
		d,
		"resetRegionPools",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareLoadBalancersResultRulesOverridesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
