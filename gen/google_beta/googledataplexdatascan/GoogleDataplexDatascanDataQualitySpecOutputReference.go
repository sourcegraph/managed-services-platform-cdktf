package googledataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledataplexdatascan/internal"
)

type GoogleDataplexDatascanDataQualitySpecOutputReference interface {
	cdktf.ComplexObject
	CatalogPublishingEnabled() any
	SetCatalogPublishingEnabled(val any)
	CatalogPublishingEnabledInput() any
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
	InternalValue() *GoogleDataplexDatascanDataQualitySpec
	SetInternalValue(val *GoogleDataplexDatascanDataQualitySpec)
	PostScanActions() GoogleDataplexDatascanDataQualitySpecPostScanActionsOutputReference
	PostScanActionsInput() *GoogleDataplexDatascanDataQualitySpecPostScanActions
	RowFilter() *string
	SetRowFilter(val *string)
	RowFilterInput() *string
	Rules() GoogleDataplexDatascanDataQualitySpecRulesList
	RulesInput() any
	SamplingPercent() *float64
	SetSamplingPercent(val *float64)
	SamplingPercentInput() *float64
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
	PutPostScanActions(value *GoogleDataplexDatascanDataQualitySpecPostScanActions)
	PutRules(value any)
	ResetCatalogPublishingEnabled()
	ResetPostScanActions()
	ResetRowFilter()
	ResetRules()
	ResetSamplingPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexDatascanDataQualitySpecOutputReference
type jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) CatalogPublishingEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"catalogPublishingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) CatalogPublishingEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"catalogPublishingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) InternalValue() *GoogleDataplexDatascanDataQualitySpec {
	var returns *GoogleDataplexDatascanDataQualitySpec
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) PostScanActions() GoogleDataplexDatascanDataQualitySpecPostScanActionsOutputReference {
	var returns GoogleDataplexDatascanDataQualitySpecPostScanActionsOutputReference
	_jsii_.Get(
		j,
		"postScanActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) PostScanActionsInput() *GoogleDataplexDatascanDataQualitySpecPostScanActions {
	var returns *GoogleDataplexDatascanDataQualitySpecPostScanActions
	_jsii_.Get(
		j,
		"postScanActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) RowFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) RowFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) Rules() GoogleDataplexDatascanDataQualitySpecRulesList {
	var returns GoogleDataplexDatascanDataQualitySpecRulesList
	_jsii_.Get(
		j,
		"rules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) RulesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"rulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SamplingPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SamplingPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"samplingPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleDataplexDatascanDataQualitySpecOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexDatascanDataQualitySpecOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexDatascanDataQualitySpecOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexDatascanDataQualitySpecOutputReference_Override(g GoogleDataplexDatascanDataQualitySpecOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataplexDatascan.GoogleDataplexDatascanDataQualitySpecOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetCatalogPublishingEnabled(val any) {
	if err := j.validateSetCatalogPublishingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogPublishingEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetInternalValue(val *GoogleDataplexDatascanDataQualitySpec) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetRowFilter(val *string) {
	if err := j.validateSetRowFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowFilter",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetSamplingPercent(val *float64) {
	if err := j.validateSetSamplingPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplingPercent",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) PutPostScanActions(value *GoogleDataplexDatascanDataQualitySpecPostScanActions) {
	if err := g.validatePutPostScanActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostScanActions",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) PutRules(value any) {
	if err := g.validatePutRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRules",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ResetCatalogPublishingEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogPublishingEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ResetPostScanActions() {
	_jsii_.InvokeVoid(
		g,
		"resetPostScanActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ResetRowFilter() {
	_jsii_.InvokeVoid(
		g,
		"resetRowFilter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ResetRules() {
	_jsii_.InvokeVoid(
		g,
		"resetRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ResetSamplingPercent() {
	_jsii_.InvokeVoid(
		g,
		"resetSamplingPercent",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleDataplexDatascanDataQualitySpecOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
