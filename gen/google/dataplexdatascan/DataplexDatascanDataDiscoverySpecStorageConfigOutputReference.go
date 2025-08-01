package dataplexdatascan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/dataplexdatascan/internal"
)

type DataplexDatascanDataDiscoverySpecStorageConfigOutputReference interface {
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
	CsvOptions() DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference
	CsvOptionsInput() *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions
	ExcludePatterns() *[]*string
	SetExcludePatterns(val *[]*string)
	ExcludePatternsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludePatterns() *[]*string
	SetIncludePatterns(val *[]*string)
	IncludePatternsInput() *[]*string
	InternalValue() *DataplexDatascanDataDiscoverySpecStorageConfig
	SetInternalValue(val *DataplexDatascanDataDiscoverySpecStorageConfig)
	JsonOptions() DataplexDatascanDataDiscoverySpecStorageConfigJsonOptionsOutputReference
	JsonOptionsInput() *DataplexDatascanDataDiscoverySpecStorageConfigJsonOptions
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
	PutCsvOptions(value *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions)
	PutJsonOptions(value *DataplexDatascanDataDiscoverySpecStorageConfigJsonOptions)
	ResetCsvOptions()
	ResetExcludePatterns()
	ResetIncludePatterns()
	ResetJsonOptions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataplexDatascanDataDiscoverySpecStorageConfigOutputReference
type jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) CsvOptions() DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference {
	var returns DataplexDatascanDataDiscoverySpecStorageConfigCsvOptionsOutputReference
	_jsii_.Get(
		j,
		"csvOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) CsvOptionsInput() *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions {
	var returns *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions
	_jsii_.Get(
		j,
		"csvOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ExcludePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ExcludePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) IncludePatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) IncludePatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includePatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) InternalValue() *DataplexDatascanDataDiscoverySpecStorageConfig {
	var returns *DataplexDatascanDataDiscoverySpecStorageConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) JsonOptions() DataplexDatascanDataDiscoverySpecStorageConfigJsonOptionsOutputReference {
	var returns DataplexDatascanDataDiscoverySpecStorageConfigJsonOptionsOutputReference
	_jsii_.Get(
		j,
		"jsonOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) JsonOptionsInput() *DataplexDatascanDataDiscoverySpecStorageConfigJsonOptions {
	var returns *DataplexDatascanDataDiscoverySpecStorageConfigJsonOptions
	_jsii_.Get(
		j,
		"jsonOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataplexDatascanDataDiscoverySpecStorageConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataplexDatascanDataDiscoverySpecStorageConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataplexDatascanDataDiscoverySpecStorageConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataDiscoverySpecStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataplexDatascanDataDiscoverySpecStorageConfigOutputReference_Override(d DataplexDatascanDataDiscoverySpecStorageConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.dataplexDatascan.DataplexDatascanDataDiscoverySpecStorageConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference)SetExcludePatterns(val *[]*string) {
	if err := j.validateSetExcludePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludePatterns",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference)SetIncludePatterns(val *[]*string) {
	if err := j.validateSetIncludePatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includePatterns",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference)SetInternalValue(val *DataplexDatascanDataDiscoverySpecStorageConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) PutCsvOptions(value *DataplexDatascanDataDiscoverySpecStorageConfigCsvOptions) {
	if err := d.validatePutCsvOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCsvOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) PutJsonOptions(value *DataplexDatascanDataDiscoverySpecStorageConfigJsonOptions) {
	if err := d.validatePutJsonOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJsonOptions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ResetCsvOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetCsvOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ResetExcludePatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludePatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ResetIncludePatterns() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludePatterns",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ResetJsonOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetJsonOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataplexDatascanDataDiscoverySpecStorageConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

