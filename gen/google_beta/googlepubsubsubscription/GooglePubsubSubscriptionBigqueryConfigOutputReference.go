package googlepubsubsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlepubsubsubscription/internal"
)

type GooglePubsubSubscriptionBigqueryConfigOutputReference interface {
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
	DropUnknownFields() any
	SetDropUnknownFields(val any)
	DropUnknownFieldsInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *GooglePubsubSubscriptionBigqueryConfig
	SetInternalValue(val *GooglePubsubSubscriptionBigqueryConfig)
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseTableSchema() any
	SetUseTableSchema(val any)
	UseTableSchemaInput() any
	UseTopicSchema() any
	SetUseTopicSchema(val any)
	UseTopicSchemaInput() any
	WriteMetadata() any
	SetWriteMetadata(val any)
	WriteMetadataInput() any
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
	ResetDropUnknownFields()
	ResetServiceAccountEmail()
	ResetUseTableSchema()
	ResetUseTopicSchema()
	ResetWriteMetadata()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GooglePubsubSubscriptionBigqueryConfigOutputReference
type jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) DropUnknownFields() any {
	var returns any
	_jsii_.Get(
		j,
		"dropUnknownFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) DropUnknownFieldsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dropUnknownFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) InternalValue() *GooglePubsubSubscriptionBigqueryConfig {
	var returns *GooglePubsubSubscriptionBigqueryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) UseTableSchema() any {
	var returns any
	_jsii_.Get(
		j,
		"useTableSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) UseTableSchemaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useTableSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) UseTopicSchema() any {
	var returns any
	_jsii_.Get(
		j,
		"useTopicSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) UseTopicSchemaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useTopicSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) WriteMetadata() any {
	var returns any
	_jsii_.Get(
		j,
		"writeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) WriteMetadataInput() any {
	var returns any
	_jsii_.Get(
		j,
		"writeMetadataInput",
		&returns,
	)
	return returns
}

func NewGooglePubsubSubscriptionBigqueryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GooglePubsubSubscriptionBigqueryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGooglePubsubSubscriptionBigqueryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePubsubSubscription.GooglePubsubSubscriptionBigqueryConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGooglePubsubSubscriptionBigqueryConfigOutputReference_Override(g GooglePubsubSubscriptionBigqueryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googlePubsubSubscription.GooglePubsubSubscriptionBigqueryConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetDropUnknownFields(val any) {
	if err := j.validateSetDropUnknownFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropUnknownFields",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetInternalValue(val *GooglePubsubSubscriptionBigqueryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetUseTableSchema(val any) {
	if err := j.validateSetUseTableSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTableSchema",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetUseTopicSchema(val any) {
	if err := j.validateSetUseTopicSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTopicSchema",
		val,
	)
}

func (j *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) SetWriteMetadata(val any) {
	if err := j.validateSetWriteMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeMetadata",
		val,
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ResetDropUnknownFields() {
	_jsii_.InvokeVoid(
		g,
		"resetDropUnknownFields",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ResetUseTableSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetUseTableSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ResetUseTopicSchema() {
	_jsii_.InvokeVoid(
		g,
		"resetUseTopicSchema",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ResetWriteMetadata() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteMetadata",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GooglePubsubSubscriptionBigqueryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
