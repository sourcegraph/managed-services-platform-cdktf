package dialogflowcxwebhook

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/dialogflowcxwebhook/internal"
)

type DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList
type jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList {
	_init_.Initialize()

	if err := validateNewDialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList{}

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxWebhook.DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList_Override(d DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dialogflowCxWebhook.DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := d.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		d,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) Get(index *float64) DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DialogflowCxWebhookServiceDirectoryGenericWebServiceSecretVersionsForRequestHeadersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

