package integrationaction

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/opsgenie/integrationaction/internal"
)

type IntegrationActionCreateOutputReference interface {
	cdktf.ComplexObject
	AlertActions() *[]*string
	SetAlertActions(val *[]*string)
	AlertActionsInput() *[]*string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AppendAttachments() interface{}
	SetAppendAttachments(val interface{})
	AppendAttachmentsInput() interface{}
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
	CustomPriority() *string
	SetCustomPriority(val *string)
	CustomPriorityInput() *string
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Entity() *string
	SetEntity(val *string)
	EntityInput() *string
	ExtraProperties() *map[string]*string
	SetExtraProperties(val *map[string]*string)
	ExtraPropertiesInput() *map[string]*string
	Filter() IntegrationActionCreateFilterList
	FilterInput() interface{}
	// Experimental.
	Fqn() *string
	IgnoreAlertActionsFromPayload() interface{}
	SetIgnoreAlertActionsFromPayload(val interface{})
	IgnoreAlertActionsFromPayloadInput() interface{}
	IgnoreExtraPropertiesFromPayload() interface{}
	SetIgnoreExtraPropertiesFromPayload(val interface{})
	IgnoreExtraPropertiesFromPayloadInput() interface{}
	IgnoreRespondersFromPayload() interface{}
	SetIgnoreRespondersFromPayload(val interface{})
	IgnoreRespondersFromPayloadInput() interface{}
	IgnoreTagsFromPayload() interface{}
	SetIgnoreTagsFromPayload(val interface{})
	IgnoreTagsFromPayloadInput() interface{}
	IgnoreTeamsFromPayload() interface{}
	SetIgnoreTeamsFromPayload(val interface{})
	IgnoreTeamsFromPayloadInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Message() *string
	SetMessage(val *string)
	MessageInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Note() *string
	SetNote(val *string)
	NoteInput() *string
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	Responders() IntegrationActionCreateRespondersList
	RespondersInput() interface{}
	Source() *string
	SetSource(val *string)
	SourceInput() *string
	Tags() *[]*string
	SetTags(val *[]*string)
	TagsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
	User() *string
	SetUser(val *string)
	UserInput() *string
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
	PutFilter(value interface{})
	PutResponders(value interface{})
	ResetAlertActions()
	ResetAlias()
	ResetAppendAttachments()
	ResetCustomPriority()
	ResetDescription()
	ResetEntity()
	ResetExtraProperties()
	ResetFilter()
	ResetIgnoreAlertActionsFromPayload()
	ResetIgnoreExtraPropertiesFromPayload()
	ResetIgnoreRespondersFromPayload()
	ResetIgnoreTagsFromPayload()
	ResetIgnoreTeamsFromPayload()
	ResetMessage()
	ResetNote()
	ResetOrder()
	ResetPriority()
	ResetResponders()
	ResetSource()
	ResetTags()
	ResetType()
	ResetUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegrationActionCreateOutputReference
type jsiiProxy_IntegrationActionCreateOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) AlertActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) AlertActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"alertActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) AppendAttachments() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) AppendAttachmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"appendAttachmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) CustomPriority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) CustomPriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Entity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) EntityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) ExtraProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) ExtraPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Filter() IntegrationActionCreateFilterList {
	var returns IntegrationActionCreateFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreAlertActionsFromPayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreAlertActionsFromPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreAlertActionsFromPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreAlertActionsFromPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreExtraPropertiesFromPayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreExtraPropertiesFromPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreExtraPropertiesFromPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreExtraPropertiesFromPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreRespondersFromPayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreRespondersFromPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreRespondersFromPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreRespondersFromPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreTagsFromPayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTagsFromPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreTagsFromPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTagsFromPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreTeamsFromPayload() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTeamsFromPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) IgnoreTeamsFromPayloadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTeamsFromPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Message() *string {
	var returns *string
	_jsii_.Get(
		j,
		"message",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) MessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Note() *string {
	var returns *string
	_jsii_.Get(
		j,
		"note",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) NoteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Responders() IntegrationActionCreateRespondersList {
	var returns IntegrationActionCreateRespondersList
	_jsii_.Get(
		j,
		"responders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) RespondersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"respondersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Tags() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) TagsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) User() *string {
	var returns *string
	_jsii_.Get(
		j,
		"user",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference) UserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userInput",
		&returns,
	)
	return returns
}


func NewIntegrationActionCreateOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) IntegrationActionCreateOutputReference {
	_init_.Initialize()

	if err := validateNewIntegrationActionCreateOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegrationActionCreateOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-opsgenie.integrationAction.IntegrationActionCreateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewIntegrationActionCreateOutputReference_Override(i IntegrationActionCreateOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-opsgenie.integrationAction.IntegrationActionCreateOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		i,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetAlertActions(val *[]*string) {
	if err := j.validateSetAlertActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alertActions",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetAlias(val *string) {
	if err := j.validateSetAliasParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetAppendAttachments(val interface{}) {
	if err := j.validateSetAppendAttachmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appendAttachments",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetCustomPriority(val *string) {
	if err := j.validateSetCustomPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customPriority",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetEntity(val *string) {
	if err := j.validateSetEntityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entity",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetExtraProperties(val *map[string]*string) {
	if err := j.validateSetExtraPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraProperties",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetIgnoreAlertActionsFromPayload(val interface{}) {
	if err := j.validateSetIgnoreAlertActionsFromPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreAlertActionsFromPayload",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetIgnoreExtraPropertiesFromPayload(val interface{}) {
	if err := j.validateSetIgnoreExtraPropertiesFromPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreExtraPropertiesFromPayload",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetIgnoreRespondersFromPayload(val interface{}) {
	if err := j.validateSetIgnoreRespondersFromPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreRespondersFromPayload",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetIgnoreTagsFromPayload(val interface{}) {
	if err := j.validateSetIgnoreTagsFromPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTagsFromPayload",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetIgnoreTeamsFromPayload(val interface{}) {
	if err := j.validateSetIgnoreTeamsFromPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTeamsFromPayload",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetMessage(val *string) {
	if err := j.validateSetMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"message",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetNote(val *string) {
	if err := j.validateSetNoteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"note",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetTags(val *[]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_IntegrationActionCreateOutputReference)SetUser(val *string) {
	if err := j.validateSetUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"user",
		val,
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) PutFilter(value interface{}) {
	if err := i.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putFilter",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) PutResponders(value interface{}) {
	if err := i.validatePutRespondersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putResponders",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetAlertActions() {
	_jsii_.InvokeVoid(
		i,
		"resetAlertActions",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetAlias() {
	_jsii_.InvokeVoid(
		i,
		"resetAlias",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetAppendAttachments() {
	_jsii_.InvokeVoid(
		i,
		"resetAppendAttachments",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetCustomPriority() {
	_jsii_.InvokeVoid(
		i,
		"resetCustomPriority",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetEntity() {
	_jsii_.InvokeVoid(
		i,
		"resetEntity",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetExtraProperties() {
	_jsii_.InvokeVoid(
		i,
		"resetExtraProperties",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		i,
		"resetFilter",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetIgnoreAlertActionsFromPayload() {
	_jsii_.InvokeVoid(
		i,
		"resetIgnoreAlertActionsFromPayload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetIgnoreExtraPropertiesFromPayload() {
	_jsii_.InvokeVoid(
		i,
		"resetIgnoreExtraPropertiesFromPayload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetIgnoreRespondersFromPayload() {
	_jsii_.InvokeVoid(
		i,
		"resetIgnoreRespondersFromPayload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetIgnoreTagsFromPayload() {
	_jsii_.InvokeVoid(
		i,
		"resetIgnoreTagsFromPayload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetIgnoreTeamsFromPayload() {
	_jsii_.InvokeVoid(
		i,
		"resetIgnoreTeamsFromPayload",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetMessage() {
	_jsii_.InvokeVoid(
		i,
		"resetMessage",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetNote() {
	_jsii_.InvokeVoid(
		i,
		"resetNote",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		i,
		"resetOrder",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		i,
		"resetPriority",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetResponders() {
	_jsii_.InvokeVoid(
		i,
		"resetResponders",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetSource() {
	_jsii_.InvokeVoid(
		i,
		"resetSource",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		i,
		"resetType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ResetUser() {
	_jsii_.InvokeVoid(
		i,
		"resetUser",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegrationActionCreateOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

