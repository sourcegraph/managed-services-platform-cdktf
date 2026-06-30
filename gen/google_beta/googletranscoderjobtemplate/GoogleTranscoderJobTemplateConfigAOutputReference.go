package googletranscoderjobtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googletranscoderjobtemplate/internal"
)

type GoogleTranscoderJobTemplateConfigAOutputReference interface {
	cdktf.ComplexObject
	AdBreaks() GoogleTranscoderJobTemplateConfigAdBreaksList
	AdBreaksInput() any
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
	EditList() GoogleTranscoderJobTemplateConfigEditListStructList
	EditListInput() any
	ElementaryStreams() GoogleTranscoderJobTemplateConfigElementaryStreamsList
	ElementaryStreamsInput() any
	Encryptions() GoogleTranscoderJobTemplateConfigEncryptionsList
	EncryptionsInput() any
	// Experimental.
	Fqn() *string
	Inputs() GoogleTranscoderJobTemplateConfigInputsList
	InputsInput() any
	InternalValue() *GoogleTranscoderJobTemplateConfigA
	SetInternalValue(val *GoogleTranscoderJobTemplateConfigA)
	Manifests() GoogleTranscoderJobTemplateConfigManifestsList
	ManifestsInput() any
	MuxStreams() GoogleTranscoderJobTemplateConfigMuxStreamsList
	MuxStreamsInput() any
	Output() GoogleTranscoderJobTemplateConfigOutputOutputReference
	OutputInput() *GoogleTranscoderJobTemplateConfigOutput
	Overlays() GoogleTranscoderJobTemplateConfigOverlaysList
	OverlaysInput() any
	PubsubDestination() GoogleTranscoderJobTemplateConfigPubsubDestinationOutputReference
	PubsubDestinationInput() *GoogleTranscoderJobTemplateConfigPubsubDestination
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
	PutAdBreaks(value any)
	PutEditList(value any)
	PutElementaryStreams(value any)
	PutEncryptions(value any)
	PutInputs(value any)
	PutManifests(value any)
	PutMuxStreams(value any)
	PutOutput(value *GoogleTranscoderJobTemplateConfigOutput)
	PutOverlays(value any)
	PutPubsubDestination(value *GoogleTranscoderJobTemplateConfigPubsubDestination)
	ResetAdBreaks()
	ResetEditList()
	ResetElementaryStreams()
	ResetEncryptions()
	ResetInputs()
	ResetManifests()
	ResetMuxStreams()
	ResetOutput()
	ResetOverlays()
	ResetPubsubDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleTranscoderJobTemplateConfigAOutputReference
type jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) AdBreaks() GoogleTranscoderJobTemplateConfigAdBreaksList {
	var returns GoogleTranscoderJobTemplateConfigAdBreaksList
	_jsii_.Get(
		j,
		"adBreaks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) AdBreaksInput() any {
	var returns any
	_jsii_.Get(
		j,
		"adBreaksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) EditList() GoogleTranscoderJobTemplateConfigEditListStructList {
	var returns GoogleTranscoderJobTemplateConfigEditListStructList
	_jsii_.Get(
		j,
		"editList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) EditListInput() any {
	var returns any
	_jsii_.Get(
		j,
		"editListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ElementaryStreams() GoogleTranscoderJobTemplateConfigElementaryStreamsList {
	var returns GoogleTranscoderJobTemplateConfigElementaryStreamsList
	_jsii_.Get(
		j,
		"elementaryStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ElementaryStreamsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"elementaryStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) Encryptions() GoogleTranscoderJobTemplateConfigEncryptionsList {
	var returns GoogleTranscoderJobTemplateConfigEncryptionsList
	_jsii_.Get(
		j,
		"encryptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) EncryptionsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"encryptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) Inputs() GoogleTranscoderJobTemplateConfigInputsList {
	var returns GoogleTranscoderJobTemplateConfigInputsList
	_jsii_.Get(
		j,
		"inputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) InputsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"inputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) InternalValue() *GoogleTranscoderJobTemplateConfigA {
	var returns *GoogleTranscoderJobTemplateConfigA
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) Manifests() GoogleTranscoderJobTemplateConfigManifestsList {
	var returns GoogleTranscoderJobTemplateConfigManifestsList
	_jsii_.Get(
		j,
		"manifests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ManifestsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"manifestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) MuxStreams() GoogleTranscoderJobTemplateConfigMuxStreamsList {
	var returns GoogleTranscoderJobTemplateConfigMuxStreamsList
	_jsii_.Get(
		j,
		"muxStreams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) MuxStreamsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"muxStreamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) Output() GoogleTranscoderJobTemplateConfigOutputOutputReference {
	var returns GoogleTranscoderJobTemplateConfigOutputOutputReference
	_jsii_.Get(
		j,
		"output",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) OutputInput() *GoogleTranscoderJobTemplateConfigOutput {
	var returns *GoogleTranscoderJobTemplateConfigOutput
	_jsii_.Get(
		j,
		"outputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) Overlays() GoogleTranscoderJobTemplateConfigOverlaysList {
	var returns GoogleTranscoderJobTemplateConfigOverlaysList
	_jsii_.Get(
		j,
		"overlays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) OverlaysInput() any {
	var returns any
	_jsii_.Get(
		j,
		"overlaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PubsubDestination() GoogleTranscoderJobTemplateConfigPubsubDestinationOutputReference {
	var returns GoogleTranscoderJobTemplateConfigPubsubDestinationOutputReference
	_jsii_.Get(
		j,
		"pubsubDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PubsubDestinationInput() *GoogleTranscoderJobTemplateConfigPubsubDestination {
	var returns *GoogleTranscoderJobTemplateConfigPubsubDestination
	_jsii_.Get(
		j,
		"pubsubDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleTranscoderJobTemplateConfigAOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleTranscoderJobTemplateConfigAOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleTranscoderJobTemplateConfigAOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigAOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleTranscoderJobTemplateConfigAOutputReference_Override(g GoogleTranscoderJobTemplateConfigAOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleTranscoderJobTemplate.GoogleTranscoderJobTemplateConfigAOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) SetInternalValue(val *GoogleTranscoderJobTemplateConfigA) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutAdBreaks(value any) {
	if err := g.validatePutAdBreaksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAdBreaks",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutEditList(value any) {
	if err := g.validatePutEditListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEditList",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutElementaryStreams(value any) {
	if err := g.validatePutElementaryStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putElementaryStreams",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutEncryptions(value any) {
	if err := g.validatePutEncryptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptions",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutInputs(value any) {
	if err := g.validatePutInputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putInputs",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutManifests(value any) {
	if err := g.validatePutManifestsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManifests",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutMuxStreams(value any) {
	if err := g.validatePutMuxStreamsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMuxStreams",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutOutput(value *GoogleTranscoderJobTemplateConfigOutput) {
	if err := g.validatePutOutputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutput",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutOverlays(value any) {
	if err := g.validatePutOverlaysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOverlays",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) PutPubsubDestination(value *GoogleTranscoderJobTemplateConfigPubsubDestination) {
	if err := g.validatePutPubsubDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPubsubDestination",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetAdBreaks() {
	_jsii_.InvokeVoid(
		g,
		"resetAdBreaks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetEditList() {
	_jsii_.InvokeVoid(
		g,
		"resetEditList",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetElementaryStreams() {
	_jsii_.InvokeVoid(
		g,
		"resetElementaryStreams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetEncryptions() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetInputs() {
	_jsii_.InvokeVoid(
		g,
		"resetInputs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetManifests() {
	_jsii_.InvokeVoid(
		g,
		"resetManifests",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetMuxStreams() {
	_jsii_.InvokeVoid(
		g,
		"resetMuxStreams",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetOutput() {
	_jsii_.InvokeVoid(
		g,
		"resetOutput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetOverlays() {
	_jsii_.InvokeVoid(
		g,
		"resetOverlays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ResetPubsubDestination() {
	_jsii_.InvokeVoid(
		g,
		"resetPubsubDestination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleTranscoderJobTemplateConfigAOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
