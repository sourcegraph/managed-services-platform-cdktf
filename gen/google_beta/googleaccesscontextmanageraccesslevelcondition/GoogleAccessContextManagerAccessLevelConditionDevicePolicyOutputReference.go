package googleaccesscontextmanageraccesslevelcondition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleaccesscontextmanageraccesslevelcondition/internal"
)

type GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference interface {
	cdktf.ComplexObject
	AllowedDeviceManagementLevels() *[]*string
	SetAllowedDeviceManagementLevels(val *[]*string)
	AllowedDeviceManagementLevelsInput() *[]*string
	AllowedEncryptionStatuses() *[]*string
	SetAllowedEncryptionStatuses(val *[]*string)
	AllowedEncryptionStatusesInput() *[]*string
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
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleAccessContextManagerAccessLevelConditionDevicePolicy
	SetInternalValue(val *GoogleAccessContextManagerAccessLevelConditionDevicePolicy)
	OsConstraints() GoogleAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsList
	OsConstraintsInput() any
	RequireAdminApproval() any
	SetRequireAdminApproval(val any)
	RequireAdminApprovalInput() any
	RequireCorpOwned() any
	SetRequireCorpOwned(val any)
	RequireCorpOwnedInput() any
	RequireScreenLock() any
	SetRequireScreenLock(val any)
	RequireScreenLockInput() any
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
	PutOsConstraints(value any)
	ResetAllowedDeviceManagementLevels()
	ResetAllowedEncryptionStatuses()
	ResetOsConstraints()
	ResetRequireAdminApproval()
	ResetRequireCorpOwned()
	ResetRequireScreenLock()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference
type jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) AllowedDeviceManagementLevels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDeviceManagementLevels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) AllowedDeviceManagementLevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDeviceManagementLevelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) AllowedEncryptionStatuses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEncryptionStatuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) AllowedEncryptionStatusesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEncryptionStatusesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) InternalValue() *GoogleAccessContextManagerAccessLevelConditionDevicePolicy {
	var returns *GoogleAccessContextManagerAccessLevelConditionDevicePolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) OsConstraints() GoogleAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsList {
	var returns GoogleAccessContextManagerAccessLevelConditionDevicePolicyOsConstraintsList
	_jsii_.Get(
		j,
		"osConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) OsConstraintsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"osConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) RequireAdminApproval() any {
	var returns any
	_jsii_.Get(
		j,
		"requireAdminApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) RequireAdminApprovalInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requireAdminApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) RequireCorpOwned() any {
	var returns any
	_jsii_.Get(
		j,
		"requireCorpOwned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) RequireCorpOwnedInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requireCorpOwnedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) RequireScreenLock() any {
	var returns any
	_jsii_.Get(
		j,
		"requireScreenLock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) RequireScreenLockInput() any {
	var returns any
	_jsii_.Get(
		j,
		"requireScreenLockInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference_Override(g GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAccessContextManagerAccessLevelCondition.GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetAllowedDeviceManagementLevels(val *[]*string) {
	if err := j.validateSetAllowedDeviceManagementLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDeviceManagementLevels",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetAllowedEncryptionStatuses(val *[]*string) {
	if err := j.validateSetAllowedEncryptionStatusesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEncryptionStatuses",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetInternalValue(val *GoogleAccessContextManagerAccessLevelConditionDevicePolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetRequireAdminApproval(val any) {
	if err := j.validateSetRequireAdminApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAdminApproval",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetRequireCorpOwned(val any) {
	if err := j.validateSetRequireCorpOwnedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCorpOwned",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetRequireScreenLock(val any) {
	if err := j.validateSetRequireScreenLockParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireScreenLock",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) PutOsConstraints(value any) {
	if err := g.validatePutOsConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOsConstraints",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ResetAllowedDeviceManagementLevels() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedDeviceManagementLevels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ResetAllowedEncryptionStatuses() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowedEncryptionStatuses",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ResetOsConstraints() {
	_jsii_.InvokeVoid(
		g,
		"resetOsConstraints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ResetRequireAdminApproval() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireAdminApproval",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ResetRequireCorpOwned() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireCorpOwned",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ResetRequireScreenLock() {
	_jsii_.InvokeVoid(
		g,
		"resetRequireScreenLock",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleAccessContextManagerAccessLevelConditionDevicePolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
