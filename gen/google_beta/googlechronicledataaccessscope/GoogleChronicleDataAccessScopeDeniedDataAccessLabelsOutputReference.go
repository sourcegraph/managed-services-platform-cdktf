package googlechronicledataaccessscope

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlechronicledataaccessscope/internal"
)

type GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference interface {
	cdktf.ComplexObject
	AssetNamespace() *string
	SetAssetNamespace(val *string)
	AssetNamespaceInput() *string
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
	DataAccessLabel() *string
	SetDataAccessLabel(val *string)
	DataAccessLabelInput() *string
	DisplayName() *string
	// Experimental.
	Fqn() *string
	IngestionLabel() GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference
	IngestionLabelInput() *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LogType() *string
	SetLogType(val *string)
	LogTypeInput() *string
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
	PutIngestionLabel(value *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel)
	ResetAssetNamespace()
	ResetDataAccessLabel()
	ResetIngestionLabel()
	ResetLogType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference
type jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) AssetNamespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) AssetNamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetNamespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) DataAccessLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) DataAccessLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataAccessLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) IngestionLabel() GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference {
	var returns GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabelOutputReference
	_jsii_.Get(
		j,
		"ingestionLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) IngestionLabelInput() *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel {
	var returns *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel
	_jsii_.Get(
		j,
		"ingestionLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) LogType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) LogTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleChronicleDataAccessScope.GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewGoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference_Override(g GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleChronicleDataAccessScope.GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		g,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetAssetNamespace(val *string) {
	if err := j.validateSetAssetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetNamespace",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetDataAccessLabel(val *string) {
	if err := j.validateSetDataAccessLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataAccessLabel",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetLogType(val *string) {
	if err := j.validateSetLogTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logType",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) PutIngestionLabel(value *GoogleChronicleDataAccessScopeDeniedDataAccessLabelsIngestionLabel) {
	if err := g.validatePutIngestionLabelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIngestionLabel",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ResetAssetNamespace() {
	_jsii_.InvokeVoid(
		g,
		"resetAssetNamespace",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ResetDataAccessLabel() {
	_jsii_.InvokeVoid(
		g,
		"resetDataAccessLabel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ResetIngestionLabel() {
	_jsii_.InvokeVoid(
		g,
		"resetIngestionLabel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ResetLogType() {
	_jsii_.InvokeVoid(
		g,
		"resetLogType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleChronicleDataAccessScopeDeniedDataAccessLabelsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

