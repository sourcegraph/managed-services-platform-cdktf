package googledataprocworkflowtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googledataprocworkflowtemplate/internal"
)

type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList
type jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList {
	_init_.Initialize()

	if err := validateNewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewGoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList_Override(g GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		g,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) Get(index *float64) GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigOutputReference {
	if err := g.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigOutputReference

	_jsii_.Invoke(
		g,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

