package datagooglecloudquotasquotainfos

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/datagooglecloudquotasquotainfos/internal"
)

type DataGoogleCloudQuotasQuotaInfosQuotaInfosList interface {
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
	Get(index *float64) DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGoogleCloudQuotasQuotaInfosQuotaInfosList
type jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataGoogleCloudQuotasQuotaInfosQuotaInfosList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataGoogleCloudQuotasQuotaInfosQuotaInfosList {
	_init_.Initialize()

	if err := validateNewDataGoogleCloudQuotasQuotaInfosQuotaInfosListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.dataGoogleCloudQuotasQuotaInfos.DataGoogleCloudQuotasQuotaInfosQuotaInfosList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataGoogleCloudQuotasQuotaInfosQuotaInfosList_Override(d DataGoogleCloudQuotasQuotaInfosQuotaInfosList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.dataGoogleCloudQuotasQuotaInfos.DataGoogleCloudQuotasQuotaInfosQuotaInfosList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) Get(index *float64) DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataGoogleCloudQuotasQuotaInfosQuotaInfosOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataGoogleCloudQuotasQuotaInfosQuotaInfosList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
