package datatfeteamprojectaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/tfe/datatfeteamprojectaccess/internal"
)

type DataTfeTeamProjectAccessWorkspaceAccessOutputReference interface {
	cdktf.ComplexObject
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
	Create() cdktf.IResolvable
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *DataTfeTeamProjectAccessWorkspaceAccess
	SetInternalValue(val *DataTfeTeamProjectAccessWorkspaceAccess)
	Locking() cdktf.IResolvable
	Move() cdktf.IResolvable
	Runs() *string
	RunTasks() cdktf.IResolvable
	SentinelMocks() *string
	StateVersions() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Variables() *string
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

// The jsii proxy struct for DataTfeTeamProjectAccessWorkspaceAccessOutputReference
type jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Create() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Delete() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) InternalValue() *DataTfeTeamProjectAccessWorkspaceAccess {
	var returns *DataTfeTeamProjectAccessWorkspaceAccess
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Locking() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"locking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Move() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"move",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Runs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) RunTasks() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"runTasks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) SentinelMocks() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sentinelMocks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) StateVersions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Variables() *string {
	var returns *string
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewDataTfeTeamProjectAccessWorkspaceAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfeTeamProjectAccessWorkspaceAccessOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfeTeamProjectAccessWorkspaceAccessOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-tfe.dataTfeTeamProjectAccess.DataTfeTeamProjectAccessWorkspaceAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataTfeTeamProjectAccessWorkspaceAccessOutputReference_Override(d DataTfeTeamProjectAccessWorkspaceAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-tfe.dataTfeTeamProjectAccess.DataTfeTeamProjectAccessWorkspaceAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference)SetInternalValue(val *DataTfeTeamProjectAccessWorkspaceAccess) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataTfeTeamProjectAccessWorkspaceAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

