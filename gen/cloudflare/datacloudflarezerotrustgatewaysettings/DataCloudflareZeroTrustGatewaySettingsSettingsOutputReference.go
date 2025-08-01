package datacloudflarezerotrustgatewaysettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/cloudflare/datacloudflarezerotrustgatewaysettings/internal"
)

type DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference interface {
	cdktf.ComplexObject
	ActivityLog() DataCloudflareZeroTrustGatewaySettingsSettingsActivityLogOutputReference
	Antivirus() DataCloudflareZeroTrustGatewaySettingsSettingsAntivirusOutputReference
	BlockPage() DataCloudflareZeroTrustGatewaySettingsSettingsBlockPageOutputReference
	BodyScanning() DataCloudflareZeroTrustGatewaySettingsSettingsBodyScanningOutputReference
	BrowserIsolation() DataCloudflareZeroTrustGatewaySettingsSettingsBrowserIsolationOutputReference
	Certificate() DataCloudflareZeroTrustGatewaySettingsSettingsCertificateOutputReference
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
	CustomCertificate() DataCloudflareZeroTrustGatewaySettingsSettingsCustomCertificateOutputReference
	ExtendedEmailMatching() DataCloudflareZeroTrustGatewaySettingsSettingsExtendedEmailMatchingOutputReference
	Fips() DataCloudflareZeroTrustGatewaySettingsSettingsFipsOutputReference
	// Experimental.
	Fqn() *string
	HostSelector() DataCloudflareZeroTrustGatewaySettingsSettingsHostSelectorOutputReference
	InternalValue() *DataCloudflareZeroTrustGatewaySettingsSettings
	SetInternalValue(val *DataCloudflareZeroTrustGatewaySettingsSettings)
	ProtocolDetection() DataCloudflareZeroTrustGatewaySettingsSettingsProtocolDetectionOutputReference
	Sandbox() DataCloudflareZeroTrustGatewaySettingsSettingsSandboxOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsDecrypt() DataCloudflareZeroTrustGatewaySettingsSettingsTlsDecryptOutputReference
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

// The jsii proxy struct for DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference
type jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) ActivityLog() DataCloudflareZeroTrustGatewaySettingsSettingsActivityLogOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsActivityLogOutputReference
	_jsii_.Get(
		j,
		"activityLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) Antivirus() DataCloudflareZeroTrustGatewaySettingsSettingsAntivirusOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsAntivirusOutputReference
	_jsii_.Get(
		j,
		"antivirus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) BlockPage() DataCloudflareZeroTrustGatewaySettingsSettingsBlockPageOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsBlockPageOutputReference
	_jsii_.Get(
		j,
		"blockPage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) BodyScanning() DataCloudflareZeroTrustGatewaySettingsSettingsBodyScanningOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsBodyScanningOutputReference
	_jsii_.Get(
		j,
		"bodyScanning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) BrowserIsolation() DataCloudflareZeroTrustGatewaySettingsSettingsBrowserIsolationOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsBrowserIsolationOutputReference
	_jsii_.Get(
		j,
		"browserIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) Certificate() DataCloudflareZeroTrustGatewaySettingsSettingsCertificateOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsCertificateOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) CustomCertificate() DataCloudflareZeroTrustGatewaySettingsSettingsCustomCertificateOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsCustomCertificateOutputReference
	_jsii_.Get(
		j,
		"customCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) ExtendedEmailMatching() DataCloudflareZeroTrustGatewaySettingsSettingsExtendedEmailMatchingOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsExtendedEmailMatchingOutputReference
	_jsii_.Get(
		j,
		"extendedEmailMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) Fips() DataCloudflareZeroTrustGatewaySettingsSettingsFipsOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsFipsOutputReference
	_jsii_.Get(
		j,
		"fips",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) HostSelector() DataCloudflareZeroTrustGatewaySettingsSettingsHostSelectorOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsHostSelectorOutputReference
	_jsii_.Get(
		j,
		"hostSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) InternalValue() *DataCloudflareZeroTrustGatewaySettingsSettings {
	var returns *DataCloudflareZeroTrustGatewaySettingsSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) ProtocolDetection() DataCloudflareZeroTrustGatewaySettingsSettingsProtocolDetectionOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsProtocolDetectionOutputReference
	_jsii_.Get(
		j,
		"protocolDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) Sandbox() DataCloudflareZeroTrustGatewaySettingsSettingsSandboxOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsSandboxOutputReference
	_jsii_.Get(
		j,
		"sandbox",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) TlsDecrypt() DataCloudflareZeroTrustGatewaySettingsSettingsTlsDecryptOutputReference {
	var returns DataCloudflareZeroTrustGatewaySettingsSettingsTlsDecryptOutputReference
	_jsii_.Get(
		j,
		"tlsDecrypt",
		&returns,
	)
	return returns
}


func NewDataCloudflareZeroTrustGatewaySettingsSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewDataCloudflareZeroTrustGatewaySettingsSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustGatewaySettings.DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataCloudflareZeroTrustGatewaySettingsSettingsOutputReference_Override(d DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustGatewaySettings.DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference)SetInternalValue(val *DataCloudflareZeroTrustGatewaySettingsSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataCloudflareZeroTrustGatewaySettingsSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

