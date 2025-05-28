package projects_iam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/projects_iam/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/projects_iam/internal"
)

// Defines an ProjectsIam based on a Terraform module.
//
// Docs at Terraform Registry: {@link https://registry.terraform.io/modules/terraform-google-modules/iam/google/8.1.0/submodules/projects_iam terraform-google-modules/iam/google//modules/projects_iam}
type ProjectsIam interface {
	cdktf.TerraformModule
	Bindings() *map[string]*[]*string
	SetBindings(val *map[string]*[]*string)
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConditionalBindings() interface{}
	SetConditionalBindings(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	MembersOutput() *string
	Mode() *string
	SetMode(val *string)
	// The tree node.
	Node() constructs.Node
	Projects() *[]*string
	SetProjects(val *[]*string)
	ProjectsOutput() *string
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	RolesOutput() *string
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	// Experimental.
	Version() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ProjectsIam
type jsiiProxy_ProjectsIam struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_ProjectsIam) Bindings() *map[string]*[]*string {
	var returns *map[string]*[]*string
	_jsii_.Get(
		j,
		"bindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) ConditionalBindings() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionalBindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) MembersOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"membersOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) Projects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"projects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) ProjectsOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectsOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) RolesOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rolesOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ProjectsIam) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


func NewProjectsIam(scope constructs.Construct, id *string, config *ProjectsIamConfig) ProjectsIam {
	_init_.Initialize()

	if err := validateNewProjectsIamParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ProjectsIam{}

	_jsii_.Create(
		"@cdktf/provider-projects_iam.ProjectsIam",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

func NewProjectsIam_Override(p ProjectsIam, scope constructs.Construct, id *string, config *ProjectsIamConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-projects_iam.ProjectsIam",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_ProjectsIam)SetBindings(val *map[string]*[]*string) {
	_jsii_.Set(
		j,
		"bindings",
		val,
	)
}

func (j *jsiiProxy_ProjectsIam)SetConditionalBindings(val interface{}) {
	if err := j.validateSetConditionalBindingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"conditionalBindings",
		val,
	)
}

func (j *jsiiProxy_ProjectsIam)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ProjectsIam)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ProjectsIam)SetMode(val *string) {
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ProjectsIam)SetProjects(val *[]*string) {
	_jsii_.Set(
		j,
		"projects",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ProjectsIam_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectsIam_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-projects_iam.ProjectsIam",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ProjectsIam_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProjectsIam_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-projects_iam.ProjectsIam",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_ProjectsIam) AddProvider(provider interface{}) {
	if err := p.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addProvider",
		[]interface{}{provider},
	)
}

func (p *jsiiProxy_ProjectsIam) GetString(output *string) *string {
	if err := p.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := p.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_ProjectsIam) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_ProjectsIam) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_ProjectsIam) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

