package googlecontainernodepool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecontainernodepool/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool google_container_node_pool}.
type GoogleContainerNodePool interface {
	cdktf.TerraformResource
	Autoscaling() GoogleContainerNodePoolAutoscalingOutputReference
	AutoscalingInput() *GoogleContainerNodePoolAutoscaling
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
	// Experimental.
	Connection() any
	// Experimental.
	SetConnection(val any)
	// Experimental.
	ConstructNodeMetadata() *map[string]any
	// Experimental.
	Count() any
	// Experimental.
	SetCount(val any)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	InitialNodeCount() *float64
	SetInitialNodeCount(val *float64)
	InitialNodeCountInput() *float64
	InstanceGroupUrls() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedInstanceGroupUrls() *[]*string
	Management() GoogleContainerNodePoolManagementOutputReference
	ManagementInput() *GoogleContainerNodePoolManagement
	MaxPodsPerNode() *float64
	SetMaxPodsPerNode(val *float64)
	MaxPodsPerNodeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	NetworkConfig() GoogleContainerNodePoolNetworkConfigOutputReference
	NetworkConfigInput() *GoogleContainerNodePoolNetworkConfig
	// The tree node.
	Node() constructs.Node
	NodeConfig() GoogleContainerNodePoolNodeConfigOutputReference
	NodeConfigInput() *GoogleContainerNodePoolNodeConfig
	NodeCount() *float64
	SetNodeCount(val *float64)
	NodeCountInput() *float64
	NodeLocations() *[]*string
	SetNodeLocations(val *[]*string)
	NodeLocationsInput() *[]*string
	Operation() *string
	PlacementPolicy() GoogleContainerNodePoolPlacementPolicyOutputReference
	PlacementPolicyInput() *GoogleContainerNodePoolPlacementPolicy
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]any
	// Experimental.
	SetProvisioners(val *[]any)
	QueuedProvisioning() GoogleContainerNodePoolQueuedProvisioningOutputReference
	QueuedProvisioningInput() *GoogleContainerNodePoolQueuedProvisioning
	// Experimental.
	RawOverrides() any
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]any
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleContainerNodePoolTimeoutsOutputReference
	TimeoutsInput() any
	UpgradeSettings() GoogleContainerNodePoolUpgradeSettingsOutputReference
	UpgradeSettingsInput() *GoogleContainerNodePoolUpgradeSettings
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value any)
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
	HasResourceMove() any
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
	MoveTo(moveTarget *string, index any)
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAutoscaling(value *GoogleContainerNodePoolAutoscaling)
	PutManagement(value *GoogleContainerNodePoolManagement)
	PutNetworkConfig(value *GoogleContainerNodePoolNetworkConfig)
	PutNodeConfig(value *GoogleContainerNodePoolNodeConfig)
	PutPlacementPolicy(value *GoogleContainerNodePoolPlacementPolicy)
	PutQueuedProvisioning(value *GoogleContainerNodePoolQueuedProvisioning)
	PutTimeouts(value *GoogleContainerNodePoolTimeouts)
	PutUpgradeSettings(value *GoogleContainerNodePoolUpgradeSettings)
	ResetAutoscaling()
	ResetId()
	ResetInitialNodeCount()
	ResetLocation()
	ResetManagement()
	ResetMaxPodsPerNode()
	ResetName()
	ResetNamePrefix()
	ResetNetworkConfig()
	ResetNodeConfig()
	ResetNodeCount()
	ResetNodeLocations()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlacementPolicy()
	ResetProject()
	ResetQueuedProvisioning()
	ResetTimeouts()
	ResetUpgradeSettings()
	ResetVersion()
	SynthesizeAttributes() *map[string]any
	SynthesizeHclAttributes() *map[string]any
	// Experimental.
	ToHclTerraform() any
	// Experimental.
	ToMetadata() any
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() any
}

