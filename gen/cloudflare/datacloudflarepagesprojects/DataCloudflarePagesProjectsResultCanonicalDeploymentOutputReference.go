package datacloudflarepagesprojects

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflarepagesprojects/internal"
)

type DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference interface {
	cdktf.ComplexObject
	Aliases() *[]*string
	BuildConfig() DataCloudflarePagesProjectsResultCanonicalDeploymentBuildConfigOutputReference
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
	CreatedOn() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeploymentTrigger() DataCloudflarePagesProjectsResultCanonicalDeploymentDeploymentTriggerOutputReference
	Environment() *string
	EnvVars() DataCloudflarePagesProjectsResultCanonicalDeploymentEnvVarsMap
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() *DataCloudflarePagesProjectsResultCanonicalDeployment
	SetInternalValue(val *DataCloudflarePagesProjectsResultCanonicalDeployment)
	IsSkipped() cdktf.IResolvable
	LatestStage() DataCloudflarePagesProjectsResultCanonicalDeploymentLatestStageOutputReference
	ModifiedOn() *string
	ProjectId() *string
	ProjectName() *string
	ShortId() *string
	Source() DataCloudflarePagesProjectsResultCanonicalDeploymentSourceOutputReference
	Stages() DataCloudflarePagesProjectsResultCanonicalDeploymentStagesList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Url() *string
	UsesFunctions() cdktf.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference
type jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Aliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"aliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) BuildConfig() DataCloudflarePagesProjectsResultCanonicalDeploymentBuildConfigOutputReference {
	var returns DataCloudflarePagesProjectsResultCanonicalDeploymentBuildConfigOutputReference
	_jsii_.Get(
		j,
		"buildConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) CreatedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) DeploymentTrigger() DataCloudflarePagesProjectsResultCanonicalDeploymentDeploymentTriggerOutputReference {
	var returns DataCloudflarePagesProjectsResultCanonicalDeploymentDeploymentTriggerOutputReference
	_jsii_.Get(
		j,
		"deploymentTrigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) EnvVars() DataCloudflarePagesProjectsResultCanonicalDeploymentEnvVarsMap {
	var returns DataCloudflarePagesProjectsResultCanonicalDeploymentEnvVarsMap
	_jsii_.Get(
		j,
		"envVars",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) InternalValue() *DataCloudflarePagesProjectsResultCanonicalDeployment {
	var returns *DataCloudflarePagesProjectsResultCanonicalDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) IsSkipped() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isSkipped",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) LatestStage() DataCloudflarePagesProjectsResultCanonicalDeploymentLatestStageOutputReference {
	var returns DataCloudflarePagesProjectsResultCanonicalDeploymentLatestStageOutputReference
	_jsii_.Get(
		j,
		"latestStage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ModifiedOn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modifiedOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ProjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ShortId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shortId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Source() DataCloudflarePagesProjectsResultCanonicalDeploymentSourceOutputReference {
	var returns DataCloudflarePagesProjectsResultCanonicalDeploymentSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Stages() DataCloudflarePagesProjectsResultCanonicalDeploymentStagesList {
	var returns DataCloudflarePagesProjectsResultCanonicalDeploymentStagesList
	_jsii_.Get(
		j,
		"stages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) UsesFunctions() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"usesFunctions",
		&returns,
	)
	return returns
}


func NewDataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflarePagesProjectsResultCanonicalDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePagesProjects.DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference_Override(d DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflarePagesProjects.DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference)SetInternalValue(val *DataCloudflarePagesProjectsResultCanonicalDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflarePagesProjectsResultCanonicalDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

