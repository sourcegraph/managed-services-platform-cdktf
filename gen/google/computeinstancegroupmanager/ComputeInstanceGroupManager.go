package computeinstancegroupmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/computeinstancegroupmanager/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_group_manager google_compute_instance_group_manager}.
type ComputeInstanceGroupManager interface {
	cdktf.TerraformResource
	AllInstancesConfig() ComputeInstanceGroupManagerAllInstancesConfigOutputReference
	AllInstancesConfigInput() *ComputeInstanceGroupManagerAllInstancesConfig
	AutoHealingPolicies() ComputeInstanceGroupManagerAutoHealingPoliciesOutputReference
	AutoHealingPoliciesInput() *ComputeInstanceGroupManagerAutoHealingPolicies
	BaseInstanceName() *string
	SetBaseInstanceName(val *string)
	BaseInstanceNameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Fingerprint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceGroup() *string
	InstanceGroupManagerId() *float64
	InstanceLifecyclePolicy() ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	InstanceLifecyclePolicyInput() *ComputeInstanceGroupManagerInstanceLifecyclePolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListManagedInstancesResults() *string
	SetListManagedInstancesResults(val *string)
	ListManagedInstancesResultsInput() *string
	Name() *string
	SetName(val *string)
	NamedPort() ComputeInstanceGroupManagerNamedPortList
	NamedPortInput() interface{}
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Operation() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ResourcePolicies() ComputeInstanceGroupManagerResourcePoliciesOutputReference
	ResourcePoliciesInput() *ComputeInstanceGroupManagerResourcePolicies
	SelfLink() *string
	StandbyPolicy() ComputeInstanceGroupManagerStandbyPolicyOutputReference
	StandbyPolicyInput() *ComputeInstanceGroupManagerStandbyPolicy
	StatefulDisk() ComputeInstanceGroupManagerStatefulDiskList
	StatefulDiskInput() interface{}
	StatefulExternalIp() ComputeInstanceGroupManagerStatefulExternalIpList
	StatefulExternalIpInput() interface{}
	StatefulInternalIp() ComputeInstanceGroupManagerStatefulInternalIpList
	StatefulInternalIpInput() interface{}
	Status() ComputeInstanceGroupManagerStatusList
	TargetPools() *[]*string
	SetTargetPools(val *[]*string)
	TargetPoolsInput() *[]*string
	TargetSize() *float64
	SetTargetSize(val *float64)
	TargetSizeInput() *float64
	TargetStoppedSize() *float64
	SetTargetStoppedSize(val *float64)
	TargetStoppedSizeInput() *float64
	TargetSuspendedSize() *float64
	SetTargetSuspendedSize(val *float64)
	TargetSuspendedSizeInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() ComputeInstanceGroupManagerTimeoutsOutputReference
	TimeoutsInput() interface{}
	UpdatePolicy() ComputeInstanceGroupManagerUpdatePolicyOutputReference
	UpdatePolicyInput() *ComputeInstanceGroupManagerUpdatePolicy
	Version() ComputeInstanceGroupManagerVersionList
	VersionInput() interface{}
	WaitForInstances() interface{}
	SetWaitForInstances(val interface{})
	WaitForInstancesInput() interface{}
	WaitForInstancesStatus() *string
	SetWaitForInstancesStatus(val *string)
	WaitForInstancesStatusInput() *string
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAllInstancesConfig(value *ComputeInstanceGroupManagerAllInstancesConfig)
	PutAutoHealingPolicies(value *ComputeInstanceGroupManagerAutoHealingPolicies)
	PutInstanceLifecyclePolicy(value *ComputeInstanceGroupManagerInstanceLifecyclePolicy)
	PutNamedPort(value interface{})
	PutResourcePolicies(value *ComputeInstanceGroupManagerResourcePolicies)
	PutStandbyPolicy(value *ComputeInstanceGroupManagerStandbyPolicy)
	PutStatefulDisk(value interface{})
	PutStatefulExternalIp(value interface{})
	PutStatefulInternalIp(value interface{})
	PutTimeouts(value *ComputeInstanceGroupManagerTimeouts)
	PutUpdatePolicy(value *ComputeInstanceGroupManagerUpdatePolicy)
	PutVersion(value interface{})
	ResetAllInstancesConfig()
	ResetAutoHealingPolicies()
	ResetDescription()
	ResetId()
	ResetInstanceLifecyclePolicy()
	ResetListManagedInstancesResults()
	ResetNamedPort()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetResourcePolicies()
	ResetStandbyPolicy()
	ResetStatefulDisk()
	ResetStatefulExternalIp()
	ResetStatefulInternalIp()
	ResetTargetPools()
	ResetTargetSize()
	ResetTargetStoppedSize()
	ResetTargetSuspendedSize()
	ResetTimeouts()
	ResetUpdatePolicy()
	ResetWaitForInstances()
	ResetWaitForInstancesStatus()
	ResetZone()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ComputeInstanceGroupManager
type jsiiProxy_ComputeInstanceGroupManager struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ComputeInstanceGroupManager) AllInstancesConfig() ComputeInstanceGroupManagerAllInstancesConfigOutputReference {
	var returns ComputeInstanceGroupManagerAllInstancesConfigOutputReference
	_jsii_.Get(
		j,
		"allInstancesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) AllInstancesConfigInput() *ComputeInstanceGroupManagerAllInstancesConfig {
	var returns *ComputeInstanceGroupManagerAllInstancesConfig
	_jsii_.Get(
		j,
		"allInstancesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) AutoHealingPolicies() ComputeInstanceGroupManagerAutoHealingPoliciesOutputReference {
	var returns ComputeInstanceGroupManagerAutoHealingPoliciesOutputReference
	_jsii_.Get(
		j,
		"autoHealingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) AutoHealingPoliciesInput() *ComputeInstanceGroupManagerAutoHealingPolicies {
	var returns *ComputeInstanceGroupManagerAutoHealingPolicies
	_jsii_.Get(
		j,
		"autoHealingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) BaseInstanceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) BaseInstanceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseInstanceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) InstanceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) InstanceGroupManagerId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instanceGroupManagerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) InstanceLifecyclePolicy() ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference {
	var returns ComputeInstanceGroupManagerInstanceLifecyclePolicyOutputReference
	_jsii_.Get(
		j,
		"instanceLifecyclePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) InstanceLifecyclePolicyInput() *ComputeInstanceGroupManagerInstanceLifecyclePolicy {
	var returns *ComputeInstanceGroupManagerInstanceLifecyclePolicy
	_jsii_.Get(
		j,
		"instanceLifecyclePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ListManagedInstancesResults() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ListManagedInstancesResultsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"listManagedInstancesResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) NamedPort() ComputeInstanceGroupManagerNamedPortList {
	var returns ComputeInstanceGroupManagerNamedPortList
	_jsii_.Get(
		j,
		"namedPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) NamedPortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namedPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ResourcePolicies() ComputeInstanceGroupManagerResourcePoliciesOutputReference {
	var returns ComputeInstanceGroupManagerResourcePoliciesOutputReference
	_jsii_.Get(
		j,
		"resourcePolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ResourcePoliciesInput() *ComputeInstanceGroupManagerResourcePolicies {
	var returns *ComputeInstanceGroupManagerResourcePolicies
	_jsii_.Get(
		j,
		"resourcePoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StandbyPolicy() ComputeInstanceGroupManagerStandbyPolicyOutputReference {
	var returns ComputeInstanceGroupManagerStandbyPolicyOutputReference
	_jsii_.Get(
		j,
		"standbyPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StandbyPolicyInput() *ComputeInstanceGroupManagerStandbyPolicy {
	var returns *ComputeInstanceGroupManagerStandbyPolicy
	_jsii_.Get(
		j,
		"standbyPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StatefulDisk() ComputeInstanceGroupManagerStatefulDiskList {
	var returns ComputeInstanceGroupManagerStatefulDiskList
	_jsii_.Get(
		j,
		"statefulDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StatefulDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StatefulExternalIp() ComputeInstanceGroupManagerStatefulExternalIpList {
	var returns ComputeInstanceGroupManagerStatefulExternalIpList
	_jsii_.Get(
		j,
		"statefulExternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StatefulExternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulExternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StatefulInternalIp() ComputeInstanceGroupManagerStatefulInternalIpList {
	var returns ComputeInstanceGroupManagerStatefulInternalIpList
	_jsii_.Get(
		j,
		"statefulInternalIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) StatefulInternalIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statefulInternalIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Status() ComputeInstanceGroupManagerStatusList {
	var returns ComputeInstanceGroupManagerStatusList
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetPoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetStoppedSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetStoppedSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetStoppedSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetStoppedSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetSuspendedSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSuspendedSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TargetSuspendedSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetSuspendedSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Timeouts() ComputeInstanceGroupManagerTimeoutsOutputReference {
	var returns ComputeInstanceGroupManagerTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) UpdatePolicy() ComputeInstanceGroupManagerUpdatePolicyOutputReference {
	var returns ComputeInstanceGroupManagerUpdatePolicyOutputReference
	_jsii_.Get(
		j,
		"updatePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) UpdatePolicyInput() *ComputeInstanceGroupManagerUpdatePolicy {
	var returns *ComputeInstanceGroupManagerUpdatePolicy
	_jsii_.Get(
		j,
		"updatePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Version() ComputeInstanceGroupManagerVersionList {
	var returns ComputeInstanceGroupManagerVersionList
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) VersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) WaitForInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) WaitForInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"waitForInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) WaitForInstancesStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) WaitForInstancesStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"waitForInstancesStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ComputeInstanceGroupManager) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_group_manager google_compute_instance_group_manager} Resource.
func NewComputeInstanceGroupManager(scope constructs.Construct, id *string, config *ComputeInstanceGroupManagerConfig) ComputeInstanceGroupManager {
	_init_.Initialize()

	if err := validateNewComputeInstanceGroupManagerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ComputeInstanceGroupManager{}

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManager",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_instance_group_manager google_compute_instance_group_manager} Resource.
func NewComputeInstanceGroupManager_Override(c ComputeInstanceGroupManager, scope constructs.Construct, id *string, config *ComputeInstanceGroupManagerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManager",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetBaseInstanceName(val *string) {
	if err := j.validateSetBaseInstanceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseInstanceName",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetListManagedInstancesResults(val *string) {
	if err := j.validateSetListManagedInstancesResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listManagedInstancesResults",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetTargetPools(val *[]*string) {
	if err := j.validateSetTargetPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPools",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetTargetSize(val *float64) {
	if err := j.validateSetTargetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSize",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetTargetStoppedSize(val *float64) {
	if err := j.validateSetTargetStoppedSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetStoppedSize",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetTargetSuspendedSize(val *float64) {
	if err := j.validateSetTargetSuspendedSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetSuspendedSize",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetWaitForInstances(val interface{}) {
	if err := j.validateSetWaitForInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstances",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetWaitForInstancesStatus(val *string) {
	if err := j.validateSetWaitForInstancesStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitForInstancesStatus",
		val,
	)
}

func (j *jsiiProxy_ComputeInstanceGroupManager)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a ComputeInstanceGroupManager resource upon running "cdktf plan <stack-name>".
func ComputeInstanceGroupManager_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateComputeInstanceGroupManager_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManager",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func ComputeInstanceGroupManager_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInstanceGroupManager_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManager",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeInstanceGroupManager_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInstanceGroupManager_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManager",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ComputeInstanceGroupManager_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateComputeInstanceGroupManager_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManager",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ComputeInstanceGroupManager_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.computeInstanceGroupManager.ComputeInstanceGroupManager",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutAllInstancesConfig(value *ComputeInstanceGroupManagerAllInstancesConfig) {
	if err := c.validatePutAllInstancesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAllInstancesConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutAutoHealingPolicies(value *ComputeInstanceGroupManagerAutoHealingPolicies) {
	if err := c.validatePutAutoHealingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAutoHealingPolicies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutInstanceLifecyclePolicy(value *ComputeInstanceGroupManagerInstanceLifecyclePolicy) {
	if err := c.validatePutInstanceLifecyclePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putInstanceLifecyclePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutNamedPort(value interface{}) {
	if err := c.validatePutNamedPortParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putNamedPort",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutResourcePolicies(value *ComputeInstanceGroupManagerResourcePolicies) {
	if err := c.validatePutResourcePoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putResourcePolicies",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutStandbyPolicy(value *ComputeInstanceGroupManagerStandbyPolicy) {
	if err := c.validatePutStandbyPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStandbyPolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutStatefulDisk(value interface{}) {
	if err := c.validatePutStatefulDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStatefulDisk",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutStatefulExternalIp(value interface{}) {
	if err := c.validatePutStatefulExternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStatefulExternalIp",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutStatefulInternalIp(value interface{}) {
	if err := c.validatePutStatefulInternalIpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putStatefulInternalIp",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutTimeouts(value *ComputeInstanceGroupManagerTimeouts) {
	if err := c.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutUpdatePolicy(value *ComputeInstanceGroupManagerUpdatePolicy) {
	if err := c.validatePutUpdatePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUpdatePolicy",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) PutVersion(value interface{}) {
	if err := c.validatePutVersionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVersion",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetAllInstancesConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetAllInstancesConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetAutoHealingPolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetAutoHealingPolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetInstanceLifecyclePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetInstanceLifecyclePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetListManagedInstancesResults() {
	_jsii_.InvokeVoid(
		c,
		"resetListManagedInstancesResults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetNamedPort() {
	_jsii_.InvokeVoid(
		c,
		"resetNamedPort",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetProject() {
	_jsii_.InvokeVoid(
		c,
		"resetProject",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetResourcePolicies() {
	_jsii_.InvokeVoid(
		c,
		"resetResourcePolicies",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetStandbyPolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetStandbyPolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetStatefulDisk() {
	_jsii_.InvokeVoid(
		c,
		"resetStatefulDisk",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetStatefulExternalIp() {
	_jsii_.InvokeVoid(
		c,
		"resetStatefulExternalIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetStatefulInternalIp() {
	_jsii_.InvokeVoid(
		c,
		"resetStatefulInternalIp",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetTargetPools() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetPools",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetTargetSize() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetTargetStoppedSize() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetStoppedSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetTargetSuspendedSize() {
	_jsii_.InvokeVoid(
		c,
		"resetTargetSuspendedSize",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetTimeouts() {
	_jsii_.InvokeVoid(
		c,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetUpdatePolicy() {
	_jsii_.InvokeVoid(
		c,
		"resetUpdatePolicy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetWaitForInstances() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForInstances",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetWaitForInstancesStatus() {
	_jsii_.InvokeVoid(
		c,
		"resetWaitForInstancesStatus",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ResetZone() {
	_jsii_.InvokeVoid(
		c,
		"resetZone",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ComputeInstanceGroupManager) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ComputeInstanceGroupManager) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