// The jsii proxy struct for GoogleContainerNodePool
type jsiiProxy_GoogleContainerNodePool struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleContainerNodePool) Autoscaling() GoogleContainerNodePoolAutoscalingOutputReference {
	var returns GoogleContainerNodePoolAutoscalingOutputReference
	_jsii_.Get(
		j,
		"autoscaling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) AutoscalingInput() *GoogleContainerNodePoolAutoscaling {
	var returns *GoogleContainerNodePoolAutoscaling
	_jsii_.Get(
		j,
		"autoscalingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Connection() any {
	var returns any
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) ConstructNodeMetadata() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Count() any {
	var returns any
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) InitialNodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) InitialNodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialNodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) InstanceGroupUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceGroupUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) ManagedInstanceGroupUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"managedInstanceGroupUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Management() GoogleContainerNodePoolManagementOutputReference {
	var returns GoogleContainerNodePoolManagementOutputReference
	_jsii_.Get(
		j,
		"management",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) ManagementInput() *GoogleContainerNodePoolManagement {
	var returns *GoogleContainerNodePoolManagement
	_jsii_.Get(
		j,
		"managementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) MaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) MaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NetworkConfig() GoogleContainerNodePoolNetworkConfigOutputReference {
	var returns GoogleContainerNodePoolNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NetworkConfigInput() *GoogleContainerNodePoolNetworkConfig {
	var returns *GoogleContainerNodePoolNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NodeConfig() GoogleContainerNodePoolNodeConfigOutputReference {
	var returns GoogleContainerNodePoolNodeConfigOutputReference
	_jsii_.Get(
		j,
		"nodeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NodeConfigInput() *GoogleContainerNodePoolNodeConfig {
	var returns *GoogleContainerNodePoolNodeConfig
	_jsii_.Get(
		j,
		"nodeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NodeCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NodeCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nodeCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NodeLocations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) NodeLocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"nodeLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Operation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) PlacementPolicy() GoogleContainerNodePoolPlacementPolicyOutputReference {
	var returns GoogleContainerNodePoolPlacementPolicyOutputReference
	_jsii_.Get(
		j,
		"placementPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) PlacementPolicyInput() *GoogleContainerNodePoolPlacementPolicy {
	var returns *GoogleContainerNodePoolPlacementPolicy
	_jsii_.Get(
		j,
		"placementPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Provisioners() *[]any {
	var returns *[]any
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) QueuedProvisioning() GoogleContainerNodePoolQueuedProvisioningOutputReference {
	var returns GoogleContainerNodePoolQueuedProvisioningOutputReference
	_jsii_.Get(
		j,
		"queuedProvisioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) QueuedProvisioningInput() *GoogleContainerNodePoolQueuedProvisioning {
	var returns *GoogleContainerNodePoolQueuedProvisioning
	_jsii_.Get(
		j,
		"queuedProvisioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) RawOverrides() any {
	var returns any
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) TerraformMetaArguments() *map[string]any {
	var returns *map[string]any
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Timeouts() GoogleContainerNodePoolTimeoutsOutputReference {
	var returns GoogleContainerNodePoolTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) TimeoutsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) UpgradeSettings() GoogleContainerNodePoolUpgradeSettingsOutputReference {
	var returns GoogleContainerNodePoolUpgradeSettingsOutputReference
	_jsii_.Get(
		j,
		"upgradeSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) UpgradeSettingsInput() *GoogleContainerNodePoolUpgradeSettings {
	var returns *GoogleContainerNodePoolUpgradeSettings
	_jsii_.Get(
		j,
		"upgradeSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleContainerNodePool) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool google_container_node_pool} Resource.
func NewGoogleContainerNodePool(scope constructs.Construct, id *string, config *GoogleContainerNodePoolConfig) GoogleContainerNodePool {
	_init_.Initialize()

	if err := validateNewGoogleContainerNodePoolParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleContainerNodePool{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePool",
		[]any{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_node_pool google_container_node_pool} Resource.
func NewGoogleContainerNodePool_Override(g GoogleContainerNodePool, scope constructs.Construct, id *string, config *GoogleContainerNodePoolConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePool",
		[]any{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetConnection(val any) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetCount(val any) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetInitialNodeCount(val *float64) {
	if err := j.validateSetInitialNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialNodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetMaxPodsPerNode(val *float64) {
	if err := j.validateSetMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetNodeCount(val *float64) {
	if err := j.validateSetNodeCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeCount",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetNodeLocations(val *[]*string) {
	if err := j.validateSetNodeLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeLocations",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetProvisioners(val *[]any) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleContainerNodePool) SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a GoogleContainerNodePool resource upon running "cdktf plan <stack-name>".
func GoogleContainerNodePool_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleContainerNodePool_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePool",
		"generateConfigForImport",
		[]any{scope, importToId, importFromId, provider},
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
func GoogleContainerNodePool_IsConstruct(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerNodePool_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePool",
		"isConstruct",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContainerNodePool_IsTerraformElement(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerNodePool_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePool",
		"isTerraformElement",
		[]any{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleContainerNodePool_IsTerraformResource(x any) *bool {
	_init_.Initialize()

	if err := validateGoogleContainerNodePool_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePool",
		"isTerraformResource",
		[]any{x},
		&returns,
	)

	return returns
}

func GoogleContainerNodePool_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleContainerNodePool.GoogleContainerNodePool",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]any{moveTarget},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) AddOverride(path *string, value any) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]any{path, value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleContainerNodePool) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleContainerNodePool) HasResourceMove() any {
	var returns any

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]any{id, provider},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) MoveTo(moveTarget *string, index any) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]any{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]any{id},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]any{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutAutoscaling(value *GoogleContainerNodePoolAutoscaling) {
	if err := g.validatePutAutoscalingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAutoscaling",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutManagement(value *GoogleContainerNodePoolManagement) {
	if err := g.validatePutManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManagement",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutNetworkConfig(value *GoogleContainerNodePoolNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutNodeConfig(value *GoogleContainerNodePoolNodeConfig) {
	if err := g.validatePutNodeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNodeConfig",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutPlacementPolicy(value *GoogleContainerNodePoolPlacementPolicy) {
	if err := g.validatePutPlacementPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPlacementPolicy",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutQueuedProvisioning(value *GoogleContainerNodePoolQueuedProvisioning) {
	if err := g.validatePutQueuedProvisioningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueuedProvisioning",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutTimeouts(value *GoogleContainerNodePoolTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) PutUpgradeSettings(value *GoogleContainerNodePoolUpgradeSettings) {
	if err := g.validatePutUpgradeSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putUpgradeSettings",
		[]any{value},
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetAutoscaling() {
	_jsii_.InvokeVoid(
		g,
		"resetAutoscaling",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetInitialNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetInitialNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetLocation() {
	_jsii_.InvokeVoid(
		g,
		"resetLocation",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetManagement() {
	_jsii_.InvokeVoid(
		g,
		"resetManagement",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		g,
		"resetMaxPodsPerNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetName() {
	_jsii_.InvokeVoid(
		g,
		"resetName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		g,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetNodeConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetNodeCount() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeCount",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetNodeLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetNodeLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetPlacementPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetPlacementPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetQueuedProvisioning() {
	_jsii_.InvokeVoid(
		g,
		"resetQueuedProvisioning",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetUpgradeSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetUpgradeSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) ResetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleContainerNodePool) SynthesizeAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) SynthesizeHclAttributes() *map[string]any {
	var returns *map[string]any

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) ToHclTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) ToMetadata() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleContainerNodePool) ToTerraform() any {
	var returns any

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
