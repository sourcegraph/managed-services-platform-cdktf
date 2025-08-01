package googlenetappvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlenetappvolume/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_volume google_netapp_volume}.
type GoogleNetappVolume interface {
	cdktf.TerraformResource
	ActiveDirectory() *string
	BackupConfig() GoogleNetappVolumeBackupConfigOutputReference
	BackupConfigInput() *GoogleNetappVolumeBackupConfig
	CapacityGib() *string
	SetCapacityGib(val *string)
	CapacityGibInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ColdTierSizeGib() *string
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
	DeletionPolicy() *string
	SetDeletionPolicy(val *string)
	DeletionPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EffectiveLabels() cdktf.StringMap
	EncryptionType() *string
	ExportPolicy() GoogleNetappVolumeExportPolicyOutputReference
	ExportPolicyInput() *GoogleNetappVolumeExportPolicy
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HasReplication() cdktf.IResolvable
	HybridReplicationParameters() GoogleNetappVolumeHybridReplicationParametersOutputReference
	HybridReplicationParametersInput() *GoogleNetappVolumeHybridReplicationParameters
	Id() *string
	SetId(val *string)
	IdInput() *string
	KerberosEnabled() interface{}
	SetKerberosEnabled(val interface{})
	KerberosEnabledInput() interface{}
	KmsConfig() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LargeCapacity() interface{}
	SetLargeCapacity(val interface{})
	LargeCapacityInput() interface{}
	LdapEnabled() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MountOptions() GoogleNetappVolumeMountOptionsList
	MultipleEndpoints() interface{}
	SetMultipleEndpoints(val interface{})
	MultipleEndpointsInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Network() *string
	// The tree node.
	Node() constructs.Node
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocols() *[]*string
	SetProtocols(val *[]*string)
	ProtocolsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PsaRange() *string
	// Experimental.
	RawOverrides() interface{}
	ReplicaZone() *string
	RestoreParameters() GoogleNetappVolumeRestoreParametersOutputReference
	RestoreParametersInput() *GoogleNetappVolumeRestoreParameters
	RestrictedActions() *[]*string
	SetRestrictedActions(val *[]*string)
	RestrictedActionsInput() *[]*string
	SecurityStyle() *string
	SetSecurityStyle(val *string)
	SecurityStyleInput() *string
	ServiceLevel() *string
	ShareName() *string
	SetShareName(val *string)
	ShareNameInput() *string
	SmbSettings() *[]*string
	SetSmbSettings(val *[]*string)
	SmbSettingsInput() *[]*string
	SnapshotDirectory() interface{}
	SetSnapshotDirectory(val interface{})
	SnapshotDirectoryInput() interface{}
	SnapshotPolicy() GoogleNetappVolumeSnapshotPolicyOutputReference
	SnapshotPolicyInput() *GoogleNetappVolumeSnapshotPolicy
	State() *string
	StateDetails() *string
	StoragePool() *string
	SetStoragePool(val *string)
	StoragePoolInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TieringPolicy() GoogleNetappVolumeTieringPolicyOutputReference
	TieringPolicyInput() *GoogleNetappVolumeTieringPolicy
	Timeouts() GoogleNetappVolumeTimeoutsOutputReference
	TimeoutsInput() interface{}
	UnixPermissions() *string
	SetUnixPermissions(val *string)
	UnixPermissionsInput() *string
	UsedGib() *string
	Zone() *string
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
	PutBackupConfig(value *GoogleNetappVolumeBackupConfig)
	PutExportPolicy(value *GoogleNetappVolumeExportPolicy)
	PutHybridReplicationParameters(value *GoogleNetappVolumeHybridReplicationParameters)
	PutRestoreParameters(value *GoogleNetappVolumeRestoreParameters)
	PutSnapshotPolicy(value *GoogleNetappVolumeSnapshotPolicy)
	PutTieringPolicy(value *GoogleNetappVolumeTieringPolicy)
	PutTimeouts(value *GoogleNetappVolumeTimeouts)
	ResetBackupConfig()
	ResetDeletionPolicy()
	ResetDescription()
	ResetExportPolicy()
	ResetHybridReplicationParameters()
	ResetId()
	ResetKerberosEnabled()
	ResetLabels()
	ResetLargeCapacity()
	ResetMultipleEndpoints()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetRestoreParameters()
	ResetRestrictedActions()
	ResetSecurityStyle()
	ResetSmbSettings()
	ResetSnapshotDirectory()
	ResetSnapshotPolicy()
	ResetTieringPolicy()
	ResetTimeouts()
	ResetUnixPermissions()
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

// The jsii proxy struct for GoogleNetappVolume
type jsiiProxy_GoogleNetappVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleNetappVolume) ActiveDirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) BackupConfig() GoogleNetappVolumeBackupConfigOutputReference {
	var returns GoogleNetappVolumeBackupConfigOutputReference
	_jsii_.Get(
		j,
		"backupConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) BackupConfigInput() *GoogleNetappVolumeBackupConfig {
	var returns *GoogleNetappVolumeBackupConfig
	_jsii_.Get(
		j,
		"backupConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) CapacityGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) CapacityGibInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ColdTierSizeGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"coldTierSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) DeletionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) DeletionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) EncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ExportPolicy() GoogleNetappVolumeExportPolicyOutputReference {
	var returns GoogleNetappVolumeExportPolicyOutputReference
	_jsii_.Get(
		j,
		"exportPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ExportPolicyInput() *GoogleNetappVolumeExportPolicy {
	var returns *GoogleNetappVolumeExportPolicy
	_jsii_.Get(
		j,
		"exportPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) HasReplication() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"hasReplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) HybridReplicationParameters() GoogleNetappVolumeHybridReplicationParametersOutputReference {
	var returns GoogleNetappVolumeHybridReplicationParametersOutputReference
	_jsii_.Get(
		j,
		"hybridReplicationParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) HybridReplicationParametersInput() *GoogleNetappVolumeHybridReplicationParameters {
	var returns *GoogleNetappVolumeHybridReplicationParameters
	_jsii_.Get(
		j,
		"hybridReplicationParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) KerberosEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) KerberosEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kerberosEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) KmsConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) LargeCapacity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"largeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) LargeCapacityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"largeCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) LdapEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"ldapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) MountOptions() GoogleNetappVolumeMountOptionsList {
	var returns GoogleNetappVolumeMountOptionsList
	_jsii_.Get(
		j,
		"mountOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) MultipleEndpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) MultipleEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Network() *string {
	var returns *string
	_jsii_.Get(
		j,
		"network",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Protocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"protocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) PsaRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"psaRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ReplicaZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) RestoreParameters() GoogleNetappVolumeRestoreParametersOutputReference {
	var returns GoogleNetappVolumeRestoreParametersOutputReference
	_jsii_.Get(
		j,
		"restoreParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) RestoreParametersInput() *GoogleNetappVolumeRestoreParameters {
	var returns *GoogleNetappVolumeRestoreParameters
	_jsii_.Get(
		j,
		"restoreParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) RestrictedActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) RestrictedActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"restrictedActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ServiceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) ShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SmbSettings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"smbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SmbSettingsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"smbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SnapshotDirectory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SnapshotDirectoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SnapshotPolicy() GoogleNetappVolumeSnapshotPolicyOutputReference {
	var returns GoogleNetappVolumeSnapshotPolicyOutputReference
	_jsii_.Get(
		j,
		"snapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) SnapshotPolicyInput() *GoogleNetappVolumeSnapshotPolicy {
	var returns *GoogleNetappVolumeSnapshotPolicy
	_jsii_.Get(
		j,
		"snapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) StateDetails() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) StoragePool() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) StoragePoolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storagePoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) TieringPolicy() GoogleNetappVolumeTieringPolicyOutputReference {
	var returns GoogleNetappVolumeTieringPolicyOutputReference
	_jsii_.Get(
		j,
		"tieringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) TieringPolicyInput() *GoogleNetappVolumeTieringPolicy {
	var returns *GoogleNetappVolumeTieringPolicy
	_jsii_.Get(
		j,
		"tieringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Timeouts() GoogleNetappVolumeTimeoutsOutputReference {
	var returns GoogleNetappVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) UnixPermissions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unixPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) UnixPermissionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unixPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) UsedGib() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usedGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleNetappVolume) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_volume google_netapp_volume} Resource.
func NewGoogleNetappVolume(scope constructs.Construct, id *string, config *GoogleNetappVolumeConfig) GoogleNetappVolume {
	_init_.Initialize()

	if err := validateNewGoogleNetappVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleNetappVolume{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_volume google_netapp_volume} Resource.
func NewGoogleNetappVolume_Override(g GoogleNetappVolume, scope constructs.Construct, id *string, config *GoogleNetappVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolume",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetCapacityGib(val *string) {
	if err := j.validateSetCapacityGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityGib",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetDeletionPolicy(val *string) {
	if err := j.validateSetDeletionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetKerberosEnabled(val interface{}) {
	if err := j.validateSetKerberosEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetLargeCapacity(val interface{}) {
	if err := j.validateSetLargeCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"largeCapacity",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetMultipleEndpoints(val interface{}) {
	if err := j.validateSetMultipleEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multipleEndpoints",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetProtocols(val *[]*string) {
	if err := j.validateSetProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocols",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetRestrictedActions(val *[]*string) {
	if err := j.validateSetRestrictedActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedActions",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetShareName(val *string) {
	if err := j.validateSetShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareName",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetSmbSettings(val *[]*string) {
	if err := j.validateSetSmbSettingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbSettings",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetSnapshotDirectory(val interface{}) {
	if err := j.validateSetSnapshotDirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotDirectory",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetStoragePool(val *string) {
	if err := j.validateSetStoragePoolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storagePool",
		val,
	)
}

func (j *jsiiProxy_GoogleNetappVolume)SetUnixPermissions(val *string) {
	if err := j.validateSetUnixPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unixPermissions",
		val,
	)
}

// Generates CDKTF code for importing a GoogleNetappVolume resource upon running "cdktf plan <stack-name>".
func GoogleNetappVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleNetappVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolume",
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
func GoogleNetappVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetappVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetappVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetappVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleNetappVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleNetappVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleNetappVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleNetappVolume.GoogleNetappVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleNetappVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleNetappVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleNetappVolume) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleNetappVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleNetappVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleNetappVolume) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleNetappVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleNetappVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) PutBackupConfig(value *GoogleNetappVolumeBackupConfig) {
	if err := g.validatePutBackupConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackupConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) PutExportPolicy(value *GoogleNetappVolumeExportPolicy) {
	if err := g.validatePutExportPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExportPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) PutHybridReplicationParameters(value *GoogleNetappVolumeHybridReplicationParameters) {
	if err := g.validatePutHybridReplicationParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putHybridReplicationParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) PutRestoreParameters(value *GoogleNetappVolumeRestoreParameters) {
	if err := g.validatePutRestoreParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putRestoreParameters",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) PutSnapshotPolicy(value *GoogleNetappVolumeSnapshotPolicy) {
	if err := g.validatePutSnapshotPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSnapshotPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) PutTieringPolicy(value *GoogleNetappVolumeTieringPolicy) {
	if err := g.validatePutTieringPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTieringPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) PutTimeouts(value *GoogleNetappVolumeTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetBackupConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetBackupConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetDeletionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetDeletionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetExportPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetExportPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetHybridReplicationParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetHybridReplicationParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetKerberosEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetKerberosEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetLargeCapacity() {
	_jsii_.InvokeVoid(
		g,
		"resetLargeCapacity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetMultipleEndpoints() {
	_jsii_.InvokeVoid(
		g,
		"resetMultipleEndpoints",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetRestoreParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetRestoreParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetRestrictedActions() {
	_jsii_.InvokeVoid(
		g,
		"resetRestrictedActions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetSecurityStyle() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityStyle",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetSmbSettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSmbSettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetSnapshotDirectory() {
	_jsii_.InvokeVoid(
		g,
		"resetSnapshotDirectory",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetSnapshotPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSnapshotPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetTieringPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetTieringPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) ResetUnixPermissions() {
	_jsii_.InvokeVoid(
		g,
		"resetUnixPermissions",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleNetappVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleNetappVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

