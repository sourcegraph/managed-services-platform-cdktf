package googlegkehubfeaturemembership

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlegkehubfeaturemembership/internal"
)

type GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference interface {
	cdktf.ComplexObject
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
	DeploymentOverrides() GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesList
	DeploymentOverridesInput() any
	Enabled() any
	SetEnabled(val any)
	EnabledInput() any
	// Experimental.
	Fqn() *string
	Git() GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGitOutputReference
	GitInput() *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit
	InternalValue() *GoogleGkeHubFeatureMembershipConfigmanagementConfigSync
	SetInternalValue(val *GoogleGkeHubFeatureMembershipConfigmanagementConfigSync)
	MetricsGcpServiceAccountEmail() *string
	SetMetricsGcpServiceAccountEmail(val *string)
	MetricsGcpServiceAccountEmailInput() *string
	Oci() GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOciOutputReference
	OciInput() *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOci
	PreventDrift() any
	SetPreventDrift(val any)
	PreventDriftInput() any
	SourceFormat() *string
	SetSourceFormat(val *string)
	SourceFormatInput() *string
	StopSyncing() any
	SetStopSyncing(val any)
	StopSyncingInput() any
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
	PutDeploymentOverrides(value any)
	PutGit(value *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit)
	PutOci(value *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOci)
	ResetDeploymentOverrides()
	ResetEnabled()
	ResetGit()
	ResetMetricsGcpServiceAccountEmail()
	ResetOci()
	ResetPreventDrift()
	ResetSourceFormat()
	ResetStopSyncing()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference
type jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) DeploymentOverrides() GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesList {
	var returns GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncDeploymentOverridesList
	_jsii_.Get(
		j,
		"deploymentOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) DeploymentOverridesInput() any {
	var returns any
	_jsii_.Get(
		j,
		"deploymentOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Enabled() any {
	var returns any
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) EnabledInput() any {
	var returns any
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Git() GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGitOutputReference {
	var returns GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGitOutputReference
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GitInput() *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit {
	var returns *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit
	_jsii_.Get(
		j,
		"gitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) InternalValue() *GoogleGkeHubFeatureMembershipConfigmanagementConfigSync {
	var returns *GoogleGkeHubFeatureMembershipConfigmanagementConfigSync
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) MetricsGcpServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGcpServiceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) MetricsGcpServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metricsGcpServiceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Oci() GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOciOutputReference {
	var returns GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOciOutputReference
	_jsii_.Get(
		j,
		"oci",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) OciInput() *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOci {
	var returns *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOci
	_jsii_.Get(
		j,
		"ociInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PreventDrift() any {
	var returns any
	_jsii_.Get(
		j,
		"preventDrift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PreventDriftInput() any {
	var returns any
	_jsii_.Get(
		j,
		"preventDriftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SourceFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SourceFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) StopSyncing() any {
	var returns any
	_jsii_.Get(
		j,
		"stopSyncing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) StopSyncingInput() any {
	var returns any
	_jsii_.Get(
		j,
		"stopSyncingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewGoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference_Override(g GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeHubFeatureMembership.GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference",
		[]any{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetEnabled(val any) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetInternalValue(val *GoogleGkeHubFeatureMembershipConfigmanagementConfigSync) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetMetricsGcpServiceAccountEmail(val *string) {
	if err := j.validateSetMetricsGcpServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metricsGcpServiceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetPreventDrift(val any) {
	if err := j.validateSetPreventDriftParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preventDrift",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetSourceFormat(val *string) {
	if err := j.validateSetSourceFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFormat",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetStopSyncing(val any) {
	if err := j.validateSetStopSyncingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stopSyncing",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PutDeploymentOverrides(value any) {
	if err := g.validatePutDeploymentOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putDeploymentOverrides",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PutGit(value *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncGit) {
	if err := g.validatePutGitParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGit",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) PutOci(value *GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOci) {
	if err := g.validatePutOciParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOci",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetDeploymentOverrides() {
	_jsii_.InvokeVoid(
		g,
		"resetDeploymentOverrides",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetGit() {
	_jsii_.InvokeVoid(
		g,
		"resetGit",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetMetricsGcpServiceAccountEmail() {
	_jsii_.InvokeVoid(
		g,
		"resetMetricsGcpServiceAccountEmail",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetOci() {
	_jsii_.InvokeVoid(
		g,
		"resetOci",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetPreventDrift() {
	_jsii_.InvokeVoid(
		g,
		"resetPreventDrift",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetSourceFormat() {
	_jsii_.InvokeVoid(
		g,
		"resetSourceFormat",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ResetStopSyncing() {
	_jsii_.InvokeVoid(
		g,
		"resetStopSyncing",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (g *jsiiProxy_GoogleGkeHubFeatureMembershipConfigmanagementConfigSyncOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
