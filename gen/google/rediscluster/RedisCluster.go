package rediscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google/rediscluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/redis_cluster google_redis_cluster}.
type RedisCluster interface {
	cdktf.TerraformResource
	AuthorizationMode() *string
	SetAuthorizationMode(val *string)
	AuthorizationModeInput() *string
	AutomatedBackupConfig() RedisClusterAutomatedBackupConfigOutputReference
	AutomatedBackupConfigInput() *RedisClusterAutomatedBackupConfig
	BackupCollection() *string
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
	CreateTime() *string
	CrossClusterReplicationConfig() RedisClusterCrossClusterReplicationConfigOutputReference
	CrossClusterReplicationConfigInput() *RedisClusterCrossClusterReplicationConfig
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiscoveryEndpoints() RedisClusterDiscoveryEndpointsList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GcsSource() RedisClusterGcsSourceOutputReference
	GcsSourceInput() *RedisClusterGcsSource
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenancePolicy() RedisClusterMaintenancePolicyOutputReference
	MaintenancePolicyInput() *RedisClusterMaintenancePolicy
	MaintenanceSchedule() RedisClusterMaintenanceScheduleList
	ManagedBackupSource() RedisClusterManagedBackupSourceOutputReference
	ManagedBackupSourceInput() *RedisClusterManagedBackupSource
	ManagedServerCa() RedisClusterManagedServerCaList
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	PersistenceConfig() RedisClusterPersistenceConfigOutputReference
	PersistenceConfigInput() *RedisClusterPersistenceConfig
	PreciseSizeGb() *float64
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
	PscConfigs() RedisClusterPscConfigsList
	PscConfigsInput() interface{}
	PscConnections() RedisClusterPscConnectionsList
	PscServiceAttachments() RedisClusterPscServiceAttachmentsList
	// Experimental.
	RawOverrides() interface{}
	RedisConfigs() *map[string]*string
	SetRedisConfigs(val *map[string]*string)
	RedisConfigsInput() *map[string]*string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	ReplicaCount() *float64
	SetReplicaCount(val *float64)
	ReplicaCountInput() *float64
	ShardCount() *float64
	SetShardCount(val *float64)
	ShardCountInput() *float64
	SizeGb() *float64
	State() *string
	StateInfo() RedisClusterStateInfoList
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RedisClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TransitEncryptionMode() *string
	SetTransitEncryptionMode(val *string)
	TransitEncryptionModeInput() *string
	Uid() *string
	ZoneDistributionConfig() RedisClusterZoneDistributionConfigOutputReference
	ZoneDistributionConfigInput() *RedisClusterZoneDistributionConfig
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
	PutAutomatedBackupConfig(value *RedisClusterAutomatedBackupConfig)
	PutCrossClusterReplicationConfig(value *RedisClusterCrossClusterReplicationConfig)
	PutGcsSource(value *RedisClusterGcsSource)
	PutMaintenancePolicy(value *RedisClusterMaintenancePolicy)
	PutManagedBackupSource(value *RedisClusterManagedBackupSource)
	PutPersistenceConfig(value *RedisClusterPersistenceConfig)
	PutPscConfigs(value interface{})
	PutTimeouts(value *RedisClusterTimeouts)
	PutZoneDistributionConfig(value *RedisClusterZoneDistributionConfig)
	ResetAuthorizationMode()
	ResetAutomatedBackupConfig()
	ResetCrossClusterReplicationConfig()
	ResetDeletionProtectionEnabled()
	ResetGcsSource()
	ResetId()
	ResetKmsKey()
	ResetMaintenancePolicy()
	ResetManagedBackupSource()
	ResetName()
	ResetNodeType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPersistenceConfig()
	ResetProject()
	ResetPscConfigs()
	ResetRedisConfigs()
	ResetRegion()
	ResetReplicaCount()
	ResetTimeouts()
	ResetTransitEncryptionMode()
	ResetZoneDistributionConfig()
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

// The jsii proxy struct for RedisCluster
type jsiiProxy_RedisCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RedisCluster) AuthorizationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) AuthorizationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) AutomatedBackupConfig() RedisClusterAutomatedBackupConfigOutputReference {
	var returns RedisClusterAutomatedBackupConfigOutputReference
	_jsii_.Get(
		j,
		"automatedBackupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) AutomatedBackupConfigInput() *RedisClusterAutomatedBackupConfig {
	var returns *RedisClusterAutomatedBackupConfig
	_jsii_.Get(
		j,
		"automatedBackupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) BackupCollection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupCollection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) CrossClusterReplicationConfig() RedisClusterCrossClusterReplicationConfigOutputReference {
	var returns RedisClusterCrossClusterReplicationConfigOutputReference
	_jsii_.Get(
		j,
		"crossClusterReplicationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) CrossClusterReplicationConfigInput() *RedisClusterCrossClusterReplicationConfig {
	var returns *RedisClusterCrossClusterReplicationConfig
	_jsii_.Get(
		j,
		"crossClusterReplicationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) DiscoveryEndpoints() RedisClusterDiscoveryEndpointsList {
	var returns RedisClusterDiscoveryEndpointsList
	_jsii_.Get(
		j,
		"discoveryEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) GcsSource() RedisClusterGcsSourceOutputReference {
	var returns RedisClusterGcsSourceOutputReference
	_jsii_.Get(
		j,
		"gcsSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) GcsSourceInput() *RedisClusterGcsSource {
	var returns *RedisClusterGcsSource
	_jsii_.Get(
		j,
		"gcsSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) MaintenancePolicy() RedisClusterMaintenancePolicyOutputReference {
	var returns RedisClusterMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) MaintenancePolicyInput() *RedisClusterMaintenancePolicy {
	var returns *RedisClusterMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) MaintenanceSchedule() RedisClusterMaintenanceScheduleList {
	var returns RedisClusterMaintenanceScheduleList
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ManagedBackupSource() RedisClusterManagedBackupSourceOutputReference {
	var returns RedisClusterManagedBackupSourceOutputReference
	_jsii_.Get(
		j,
		"managedBackupSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ManagedBackupSourceInput() *RedisClusterManagedBackupSource {
	var returns *RedisClusterManagedBackupSource
	_jsii_.Get(
		j,
		"managedBackupSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ManagedServerCa() RedisClusterManagedServerCaList {
	var returns RedisClusterManagedServerCaList
	_jsii_.Get(
		j,
		"managedServerCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) PersistenceConfig() RedisClusterPersistenceConfigOutputReference {
	var returns RedisClusterPersistenceConfigOutputReference
	_jsii_.Get(
		j,
		"persistenceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) PersistenceConfigInput() *RedisClusterPersistenceConfig {
	var returns *RedisClusterPersistenceConfig
	_jsii_.Get(
		j,
		"persistenceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) PreciseSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preciseSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) PscConfigs() RedisClusterPscConfigsList {
	var returns RedisClusterPscConfigsList
	_jsii_.Get(
		j,
		"pscConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) PscConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pscConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) PscConnections() RedisClusterPscConnectionsList {
	var returns RedisClusterPscConnectionsList
	_jsii_.Get(
		j,
		"pscConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) PscServiceAttachments() RedisClusterPscServiceAttachmentsList {
	var returns RedisClusterPscServiceAttachmentsList
	_jsii_.Get(
		j,
		"pscServiceAttachments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) RedisConfigs() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redisConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) RedisConfigsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"redisConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ReplicaCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ReplicaCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"replicaCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ShardCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ShardCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"shardCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) SizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) StateInfo() RedisClusterStateInfoList {
	var returns RedisClusterStateInfoList
	_jsii_.Get(
		j,
		"stateInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Timeouts() RedisClusterTimeoutsOutputReference {
	var returns RedisClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) TransitEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) TransitEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transitEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ZoneDistributionConfig() RedisClusterZoneDistributionConfigOutputReference {
	var returns RedisClusterZoneDistributionConfigOutputReference
	_jsii_.Get(
		j,
		"zoneDistributionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedisCluster) ZoneDistributionConfigInput() *RedisClusterZoneDistributionConfig {
	var returns *RedisClusterZoneDistributionConfig
	_jsii_.Get(
		j,
		"zoneDistributionConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/redis_cluster google_redis_cluster} Resource.
func NewRedisCluster(scope constructs.Construct, id *string, config *RedisClusterConfig) RedisCluster {
	_init_.Initialize()

	if err := validateNewRedisClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedisCluster{}

	_jsii_.Create(
		"@cdktf/provider-google.redisCluster.RedisCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/redis_cluster google_redis_cluster} Resource.
func NewRedisCluster_Override(r RedisCluster, scope constructs.Construct, id *string, config *RedisClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.redisCluster.RedisCluster",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedisCluster)SetAuthorizationMode(val *string) {
	if err := j.validateSetAuthorizationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationMode",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetRedisConfigs(val *map[string]*string) {
	if err := j.validateSetRedisConfigsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redisConfigs",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetReplicaCount(val *float64) {
	if err := j.validateSetReplicaCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaCount",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetShardCount(val *float64) {
	if err := j.validateSetShardCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shardCount",
		val,
	)
}

func (j *jsiiProxy_RedisCluster)SetTransitEncryptionMode(val *string) {
	if err := j.validateSetTransitEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionMode",
		val,
	)
}

// Generates CDKTF code for importing a RedisCluster resource upon running "cdktf plan <stack-name>".
func RedisCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRedisCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.redisCluster.RedisCluster",
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
func RedisCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedisCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.redisCluster.RedisCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedisCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedisCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.redisCluster.RedisCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedisCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedisCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google.redisCluster.RedisCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedisCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google.redisCluster.RedisCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedisCluster) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RedisCluster) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedisCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RedisCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedisCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RedisCluster) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RedisCluster) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedisCluster) PutAutomatedBackupConfig(value *RedisClusterAutomatedBackupConfig) {
	if err := r.validatePutAutomatedBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAutomatedBackupConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutCrossClusterReplicationConfig(value *RedisClusterCrossClusterReplicationConfig) {
	if err := r.validatePutCrossClusterReplicationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCrossClusterReplicationConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutGcsSource(value *RedisClusterGcsSource) {
	if err := r.validatePutGcsSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGcsSource",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutMaintenancePolicy(value *RedisClusterMaintenancePolicy) {
	if err := r.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutManagedBackupSource(value *RedisClusterManagedBackupSource) {
	if err := r.validatePutManagedBackupSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putManagedBackupSource",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutPersistenceConfig(value *RedisClusterPersistenceConfig) {
	if err := r.validatePutPersistenceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPersistenceConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutPscConfigs(value interface{}) {
	if err := r.validatePutPscConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPscConfigs",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutTimeouts(value *RedisClusterTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) PutZoneDistributionConfig(value *RedisClusterZoneDistributionConfig) {
	if err := r.validatePutZoneDistributionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putZoneDistributionConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedisCluster) ResetAuthorizationMode() {
	_jsii_.InvokeVoid(
		r,
		"resetAuthorizationMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetAutomatedBackupConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomatedBackupConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetCrossClusterReplicationConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetCrossClusterReplicationConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetGcsSource() {
	_jsii_.InvokeVoid(
		r,
		"resetGcsSource",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetKmsKey() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		r,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetManagedBackupSource() {
	_jsii_.InvokeVoid(
		r,
		"resetManagedBackupSource",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetName() {
	_jsii_.InvokeVoid(
		r,
		"resetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetNodeType() {
	_jsii_.InvokeVoid(
		r,
		"resetNodeType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetPersistenceConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetPersistenceConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetProject() {
	_jsii_.InvokeVoid(
		r,
		"resetProject",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetPscConfigs() {
	_jsii_.InvokeVoid(
		r,
		"resetPscConfigs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetRedisConfigs() {
	_jsii_.InvokeVoid(
		r,
		"resetRedisConfigs",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetReplicaCount() {
	_jsii_.InvokeVoid(
		r,
		"resetReplicaCount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetTransitEncryptionMode() {
	_jsii_.InvokeVoid(
		r,
		"resetTransitEncryptionMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) ResetZoneDistributionConfig() {
	_jsii_.InvokeVoid(
		r,
		"resetZoneDistributionConfig",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedisCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedisCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

