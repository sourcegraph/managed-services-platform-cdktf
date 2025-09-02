package googlealloydbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlealloydbinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_alloydb_instance google_alloydb_instance}.
type GoogleAlloydbInstance interface {
	cdktf.TerraformResource
	ActivationPolicy() *string
	SetActivationPolicy(val *string)
	ActivationPolicyInput() *string
	Annotations() *map[string]*string
	SetAnnotations(val *map[string]*string)
	AnnotationsInput() *map[string]*string
	AvailabilityType() *string
	SetAvailabilityType(val *string)
	AvailabilityTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientConnectionConfig() GoogleAlloydbInstanceClientConnectionConfigOutputReference
	ClientConnectionConfigInput() *GoogleAlloydbInstanceClientConnectionConfig
	Cluster() *string
	SetCluster(val *string)
	ClusterInput() *string
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
	DatabaseFlags() *map[string]*string
	SetDatabaseFlags(val *map[string]*string)
	DatabaseFlagsInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	EffectiveAnnotations() cdktf.StringMap
	EffectiveLabels() cdktf.StringMap
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GceZone() *string
	SetGceZone(val *string)
	GceZoneInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InstanceId() *string
	SetInstanceId(val *string)
	InstanceIdInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	IpAddress() *string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MachineConfig() GoogleAlloydbInstanceMachineConfigOutputReference
	MachineConfigInput() *GoogleAlloydbInstanceMachineConfig
	Name() *string
	NetworkConfig() GoogleAlloydbInstanceNetworkConfigOutputReference
	NetworkConfigInput() *GoogleAlloydbInstanceNetworkConfig
	// The tree node.
	Node() constructs.Node
	ObservabilityConfig() GoogleAlloydbInstanceObservabilityConfigOutputReference
	ObservabilityConfigInput() *GoogleAlloydbInstanceObservabilityConfig
	OutboundPublicIpAddresses() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PscInstanceConfig() GoogleAlloydbInstancePscInstanceConfigOutputReference
	PscInstanceConfigInput() *GoogleAlloydbInstancePscInstanceConfig
	PublicIpAddress() *string
	QueryInsightsConfig() GoogleAlloydbInstanceQueryInsightsConfigOutputReference
	QueryInsightsConfigInput() *GoogleAlloydbInstanceQueryInsightsConfig
	// Experimental.
	RawOverrides() interface{}
	ReadPoolConfig() GoogleAlloydbInstanceReadPoolConfigOutputReference
	ReadPoolConfigInput() *GoogleAlloydbInstanceReadPoolConfig
	Reconciling() cdktf.IResolvable
	State() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleAlloydbInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uid() *string
	UpdateTime() *string
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
	PutClientConnectionConfig(value *GoogleAlloydbInstanceClientConnectionConfig)
	PutMachineConfig(value *GoogleAlloydbInstanceMachineConfig)
	PutNetworkConfig(value *GoogleAlloydbInstanceNetworkConfig)
	PutObservabilityConfig(value *GoogleAlloydbInstanceObservabilityConfig)
	PutPscInstanceConfig(value *GoogleAlloydbInstancePscInstanceConfig)
	PutQueryInsightsConfig(value *GoogleAlloydbInstanceQueryInsightsConfig)
	PutReadPoolConfig(value *GoogleAlloydbInstanceReadPoolConfig)
	PutTimeouts(value *GoogleAlloydbInstanceTimeouts)
	ResetActivationPolicy()
	ResetAnnotations()
	ResetAvailabilityType()
	ResetClientConnectionConfig()
	ResetDatabaseFlags()
	ResetDisplayName()
	ResetGceZone()
	ResetId()
	ResetLabels()
	ResetMachineConfig()
	ResetNetworkConfig()
	ResetObservabilityConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPscInstanceConfig()
	ResetQueryInsightsConfig()
	ResetReadPoolConfig()
	ResetTimeouts()
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

// The jsii proxy struct for GoogleAlloydbInstance
type jsiiProxy_GoogleAlloydbInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleAlloydbInstance) ActivationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ActivationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"activationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Annotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) AnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) AvailabilityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) AvailabilityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ClientConnectionConfig() GoogleAlloydbInstanceClientConnectionConfigOutputReference {
	var returns GoogleAlloydbInstanceClientConnectionConfigOutputReference
	_jsii_.Get(
		j,
		"clientConnectionConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ClientConnectionConfigInput() *GoogleAlloydbInstanceClientConnectionConfig {
	var returns *GoogleAlloydbInstanceClientConnectionConfig
	_jsii_.Get(
		j,
		"clientConnectionConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Cluster() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ClusterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DatabaseFlags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DatabaseFlagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"databaseFlagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) EffectiveAnnotations() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) GceZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gceZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) GceZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gceZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) IpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) MachineConfig() GoogleAlloydbInstanceMachineConfigOutputReference {
	var returns GoogleAlloydbInstanceMachineConfigOutputReference
	_jsii_.Get(
		j,
		"machineConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) MachineConfigInput() *GoogleAlloydbInstanceMachineConfig {
	var returns *GoogleAlloydbInstanceMachineConfig
	_jsii_.Get(
		j,
		"machineConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) NetworkConfig() GoogleAlloydbInstanceNetworkConfigOutputReference {
	var returns GoogleAlloydbInstanceNetworkConfigOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) NetworkConfigInput() *GoogleAlloydbInstanceNetworkConfig {
	var returns *GoogleAlloydbInstanceNetworkConfig
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ObservabilityConfig() GoogleAlloydbInstanceObservabilityConfigOutputReference {
	var returns GoogleAlloydbInstanceObservabilityConfigOutputReference
	_jsii_.Get(
		j,
		"observabilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ObservabilityConfigInput() *GoogleAlloydbInstanceObservabilityConfig {
	var returns *GoogleAlloydbInstanceObservabilityConfig
	_jsii_.Get(
		j,
		"observabilityConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) OutboundPublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundPublicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) PscInstanceConfig() GoogleAlloydbInstancePscInstanceConfigOutputReference {
	var returns GoogleAlloydbInstancePscInstanceConfigOutputReference
	_jsii_.Get(
		j,
		"pscInstanceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) PscInstanceConfigInput() *GoogleAlloydbInstancePscInstanceConfig {
	var returns *GoogleAlloydbInstancePscInstanceConfig
	_jsii_.Get(
		j,
		"pscInstanceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) QueryInsightsConfig() GoogleAlloydbInstanceQueryInsightsConfigOutputReference {
	var returns GoogleAlloydbInstanceQueryInsightsConfigOutputReference
	_jsii_.Get(
		j,
		"queryInsightsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) QueryInsightsConfigInput() *GoogleAlloydbInstanceQueryInsightsConfig {
	var returns *GoogleAlloydbInstanceQueryInsightsConfig
	_jsii_.Get(
		j,
		"queryInsightsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ReadPoolConfig() GoogleAlloydbInstanceReadPoolConfigOutputReference {
	var returns GoogleAlloydbInstanceReadPoolConfigOutputReference
	_jsii_.Get(
		j,
		"readPoolConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) ReadPoolConfigInput() *GoogleAlloydbInstanceReadPoolConfig {
	var returns *GoogleAlloydbInstanceReadPoolConfig
	_jsii_.Get(
		j,
		"readPoolConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Reconciling() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"reconciling",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Timeouts() GoogleAlloydbInstanceTimeoutsOutputReference {
	var returns GoogleAlloydbInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) Uid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleAlloydbInstance) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_alloydb_instance google_alloydb_instance} Resource.
func NewGoogleAlloydbInstance(scope constructs.Construct, id *string, config *GoogleAlloydbInstanceConfig) GoogleAlloydbInstance {
	_init_.Initialize()

	if err := validateNewGoogleAlloydbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleAlloydbInstance{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_alloydb_instance google_alloydb_instance} Resource.
func NewGoogleAlloydbInstance_Override(g GoogleAlloydbInstance, scope constructs.Construct, id *string, config *GoogleAlloydbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstance",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetActivationPolicy(val *string) {
	if err := j.validateSetActivationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activationPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetAnnotations(val *map[string]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetAvailabilityType(val *string) {
	if err := j.validateSetAvailabilityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityType",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetCluster(val *string) {
	if err := j.validateSetClusterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cluster",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetDatabaseFlags(val *map[string]*string) {
	if err := j.validateSetDatabaseFlagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseFlags",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetGceZone(val *string) {
	if err := j.validateSetGceZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gceZone",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetInstanceId(val *string) {
	if err := j.validateSetInstanceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceId",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleAlloydbInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a GoogleAlloydbInstance resource upon running "cdktf plan <stack-name>".
func GoogleAlloydbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstance",
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
func GoogleAlloydbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAlloydbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleAlloydbInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleAlloydbInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleAlloydbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleAlloydbInstance.GoogleAlloydbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleAlloydbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleAlloydbInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleAlloydbInstance) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutClientConnectionConfig(value *GoogleAlloydbInstanceClientConnectionConfig) {
	if err := g.validatePutClientConnectionConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putClientConnectionConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutMachineConfig(value *GoogleAlloydbInstanceMachineConfig) {
	if err := g.validatePutMachineConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMachineConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutNetworkConfig(value *GoogleAlloydbInstanceNetworkConfig) {
	if err := g.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutObservabilityConfig(value *GoogleAlloydbInstanceObservabilityConfig) {
	if err := g.validatePutObservabilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putObservabilityConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutPscInstanceConfig(value *GoogleAlloydbInstancePscInstanceConfig) {
	if err := g.validatePutPscInstanceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPscInstanceConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutQueryInsightsConfig(value *GoogleAlloydbInstanceQueryInsightsConfig) {
	if err := g.validatePutQueryInsightsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putQueryInsightsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutReadPoolConfig(value *GoogleAlloydbInstanceReadPoolConfig) {
	if err := g.validatePutReadPoolConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putReadPoolConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) PutTimeouts(value *GoogleAlloydbInstanceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetActivationPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetActivationPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetAnnotations() {
	_jsii_.InvokeVoid(
		g,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetAvailabilityType() {
	_jsii_.InvokeVoid(
		g,
		"resetAvailabilityType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetClientConnectionConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetClientConnectionConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetDatabaseFlags() {
	_jsii_.InvokeVoid(
		g,
		"resetDatabaseFlags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetDisplayName() {
	_jsii_.InvokeVoid(
		g,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetGceZone() {
	_jsii_.InvokeVoid(
		g,
		"resetGceZone",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetMachineConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetMachineConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetObservabilityConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetObservabilityConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetPscInstanceConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetPscInstanceConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetQueryInsightsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetQueryInsightsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetReadPoolConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetReadPoolConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleAlloydbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleAlloydbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

