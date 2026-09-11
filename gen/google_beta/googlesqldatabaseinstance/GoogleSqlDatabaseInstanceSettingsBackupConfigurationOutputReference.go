package googlesqldatabaseinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlesqldatabaseinstance/internal"
)

type GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference interface {
	cdktf.ComplexObject
	BackupRetentionSettings() GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettingsOutputReference
	BackupRetentionSettingsInput() *GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings
	BinaryLogEnabled() any
	SetBinaryLogEnabled(val any)
	BinaryLogEnabledInput() any
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
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
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleSqlDatabaseInstanceSettingsBackupConfiguration
	SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsBackupConfiguration)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	PointInTimeRecoveryEnabled() any
	SetPointInTimeRecoveryEnabled(val any)
	PointInTimeRecoveryEnabledInput() any
	StartTime() *string
	SetStartTime(val *string)
	StartTimeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransactionLogRetentionDays() *float64
	SetTransactionLogRetentionDays(val *float64)
	TransactionLogRetentionDaysInput() *float64
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
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
	PutBackupRetentionSettings(value *GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings)
	ResetBackupRetentionSettings()
	ResetBinaryLogEnabled()
	ResetEnabled()
	ResetLocation()
	ResetPointInTimeRecoveryEnabled()
	ResetStartTime()
	ResetTransactionLogRetentionDays()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference
type jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) BackupRetentionSettings() GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettingsOutputReference {
	var returns GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettingsOutputReference
	_jsii_.Get(
		j,
		"backupRetentionSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) BackupRetentionSettingsInput() *GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings {
	var returns *GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings
	_jsii_.Get(
		j,
		"backupRetentionSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) BinaryLogEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"binaryLogEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) BinaryLogEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"binaryLogEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) InternalValue() *GoogleSqlDatabaseInstanceSettingsBackupConfiguration {
	var returns *GoogleSqlDatabaseInstanceSettingsBackupConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) PointInTimeRecoveryEnabled() any {
	var returns any
	_jsii_.Get(
		j,
		"pointInTimeRecoveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) PointInTimeRecoveryEnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"pointInTimeRecoveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) StartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) TransactionLogRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionLogRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) TransactionLogRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"transactionLogRetentionDaysInput",
		&returns,
	)
	return returns
}

func NewGoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference_Override(g GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleSqlDatabaseInstance.GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetBinaryLogEnabled(val any) {
	if err := j.validateSetBinaryLogEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binaryLogEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetInternalValue(val *GoogleSqlDatabaseInstanceSettingsBackupConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetPointInTimeRecoveryEnabled(val any) {
	if err := j.validateSetPointInTimeRecoveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pointInTimeRecoveryEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetStartTime(val *string) {
	if err := j.validateSetStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startTime",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) SetTransactionLogRetentionDays(val *float64) {
	if err := j.validateSetTransactionLogRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transactionLogRetentionDays",
		val,
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) PutBackupRetentionSettings(value *GoogleSqlDatabaseInstanceSettingsBackupConfigurationBackupRetentionSettings) {
	if err := g.validatePutBackupRetentionSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackupRetentionSettings",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ResetBackupRetentionSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupRetentionSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ResetBinaryLogEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetBinaryLogEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ResetPointInTimeRecoveryEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetPointInTimeRecoveryEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ResetStartTime() {
	_jsii_.InvokeVoid(
		g,
		"resetStartTime",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ResetTransactionLogRetentionDays() {
	_jsii_.InvokeVoid(
		g,
		"resetTransactionLogRetentionDays",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		g,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleSqlDatabaseInstanceSettingsBackupConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
