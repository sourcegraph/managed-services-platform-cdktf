package googledataplexzone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledataplexzone/internal"
)

type GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference interface {
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
	Delimiter() *string
	SetDelimiter(val *string)
	DelimiterInput() *string
	DisableTypeInference() any
	SetDisableTypeInference(val any)
	DisableTypeInferenceInput() any
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	// Experimental.
	Fqn() *string
	HeaderRows() *float64
	SetHeaderRows(val *float64)
	HeaderRowsInput() *float64
	InternalValue() *GoogleDataplexZoneDiscoverySpecCsvOptions
	SetInternalValue(val *GoogleDataplexZoneDiscoverySpecCsvOptions)
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
	ResetDelimiter()
	ResetDisableTypeInference()
	ResetEncoding()
	ResetHeaderRows()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference
type jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) Delimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) DelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) DisableTypeInference() any {
	var returns any
	_jsii_.Get(
		j,
		"disableTypeInference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) DisableTypeInferenceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"disableTypeInferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) HeaderRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"headerRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) HeaderRowsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"headerRowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) InternalValue() *GoogleDataplexZoneDiscoverySpecCsvOptions {
	var returns *GoogleDataplexZoneDiscoverySpecCsvOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleDataplexZoneDiscoverySpecCsvOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataplexZone.GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference_Override(g GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataplexZone.GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetDelimiter(val *string) {
	if err := j.validateSetDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delimiter",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetDisableTypeInference(val any) {
	if err := j.validateSetDisableTypeInferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTypeInference",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetHeaderRows(val *float64) {
	if err := j.validateSetHeaderRowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerRows",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetInternalValue(val *GoogleDataplexZoneDiscoverySpecCsvOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ResetDelimiter() {
	_jsii_.InvokeVoid(
		g,
		"resetDelimiter",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ResetDisableTypeInference() {
	_jsii_.InvokeVoid(
		g,
		"resetDisableTypeInference",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ResetEncoding() {
	_jsii_.InvokeVoid(
		g,
		"resetEncoding",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ResetHeaderRows() {
	_jsii_.InvokeVoid(
		g,
		"resetHeaderRows",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleDataplexZoneDiscoverySpecCsvOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
