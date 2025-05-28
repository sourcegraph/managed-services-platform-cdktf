package googleedgecontainercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googleedgecontainercluster/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_edgecontainer_cluster google_edgecontainer_cluster}.
type GoogleEdgecontainerCluster interface {
	cdktf.TerraformResource
	Authorization() GoogleEdgecontainerClusterAuthorizationOutputReference
	AuthorizationInput() *GoogleEdgecontainerClusterAuthorization
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterCaCertificate() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPlane() GoogleEdgecontainerClusterControlPlaneOutputReference
	ControlPlaneEncryption() GoogleEdgecontainerClusterControlPlaneEncryptionOutputReference
	ControlPlaneEncryptionInput() *GoogleEdgecontainerClusterControlPlaneEncryption
	ControlPlaneInput() *GoogleEdgecontainerClusterControlPlane
	ControlPlaneVersion() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreateTime() *string
	DefaultMaxPodsPerNode() *float64
	SetDefaultMaxPodsPerNode(val *float64)
	DefaultMaxPodsPerNodeInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EffectiveLabels() cdktf.StringMap
	Endpoint() *string
	ExternalLoadBalancerIpv4AddressPools() *[]*string
	SetExternalLoadBalancerIpv4AddressPools(val *[]*string)
	ExternalLoadBalancerIpv4AddressPoolsInput() *[]*string
	Fleet() GoogleEdgecontainerClusterFleetOutputReference
	FleetInput() *GoogleEdgecontainerClusterFleet
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
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaintenanceEvents() GoogleEdgecontainerClusterMaintenanceEventsList
	MaintenancePolicy() GoogleEdgecontainerClusterMaintenancePolicyOutputReference
	MaintenancePolicyInput() *GoogleEdgecontainerClusterMaintenancePolicy
	Name() *string
	SetName(val *string)
	NameInput() *string
	Networking() GoogleEdgecontainerClusterNetworkingOutputReference
	NetworkingInput() *GoogleEdgecontainerClusterNetworking
	// The tree node.
	Node() constructs.Node
	NodeVersion() *string
	Port() *float64
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
	ReleaseChannel() *string
	SetReleaseChannel(val *string)
	ReleaseChannelInput() *string
	Status() *string
	SystemAddonsConfig() GoogleEdgecontainerClusterSystemAddonsConfigOutputReference
	SystemAddonsConfigInput() *GoogleEdgecontainerClusterSystemAddonsConfig
	TargetVersion() *string
	SetTargetVersion(val *string)
	TargetVersionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	TerraformLabels() cdktf.StringMap
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleEdgecontainerClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAuthorization(value *GoogleEdgecontainerClusterAuthorization)
	PutControlPlane(value *GoogleEdgecontainerClusterControlPlane)
	PutControlPlaneEncryption(value *GoogleEdgecontainerClusterControlPlaneEncryption)
	PutFleet(value *GoogleEdgecontainerClusterFleet)
	PutMaintenancePolicy(value *GoogleEdgecontainerClusterMaintenancePolicy)
	PutNetworking(value *GoogleEdgecontainerClusterNetworking)
	PutSystemAddonsConfig(value *GoogleEdgecontainerClusterSystemAddonsConfig)
	PutTimeouts(value *GoogleEdgecontainerClusterTimeouts)
	ResetControlPlane()
	ResetControlPlaneEncryption()
	ResetDefaultMaxPodsPerNode()
	ResetExternalLoadBalancerIpv4AddressPools()
	ResetId()
	ResetLabels()
	ResetMaintenancePolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProject()
	ResetReleaseChannel()
	ResetSystemAddonsConfig()
	ResetTargetVersion()
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

// The jsii proxy struct for GoogleEdgecontainerCluster
type jsiiProxy_GoogleEdgecontainerCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Authorization() GoogleEdgecontainerClusterAuthorizationOutputReference {
	var returns GoogleEdgecontainerClusterAuthorizationOutputReference
	_jsii_.Get(
		j,
		"authorization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) AuthorizationInput() *GoogleEdgecontainerClusterAuthorization {
	var returns *GoogleEdgecontainerClusterAuthorization
	_jsii_.Get(
		j,
		"authorizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ClusterCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ControlPlane() GoogleEdgecontainerClusterControlPlaneOutputReference {
	var returns GoogleEdgecontainerClusterControlPlaneOutputReference
	_jsii_.Get(
		j,
		"controlPlane",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ControlPlaneEncryption() GoogleEdgecontainerClusterControlPlaneEncryptionOutputReference {
	var returns GoogleEdgecontainerClusterControlPlaneEncryptionOutputReference
	_jsii_.Get(
		j,
		"controlPlaneEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ControlPlaneEncryptionInput() *GoogleEdgecontainerClusterControlPlaneEncryption {
	var returns *GoogleEdgecontainerClusterControlPlaneEncryption
	_jsii_.Get(
		j,
		"controlPlaneEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ControlPlaneInput() *GoogleEdgecontainerClusterControlPlane {
	var returns *GoogleEdgecontainerClusterControlPlane
	_jsii_.Get(
		j,
		"controlPlaneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ControlPlaneVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) CreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) DefaultMaxPodsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) DefaultMaxPodsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultMaxPodsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) EffectiveLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"effectiveLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ExternalLoadBalancerIpv4AddressPools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalLoadBalancerIpv4AddressPools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ExternalLoadBalancerIpv4AddressPoolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalLoadBalancerIpv4AddressPoolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Fleet() GoogleEdgecontainerClusterFleetOutputReference {
	var returns GoogleEdgecontainerClusterFleetOutputReference
	_jsii_.Get(
		j,
		"fleet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) FleetInput() *GoogleEdgecontainerClusterFleet {
	var returns *GoogleEdgecontainerClusterFleet
	_jsii_.Get(
		j,
		"fleetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) MaintenanceEvents() GoogleEdgecontainerClusterMaintenanceEventsList {
	var returns GoogleEdgecontainerClusterMaintenanceEventsList
	_jsii_.Get(
		j,
		"maintenanceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) MaintenancePolicy() GoogleEdgecontainerClusterMaintenancePolicyOutputReference {
	var returns GoogleEdgecontainerClusterMaintenancePolicyOutputReference
	_jsii_.Get(
		j,
		"maintenancePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) MaintenancePolicyInput() *GoogleEdgecontainerClusterMaintenancePolicy {
	var returns *GoogleEdgecontainerClusterMaintenancePolicy
	_jsii_.Get(
		j,
		"maintenancePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Networking() GoogleEdgecontainerClusterNetworkingOutputReference {
	var returns GoogleEdgecontainerClusterNetworkingOutputReference
	_jsii_.Get(
		j,
		"networking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) NetworkingInput() *GoogleEdgecontainerClusterNetworking {
	var returns *GoogleEdgecontainerClusterNetworking
	_jsii_.Get(
		j,
		"networkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) NodeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ReleaseChannel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) ReleaseChannelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseChannelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) SystemAddonsConfig() GoogleEdgecontainerClusterSystemAddonsConfigOutputReference {
	var returns GoogleEdgecontainerClusterSystemAddonsConfigOutputReference
	_jsii_.Get(
		j,
		"systemAddonsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) SystemAddonsConfigInput() *GoogleEdgecontainerClusterSystemAddonsConfig {
	var returns *GoogleEdgecontainerClusterSystemAddonsConfig
	_jsii_.Get(
		j,
		"systemAddonsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) TargetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) TargetVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) TerraformLabels() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"terraformLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) Timeouts() GoogleEdgecontainerClusterTimeoutsOutputReference {
	var returns GoogleEdgecontainerClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEdgecontainerCluster) UpdateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateTime",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_edgecontainer_cluster google_edgecontainer_cluster} Resource.
func NewGoogleEdgecontainerCluster(scope constructs.Construct, id *string, config *GoogleEdgecontainerClusterConfig) GoogleEdgecontainerCluster {
	_init_.Initialize()

	if err := validateNewGoogleEdgecontainerClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEdgecontainerCluster{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleEdgecontainerCluster.GoogleEdgecontainerCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_edgecontainer_cluster google_edgecontainer_cluster} Resource.
func NewGoogleEdgecontainerCluster_Override(g GoogleEdgecontainerCluster, scope constructs.Construct, id *string, config *GoogleEdgecontainerClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleEdgecontainerCluster.GoogleEdgecontainerCluster",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetDefaultMaxPodsPerNode(val *float64) {
	if err := j.validateSetDefaultMaxPodsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultMaxPodsPerNode",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetExternalLoadBalancerIpv4AddressPools(val *[]*string) {
	if err := j.validateSetExternalLoadBalancerIpv4AddressPoolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalLoadBalancerIpv4AddressPools",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetReleaseChannel(val *string) {
	if err := j.validateSetReleaseChannelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseChannel",
		val,
	)
}

func (j *jsiiProxy_GoogleEdgecontainerCluster)SetTargetVersion(val *string) {
	if err := j.validateSetTargetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetVersion",
		val,
	)
}

// Generates CDKTF code for importing a GoogleEdgecontainerCluster resource upon running "cdktf plan <stack-name>".
func GoogleEdgecontainerCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleEdgecontainerCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleEdgecontainerCluster.GoogleEdgecontainerCluster",
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
func GoogleEdgecontainerCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleEdgecontainerCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleEdgecontainerCluster.GoogleEdgecontainerCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleEdgecontainerCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleEdgecontainerCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleEdgecontainerCluster.GoogleEdgecontainerCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleEdgecontainerCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleEdgecontainerCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleEdgecontainerCluster.GoogleEdgecontainerCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleEdgecontainerCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleEdgecontainerCluster.GoogleEdgecontainerCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEdgecontainerCluster) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutAuthorization(value *GoogleEdgecontainerClusterAuthorization) {
	if err := g.validatePutAuthorizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putAuthorization",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutControlPlane(value *GoogleEdgecontainerClusterControlPlane) {
	if err := g.validatePutControlPlaneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlane",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutControlPlaneEncryption(value *GoogleEdgecontainerClusterControlPlaneEncryption) {
	if err := g.validatePutControlPlaneEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putControlPlaneEncryption",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutFleet(value *GoogleEdgecontainerClusterFleet) {
	if err := g.validatePutFleetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFleet",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutMaintenancePolicy(value *GoogleEdgecontainerClusterMaintenancePolicy) {
	if err := g.validatePutMaintenancePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMaintenancePolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutNetworking(value *GoogleEdgecontainerClusterNetworking) {
	if err := g.validatePutNetworkingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNetworking",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutSystemAddonsConfig(value *GoogleEdgecontainerClusterSystemAddonsConfig) {
	if err := g.validatePutSystemAddonsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSystemAddonsConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) PutTimeouts(value *GoogleEdgecontainerClusterTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetControlPlane() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlane",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetControlPlaneEncryption() {
	_jsii_.InvokeVoid(
		g,
		"resetControlPlaneEncryption",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetDefaultMaxPodsPerNode() {
	_jsii_.InvokeVoid(
		g,
		"resetDefaultMaxPodsPerNode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetExternalLoadBalancerIpv4AddressPools() {
	_jsii_.InvokeVoid(
		g,
		"resetExternalLoadBalancerIpv4AddressPools",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetMaintenancePolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetMaintenancePolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetReleaseChannel() {
	_jsii_.InvokeVoid(
		g,
		"resetReleaseChannel",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetSystemAddonsConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetSystemAddonsConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetTargetVersion() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetVersion",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEdgecontainerCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

