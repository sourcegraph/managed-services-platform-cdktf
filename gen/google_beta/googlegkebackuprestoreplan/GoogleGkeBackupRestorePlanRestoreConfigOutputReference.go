package googlegkebackuprestoreplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlegkebackuprestoreplan/internal"
)

type GoogleGkeBackupRestorePlanRestoreConfigOutputReference interface {
	cdktf.ComplexObject
	AllNamespaces() interface{}
	SetAllNamespaces(val interface{})
	AllNamespacesInput() interface{}
	ClusterResourceConflictPolicy() *string
	SetClusterResourceConflictPolicy(val *string)
	ClusterResourceConflictPolicyInput() *string
	ClusterResourceRestoreScope() GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference
	ClusterResourceRestoreScopeInput() *GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope
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
	ExcludedNamespaces() GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespacesOutputReference
	ExcludedNamespacesInput() *GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespaces
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleGkeBackupRestorePlanRestoreConfig
	SetInternalValue(val *GoogleGkeBackupRestorePlanRestoreConfig)
	NamespacedResourceRestoreMode() *string
	SetNamespacedResourceRestoreMode(val *string)
	NamespacedResourceRestoreModeInput() *string
	NoNamespaces() interface{}
	SetNoNamespaces(val interface{})
	NoNamespacesInput() interface{}
	SelectedApplications() GoogleGkeBackupRestorePlanRestoreConfigSelectedApplicationsOutputReference
	SelectedApplicationsInput() *GoogleGkeBackupRestorePlanRestoreConfigSelectedApplications
	SelectedNamespaces() GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespacesOutputReference
	SelectedNamespacesInput() *GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespaces
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformationRules() GoogleGkeBackupRestorePlanRestoreConfigTransformationRulesList
	TransformationRulesInput() interface{}
	VolumeDataRestorePolicy() *string
	SetVolumeDataRestorePolicy(val *string)
	VolumeDataRestorePolicyInput() *string
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
	PutClusterResourceRestoreScope(value *GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope)
	PutExcludedNamespaces(value *GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespaces)
	PutSelectedApplications(value *GoogleGkeBackupRestorePlanRestoreConfigSelectedApplications)
	PutSelectedNamespaces(value *GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespaces)
	PutTransformationRules(value interface{})
	ResetAllNamespaces()
	ResetClusterResourceConflictPolicy()
	ResetClusterResourceRestoreScope()
	ResetExcludedNamespaces()
	ResetNamespacedResourceRestoreMode()
	ResetNoNamespaces()
	ResetSelectedApplications()
	ResetSelectedNamespaces()
	ResetTransformationRules()
	ResetVolumeDataRestorePolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleGkeBackupRestorePlanRestoreConfigOutputReference
type jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) AllNamespaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) AllNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceConflictPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceConflictPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceConflictPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceConflictPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceRestoreScope() GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference {
	var returns GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScopeOutputReference
	_jsii_.Get(
		j,
		"clusterResourceRestoreScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ClusterResourceRestoreScopeInput() *GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope {
	var returns *GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope
	_jsii_.Get(
		j,
		"clusterResourceRestoreScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ExcludedNamespaces() GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespacesOutputReference {
	var returns GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespacesOutputReference
	_jsii_.Get(
		j,
		"excludedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ExcludedNamespacesInput() *GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespaces {
	var returns *GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespaces
	_jsii_.Get(
		j,
		"excludedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) InternalValue() *GoogleGkeBackupRestorePlanRestoreConfig {
	var returns *GoogleGkeBackupRestorePlanRestoreConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) NamespacedResourceRestoreMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespacedResourceRestoreMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) NamespacedResourceRestoreModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespacedResourceRestoreModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) NoNamespaces() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) NoNamespacesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) SelectedApplications() GoogleGkeBackupRestorePlanRestoreConfigSelectedApplicationsOutputReference {
	var returns GoogleGkeBackupRestorePlanRestoreConfigSelectedApplicationsOutputReference
	_jsii_.Get(
		j,
		"selectedApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) SelectedApplicationsInput() *GoogleGkeBackupRestorePlanRestoreConfigSelectedApplications {
	var returns *GoogleGkeBackupRestorePlanRestoreConfigSelectedApplications
	_jsii_.Get(
		j,
		"selectedApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) SelectedNamespaces() GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespacesOutputReference {
	var returns GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespacesOutputReference
	_jsii_.Get(
		j,
		"selectedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) SelectedNamespacesInput() *GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespaces {
	var returns *GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespaces
	_jsii_.Get(
		j,
		"selectedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) TransformationRules() GoogleGkeBackupRestorePlanRestoreConfigTransformationRulesList {
	var returns GoogleGkeBackupRestorePlanRestoreConfigTransformationRulesList
	_jsii_.Get(
		j,
		"transformationRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) TransformationRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformationRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) VolumeDataRestorePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeDataRestorePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) VolumeDataRestorePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeDataRestorePolicyInput",
		&returns,
	)
	return returns
}


func NewGoogleGkeBackupRestorePlanRestoreConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleGkeBackupRestorePlanRestoreConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleGkeBackupRestorePlanRestoreConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeBackupRestorePlan.GoogleGkeBackupRestorePlanRestoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleGkeBackupRestorePlanRestoreConfigOutputReference_Override(g GoogleGkeBackupRestorePlanRestoreConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleGkeBackupRestorePlan.GoogleGkeBackupRestorePlanRestoreConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetAllNamespaces(val interface{}) {
	if err := j.validateSetAllNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allNamespaces",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetClusterResourceConflictPolicy(val *string) {
	if err := j.validateSetClusterResourceConflictPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterResourceConflictPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetInternalValue(val *GoogleGkeBackupRestorePlanRestoreConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetNamespacedResourceRestoreMode(val *string) {
	if err := j.validateSetNamespacedResourceRestoreModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespacedResourceRestoreMode",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetNoNamespaces(val interface{}) {
	if err := j.validateSetNoNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noNamespaces",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference)SetVolumeDataRestorePolicy(val *string) {
	if err := j.validateSetVolumeDataRestorePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeDataRestorePolicy",
		val,
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) PutClusterResourceRestoreScope(value *GoogleGkeBackupRestorePlanRestoreConfigClusterResourceRestoreScope) {
	if err := g.validatePutClusterResourceRestoreScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClusterResourceRestoreScope",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) PutExcludedNamespaces(value *GoogleGkeBackupRestorePlanRestoreConfigExcludedNamespaces) {
	if err := g.validatePutExcludedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludedNamespaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) PutSelectedApplications(value *GoogleGkeBackupRestorePlanRestoreConfigSelectedApplications) {
	if err := g.validatePutSelectedApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedApplications",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) PutSelectedNamespaces(value *GoogleGkeBackupRestorePlanRestoreConfigSelectedNamespaces) {
	if err := g.validatePutSelectedNamespacesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSelectedNamespaces",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) PutTransformationRules(value interface{}) {
	if err := g.validatePutTransformationRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTransformationRules",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetAllNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetAllNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetClusterResourceConflictPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterResourceConflictPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetClusterResourceRestoreScope() {
	_jsii_.InvokeVoid(
		g,
		"resetClusterResourceRestoreScope",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetExcludedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetNamespacedResourceRestoreMode() {
	_jsii_.InvokeVoid(
		g,
		"resetNamespacedResourceRestoreMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetNoNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetNoNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetSelectedApplications() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedApplications",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetSelectedNamespaces() {
	_jsii_.InvokeVoid(
		g,
		"resetSelectedNamespaces",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetTransformationRules() {
	_jsii_.InvokeVoid(
		g,
		"resetTransformationRules",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ResetVolumeDataRestorePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetVolumeDataRestorePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleGkeBackupRestorePlanRestoreConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

