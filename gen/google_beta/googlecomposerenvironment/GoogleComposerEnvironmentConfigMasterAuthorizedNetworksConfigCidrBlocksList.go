package googlecomposerenvironment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomposerenvironment/internal"
)

type GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList
type jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList {
	_init_.Initialize()

	if err := validateNewGoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList_Override(g GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComposerEnvironment.GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) Get(index *float64) GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComposerEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlocksList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

