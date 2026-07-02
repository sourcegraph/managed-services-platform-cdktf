package googlegkebackupbackupplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlegkebackupbackupplan/internal"
)

type GoogleGkeBackupBackupPlanBackupConfigOutputReference interface {
	cdktf.ComplexObject
	AllNamespaces() any
	SetAllNamespaces(val any)
	AllNamespacesInput() any
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
	EncryptionKey() GoogleGkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	EncryptionKeyInput() *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey
	// Experimental.
	Fqn() *string
	IncludeSecrets() any
	SetIncludeSecrets(val any)
	IncludeSecretsInput() any
	IncludeVolumeData() any
	SetIncludeVolumeData(val any)
	IncludeVolumeDataInput() any
	InternalValue() *GoogleGkeBackupBackupPlanBackupConfig
	SetInternalValue(val *GoogleGkeBackupBackupPlanBackupConfig)
	PermissiveMode() any
	SetPermissiveMode(val any)
	PermissiveModeInput() any
	SelectedApplications() GoogleGkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	SelectedApplicationsInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications
	SelectedNamespaces() GoogleGkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	SelectedNamespacesInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces
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
	PutEncryptionKey(value *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey)
	PutSelectedApplications(value *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications)
	PutSelectedNamespaces(value *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces)
	ResetAllNamespaces()
	ResetEncryptionKey()
	ResetIncludeSecrets()
	ResetIncludeVolumeData()
	ResetPermissiveMode()
	ResetSelectedApplications()
	ResetSelectedNamespaces()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeBackupBackupPlanBackupConfigOutputReference
type jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) AllNamespaces() any {
	var returns any
	_jsii_.Get(
		j,
		"allNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) AllNamespacesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"allNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) EncryptionKey() GoogleGkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupConfigEncryptionKeyOutputReference
	_jsii_.Get(
		j,
		"encryptionKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) EncryptionKeyInput() *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey {
	var returns *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey
	_jsii_.Get(
		j,
		"encryptionKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeSecrets() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeSecretsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeData() any {
	var returns any
	_jsii_.Get(
		j,
		"includeVolumeData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) IncludeVolumeDataInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeVolumeDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) InternalValue() *GoogleGkeBackupBackupPlanBackupConfig {
	var returns *GoogleGkeBackupBackupPlanBackupConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PermissiveMode() any {
	var returns any
	_jsii_.Get(
		j,
		"permissiveMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PermissiveModeInput() any {
	var returns any
	_jsii_.Get(
		j,
		"permissiveModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedApplications() GoogleGkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupConfigSelectedApplicationsOutputReference
	_jsii_.Get(
		j,
		"selectedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedApplicationsInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications {
	var returns *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications
	_jsii_.Get(
		j,
		"selectedApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespaces() GoogleGkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference {
	var returns GoogleGkeBackupBackupPlanBackupConfigSelectedNamespacesOutputReference
	_jsii_.Get(
		j,
		"selectedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SelectedNamespacesInput() *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces {
	var returns *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces
	_jsii_.Get(
		j,
		"selectedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleGkeBackupBackupPlanBackupConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeBackupBackupPlanBackupConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeBackupBackupPlanBackupConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeBackupBackupPlan.GoogleGkeBackupBackupPlanBackupConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeBackupBackupPlanBackupConfigOutputReference_Override(g GoogleGkeBackupBackupPlanBackupConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeBackupBackupPlan.GoogleGkeBackupBackupPlanBackupConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetAllNamespaces(val any) {
	if err := j.validateSetAllNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allNamespaces",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetIncludeSecrets(val any) {
	if err := j.validateSetIncludeSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSecrets",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetIncludeVolumeData(val any) {
	if err := j.validateSetIncludeVolumeDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeVolumeData",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetInternalValue(val *GoogleGkeBackupBackupPlanBackupConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetPermissiveMode(val any) {
	if err := j.validateSetPermissiveModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissiveMode",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PutEncryptionKey(value *GoogleGkeBackupBackupPlanBackupConfigEncryptionKey) {
	if err := g.validatePutEncryptionKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putEncryptionKey",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PutSelectedApplications(value *GoogleGkeBackupBackupPlanBackupConfigSelectedApplications) {
	if err := g.validatePutSelectedApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedApplications",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) PutSelectedNamespaces(value *GoogleGkeBackupBackupPlanBackupConfigSelectedNamespaces) {
	if err := g.validatePutSelectedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedNamespaces",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetAllNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetAllNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetEncryptionKey() {
	_jsii_.InvokeVoid(
		g,
		"resetEncryptionKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeSecrets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeSecrets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetIncludeVolumeData() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludeVolumeData",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetPermissiveMode() {
	_jsii_.InvokeVoid(
		g,
		"resetPermissiveMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedApplications() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedApplications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ResetSelectedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleGkeBackupBackupPlanBackupConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
