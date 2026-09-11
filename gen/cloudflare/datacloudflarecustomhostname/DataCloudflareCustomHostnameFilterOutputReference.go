package datacloudflarecustomhostname

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflarecustomhostname/internal"
)

type DataCloudflareCustomHostnameFilterOutputReference interface {
	cdktf.ComplexObject
	CertificateAuthority() *string
	SetCertificateAuthority(val *string)
	CertificateAuthorityInput() *string
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
	CustomOriginServer() *string
	SetCustomOriginServer(val *string)
	CustomOriginServerInput() *string
	Direction() *string
	SetDirection(val *string)
	DirectionInput() *string
	// Experimental.
	Fqn() *string
	Hostname() DataCloudflareCustomHostnameFilterHostnameOutputReference
	HostnameInput() interface{}
	HostnameStatus() *string
	SetHostnameStatus(val *string)
	HostnameStatusInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Order() *string
	SetOrder(val *string)
	OrderInput() *string
	Ssl() *float64
	SetSsl(val *float64)
	SslInput() *float64
	SslStatus() *string
	SetSslStatus(val *string)
	SslStatusInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Wildcard() interface{}
	SetWildcard(val interface{})
	WildcardInput() interface{}
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
	PutHostname(value *DataCloudflareCustomHostnameFilterHostname)
	ResetCertificateAuthority()
	ResetCustomOriginServer()
	ResetDirection()
	ResetHostname()
	ResetHostnameStatus()
	ResetId()
	ResetOrder()
	ResetSsl()
	ResetSslStatus()
	ResetWildcard()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataCloudflareCustomHostnameFilterOutputReference
type jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) CertificateAuthority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) CertificateAuthorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) CustomOriginServer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customOriginServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) CustomOriginServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customOriginServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Direction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"direction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) DirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"directionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Hostname() DataCloudflareCustomHostnameFilterHostnameOutputReference {
	var returns DataCloudflareCustomHostnameFilterHostnameOutputReference
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) HostnameInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) HostnameStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) HostnameStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Order() *string {
	var returns *string
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) OrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Ssl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ssl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) SslInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) SslStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) SslStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Wildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) WildcardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcardInput",
		&returns,
	)
	return returns
}


func NewDataCloudflareCustomHostnameFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareCustomHostnameFilterOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareCustomHostnameFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareCustomHostname.DataCloudflareCustomHostnameFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareCustomHostnameFilterOutputReference_Override(d DataCloudflareCustomHostnameFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareCustomHostname.DataCloudflareCustomHostnameFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetCertificateAuthority(val *string) {
	if err := j.validateSetCertificateAuthorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateAuthority",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetCustomOriginServer(val *string) {
	if err := j.validateSetCustomOriginServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customOriginServer",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetDirection(val *string) {
	if err := j.validateSetDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"direction",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetHostnameStatus(val *string) {
	if err := j.validateSetHostnameStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameStatus",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetOrder(val *string) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetSsl(val *float64) {
	if err := j.validateSetSslParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ssl",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetSslStatus(val *string) {
	if err := j.validateSetSslStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslStatus",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference)SetWildcard(val interface{}) {
	if err := j.validateSetWildcardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildcard",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) PutHostname(value *DataCloudflareCustomHostnameFilterHostname) {
	if err := d.validatePutHostnameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHostname",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetCertificateAuthority() {
	_jsii_.InvokeVoid(
		d,
		"resetCertificateAuthority",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetCustomOriginServer() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomOriginServer",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		d,
		"resetHostname",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetHostnameStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetHostnameStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		d,
		"resetOrder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetSsl() {
	_jsii_.InvokeVoid(
		d,
		"resetSsl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetSslStatus() {
	_jsii_.InvokeVoid(
		d,
		"resetSslStatus",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ResetWildcard() {
	_jsii_.InvokeVoid(
		d,
		"resetWildcard",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareCustomHostnameFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

