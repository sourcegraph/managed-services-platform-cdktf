package osconfigv2policyorchestratorfororganization

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/osconfigv2policyorchestratorfororganization/internal"
)

type OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList
type jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList {
	_init_.Initialize()

	if err := validateNewOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList{}

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestratorForOrganization.OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList_Override(o OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.osConfigV2PolicyOrchestratorForOrganization.OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) Get(index *float64) OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsConfigV2PolicyOrchestratorForOrganizationOrchestrationStateCurrentIterationStateErrorDetailsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

