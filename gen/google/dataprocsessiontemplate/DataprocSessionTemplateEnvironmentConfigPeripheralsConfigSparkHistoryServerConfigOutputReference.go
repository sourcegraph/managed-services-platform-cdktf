package dataprocsessiontemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/dataprocsessiontemplate/internal"
)

type DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataprocCluster() *string
	SetDataprocCluster(val *string)
	DataprocClusterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig
	SetInternalValue(val *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig)
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
	ResetDataprocCluster()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference
type jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) DataprocCluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocCluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) DataprocClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataprocClusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) InternalValue() *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig {
	var returns *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataprocSessionTemplate.DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference_Override(d DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataprocSessionTemplate.DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference)SetDataprocCluster(val *string) {
	if err := j.validateSetDataprocClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataprocCluster",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference)SetInternalValue(val *DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) ResetDataprocCluster() {
	_jsii_.InvokeVoid(
		d,
		"resetDataprocCluster",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataprocSessionTemplateEnvironmentConfigPeripheralsConfigSparkHistoryServerConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

