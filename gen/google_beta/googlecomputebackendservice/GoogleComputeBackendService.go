package googlecomputebackendservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/managed-services-platform-cdktf/gen/google_beta/googlecomputebackendservice/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_backend_service google_compute_backend_service}.
type GoogleComputeBackendService interface {
	cdktf.TerraformResource
	AffinityCookieTtlSec() *float64
	SetAffinityCookieTtlSec(val *float64)
	AffinityCookieTtlSecInput() *float64
	Backend() GoogleComputeBackendServiceBackendList
	BackendInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CdnPolicy() GoogleComputeBackendServiceCdnPolicyOutputReference
	CdnPolicyInput() *GoogleComputeBackendServiceCdnPolicy
	CircuitBreakers() GoogleComputeBackendServiceCircuitBreakersOutputReference
	CircuitBreakersInput() *GoogleComputeBackendServiceCircuitBreakers
	CompressionMode() *string
	SetCompressionMode(val *string)
	CompressionModeInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	ConnectionDrainingTimeoutSec() *float64
	SetConnectionDrainingTimeoutSec(val *float64)
	ConnectionDrainingTimeoutSecInput() *float64
	ConsistentHash() GoogleComputeBackendServiceConsistentHashOutputReference
	ConsistentHashInput() *GoogleComputeBackendServiceConsistentHash
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTimestamp() *string
	CustomRequestHeaders() *[]*string
	SetCustomRequestHeaders(val *[]*string)
	CustomRequestHeadersInput() *[]*string
	CustomResponseHeaders() *[]*string
	SetCustomResponseHeaders(val *[]*string)
	CustomResponseHeadersInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EdgeSecurityPolicy() *string
	SetEdgeSecurityPolicy(val *string)
	EdgeSecurityPolicyInput() *string
	EnableCdn() interface{}
	SetEnableCdn(val interface{})
	EnableCdnInput() interface{}
	Fingerprint() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeneratedId() *float64
	HealthChecks() *[]*string
	SetHealthChecks(val *[]*string)
	HealthChecksInput() *[]*string
	Iap() GoogleComputeBackendServiceIapOutputReference
	IapInput() *GoogleComputeBackendServiceIap
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpAddressSelectionPolicy() *string
	SetIpAddressSelectionPolicy(val *string)
	IpAddressSelectionPolicyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LoadBalancingScheme() *string
	SetLoadBalancingScheme(val *string)
	LoadBalancingSchemeInput() *string
	LocalityLbPolicies() GoogleComputeBackendServiceLocalityLbPoliciesList
	LocalityLbPoliciesInput() interface{}
	LocalityLbPolicy() *string
	SetLocalityLbPolicy(val *string)
	LocalityLbPolicyInput() *string
	LogConfig() GoogleComputeBackendServiceLogConfigOutputReference
	LogConfigInput() *GoogleComputeBackendServiceLogConfig
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutlierDetection() GoogleComputeBackendServiceOutlierDetectionOutputReference
	OutlierDetectionInput() *GoogleComputeBackendServiceOutlierDetection
	PortName() *string
	SetPortName(val *string)
	PortNameInput() *string
	Project() *string
	SetProject(val *string)
	ProjectInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	SecurityPolicy() *string
	SetSecurityPolicy(val *string)
	SecurityPolicyInput() *string
	SecuritySettings() GoogleComputeBackendServiceSecuritySettingsOutputReference
	SecuritySettingsInput() *GoogleComputeBackendServiceSecuritySettings
	SelfLink() *string
	ServiceLbPolicy() *string
	SetServiceLbPolicy(val *string)
	ServiceLbPolicyInput() *string
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityInput() *string
	StrongSessionAffinityCookie() GoogleComputeBackendServiceStrongSessionAffinityCookieOutputReference
	StrongSessionAffinityCookieInput() *GoogleComputeBackendServiceStrongSessionAffinityCookie
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GoogleComputeBackendServiceTimeoutsOutputReference
	TimeoutSec() *float64
	SetTimeoutSec(val *float64)
	TimeoutSecInput() *float64
	TimeoutsInput() interface{}
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
	PutBackend(value interface{})
	PutCdnPolicy(value *GoogleComputeBackendServiceCdnPolicy)
	PutCircuitBreakers(value *GoogleComputeBackendServiceCircuitBreakers)
	PutConsistentHash(value *GoogleComputeBackendServiceConsistentHash)
	PutIap(value *GoogleComputeBackendServiceIap)
	PutLocalityLbPolicies(value interface{})
	PutLogConfig(value *GoogleComputeBackendServiceLogConfig)
	PutOutlierDetection(value *GoogleComputeBackendServiceOutlierDetection)
	PutSecuritySettings(value *GoogleComputeBackendServiceSecuritySettings)
	PutStrongSessionAffinityCookie(value *GoogleComputeBackendServiceStrongSessionAffinityCookie)
	PutTimeouts(value *GoogleComputeBackendServiceTimeouts)
	ResetAffinityCookieTtlSec()
	ResetBackend()
	ResetCdnPolicy()
	ResetCircuitBreakers()
	ResetCompressionMode()
	ResetConnectionDrainingTimeoutSec()
	ResetConsistentHash()
	ResetCustomRequestHeaders()
	ResetCustomResponseHeaders()
	ResetDescription()
	ResetEdgeSecurityPolicy()
	ResetEnableCdn()
	ResetHealthChecks()
	ResetIap()
	ResetId()
	ResetIpAddressSelectionPolicy()
	ResetLoadBalancingScheme()
	ResetLocalityLbPolicies()
	ResetLocalityLbPolicy()
	ResetLogConfig()
	ResetOutlierDetection()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPortName()
	ResetProject()
	ResetProtocol()
	ResetSecurityPolicy()
	ResetSecuritySettings()
	ResetServiceLbPolicy()
	ResetSessionAffinity()
	ResetStrongSessionAffinityCookie()
	ResetTimeouts()
	ResetTimeoutSec()
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

// The jsii proxy struct for GoogleComputeBackendService
type jsiiProxy_GoogleComputeBackendService struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GoogleComputeBackendService) AffinityCookieTtlSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"affinityCookieTtlSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) AffinityCookieTtlSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"affinityCookieTtlSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Backend() GoogleComputeBackendServiceBackendList {
	var returns GoogleComputeBackendServiceBackendList
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) BackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CdnPolicy() GoogleComputeBackendServiceCdnPolicyOutputReference {
	var returns GoogleComputeBackendServiceCdnPolicyOutputReference
	_jsii_.Get(
		j,
		"cdnPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CdnPolicyInput() *GoogleComputeBackendServiceCdnPolicy {
	var returns *GoogleComputeBackendServiceCdnPolicy
	_jsii_.Get(
		j,
		"cdnPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CircuitBreakers() GoogleComputeBackendServiceCircuitBreakersOutputReference {
	var returns GoogleComputeBackendServiceCircuitBreakersOutputReference
	_jsii_.Get(
		j,
		"circuitBreakers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CircuitBreakersInput() *GoogleComputeBackendServiceCircuitBreakers {
	var returns *GoogleComputeBackendServiceCircuitBreakers
	_jsii_.Get(
		j,
		"circuitBreakersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CompressionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CompressionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ConnectionDrainingTimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ConnectionDrainingTimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionDrainingTimeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ConsistentHash() GoogleComputeBackendServiceConsistentHashOutputReference {
	var returns GoogleComputeBackendServiceConsistentHashOutputReference
	_jsii_.Get(
		j,
		"consistentHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ConsistentHashInput() *GoogleComputeBackendServiceConsistentHash {
	var returns *GoogleComputeBackendServiceConsistentHash
	_jsii_.Get(
		j,
		"consistentHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CreationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CustomRequestHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customRequestHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CustomRequestHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customRequestHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CustomResponseHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customResponseHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) CustomResponseHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"customResponseHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) EdgeSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) EdgeSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeSecurityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) EnableCdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) EnableCdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableCdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Fingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) GeneratedId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"generatedId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) HealthChecks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) HealthChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"healthChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Iap() GoogleComputeBackendServiceIapOutputReference {
	var returns GoogleComputeBackendServiceIapOutputReference
	_jsii_.Get(
		j,
		"iap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) IapInput() *GoogleComputeBackendServiceIap {
	var returns *GoogleComputeBackendServiceIap
	_jsii_.Get(
		j,
		"iapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) IpAddressSelectionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressSelectionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) IpAddressSelectionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressSelectionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LoadBalancingScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LoadBalancingSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LocalityLbPolicies() GoogleComputeBackendServiceLocalityLbPoliciesList {
	var returns GoogleComputeBackendServiceLocalityLbPoliciesList
	_jsii_.Get(
		j,
		"localityLbPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LocalityLbPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localityLbPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LocalityLbPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityLbPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LocalityLbPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityLbPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LogConfig() GoogleComputeBackendServiceLogConfigOutputReference {
	var returns GoogleComputeBackendServiceLogConfigOutputReference
	_jsii_.Get(
		j,
		"logConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) LogConfigInput() *GoogleComputeBackendServiceLogConfig {
	var returns *GoogleComputeBackendServiceLogConfig
	_jsii_.Get(
		j,
		"logConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) OutlierDetection() GoogleComputeBackendServiceOutlierDetectionOutputReference {
	var returns GoogleComputeBackendServiceOutlierDetectionOutputReference
	_jsii_.Get(
		j,
		"outlierDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) OutlierDetectionInput() *GoogleComputeBackendServiceOutlierDetection {
	var returns *GoogleComputeBackendServiceOutlierDetection
	_jsii_.Get(
		j,
		"outlierDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Project() *string {
	var returns *string
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ProjectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) SecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) SecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) SecuritySettings() GoogleComputeBackendServiceSecuritySettingsOutputReference {
	var returns GoogleComputeBackendServiceSecuritySettingsOutputReference
	_jsii_.Get(
		j,
		"securitySettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) SecuritySettingsInput() *GoogleComputeBackendServiceSecuritySettings {
	var returns *GoogleComputeBackendServiceSecuritySettings
	_jsii_.Get(
		j,
		"securitySettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) SelfLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selfLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ServiceLbPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLbPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) ServiceLbPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceLbPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) StrongSessionAffinityCookie() GoogleComputeBackendServiceStrongSessionAffinityCookieOutputReference {
	var returns GoogleComputeBackendServiceStrongSessionAffinityCookieOutputReference
	_jsii_.Get(
		j,
		"strongSessionAffinityCookie",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) StrongSessionAffinityCookieInput() *GoogleComputeBackendServiceStrongSessionAffinityCookie {
	var returns *GoogleComputeBackendServiceStrongSessionAffinityCookie
	_jsii_.Get(
		j,
		"strongSessionAffinityCookieInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) Timeouts() GoogleComputeBackendServiceTimeoutsOutputReference {
	var returns GoogleComputeBackendServiceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) TimeoutSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) TimeoutSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeBackendService) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_backend_service google_compute_backend_service} Resource.
func NewGoogleComputeBackendService(scope constructs.Construct, id *string, config *GoogleComputeBackendServiceConfig) GoogleComputeBackendService {
	_init_.Initialize()

	if err := validateNewGoogleComputeBackendServiceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeBackendService{}

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendService",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_compute_backend_service google_compute_backend_service} Resource.
func NewGoogleComputeBackendService_Override(g GoogleComputeBackendService, scope constructs.Construct, id *string, config *GoogleComputeBackendServiceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendService",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetAffinityCookieTtlSec(val *float64) {
	if err := j.validateSetAffinityCookieTtlSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affinityCookieTtlSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetCompressionMode(val *string) {
	if err := j.validateSetCompressionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionMode",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetConnectionDrainingTimeoutSec(val *float64) {
	if err := j.validateSetConnectionDrainingTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionDrainingTimeoutSec",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetCustomRequestHeaders(val *[]*string) {
	if err := j.validateSetCustomRequestHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customRequestHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetCustomResponseHeaders(val *[]*string) {
	if err := j.validateSetCustomResponseHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customResponseHeaders",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetEdgeSecurityPolicy(val *string) {
	if err := j.validateSetEdgeSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeSecurityPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetEnableCdn(val interface{}) {
	if err := j.validateSetEnableCdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCdn",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetHealthChecks(val *[]*string) {
	if err := j.validateSetHealthChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthChecks",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetIpAddressSelectionPolicy(val *string) {
	if err := j.validateSetIpAddressSelectionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressSelectionPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetLoadBalancingScheme(val *string) {
	if err := j.validateSetLoadBalancingSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancingScheme",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetLocalityLbPolicy(val *string) {
	if err := j.validateSetLocalityLbPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localityLbPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetProject(val *string) {
	if err := j.validateSetProjectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"project",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetSecurityPolicy(val *string) {
	if err := j.validateSetSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetServiceLbPolicy(val *string) {
	if err := j.validateSetServiceLbPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceLbPolicy",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeBackendService)SetTimeoutSec(val *float64) {
	if err := j.validateSetTimeoutSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSec",
		val,
	)
}

// Generates CDKTF code for importing a GoogleComputeBackendService resource upon running "cdktf plan <stack-name>".
func GoogleComputeBackendService_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGoogleComputeBackendService_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendService",
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
func GoogleComputeBackendService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeBackendService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeBackendService_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeBackendService_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendService",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GoogleComputeBackendService_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGoogleComputeBackendService_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendService",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GoogleComputeBackendService_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-google_beta.googleComputeBackendService.GoogleComputeBackendService",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GoogleComputeBackendService) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeBackendService) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeBackendService) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendService) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeBackendService) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutBackend(value interface{}) {
	if err := g.validatePutBackendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBackend",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutCdnPolicy(value *GoogleComputeBackendServiceCdnPolicy) {
	if err := g.validatePutCdnPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCdnPolicy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutCircuitBreakers(value *GoogleComputeBackendServiceCircuitBreakers) {
	if err := g.validatePutCircuitBreakersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCircuitBreakers",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutConsistentHash(value *GoogleComputeBackendServiceConsistentHash) {
	if err := g.validatePutConsistentHashParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putConsistentHash",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutIap(value *GoogleComputeBackendServiceIap) {
	if err := g.validatePutIapParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIap",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutLocalityLbPolicies(value interface{}) {
	if err := g.validatePutLocalityLbPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLocalityLbPolicies",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutLogConfig(value *GoogleComputeBackendServiceLogConfig) {
	if err := g.validatePutLogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putLogConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutOutlierDetection(value *GoogleComputeBackendServiceOutlierDetection) {
	if err := g.validatePutOutlierDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOutlierDetection",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutSecuritySettings(value *GoogleComputeBackendServiceSecuritySettings) {
	if err := g.validatePutSecuritySettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSecuritySettings",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutStrongSessionAffinityCookie(value *GoogleComputeBackendServiceStrongSessionAffinityCookie) {
	if err := g.validatePutStrongSessionAffinityCookieParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStrongSessionAffinityCookie",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) PutTimeouts(value *GoogleComputeBackendServiceTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetAffinityCookieTtlSec() {
	_jsii_.InvokeVoid(
		g,
		"resetAffinityCookieTtlSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetBackend() {
	_jsii_.InvokeVoid(
		g,
		"resetBackend",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetCdnPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetCdnPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetCircuitBreakers() {
	_jsii_.InvokeVoid(
		g,
		"resetCircuitBreakers",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetCompressionMode() {
	_jsii_.InvokeVoid(
		g,
		"resetCompressionMode",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetConnectionDrainingTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetConnectionDrainingTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetConsistentHash() {
	_jsii_.InvokeVoid(
		g,
		"resetConsistentHash",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetCustomRequestHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomRequestHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetCustomResponseHeaders() {
	_jsii_.InvokeVoid(
		g,
		"resetCustomResponseHeaders",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetEdgeSecurityPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetEdgeSecurityPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetEnableCdn() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableCdn",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetHealthChecks() {
	_jsii_.InvokeVoid(
		g,
		"resetHealthChecks",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetIap() {
	_jsii_.InvokeVoid(
		g,
		"resetIap",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetIpAddressSelectionPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetIpAddressSelectionPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetLoadBalancingScheme() {
	_jsii_.InvokeVoid(
		g,
		"resetLoadBalancingScheme",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetLocalityLbPolicies() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalityLbPolicies",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetLocalityLbPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetLocalityLbPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetLogConfig() {
	_jsii_.InvokeVoid(
		g,
		"resetLogConfig",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetOutlierDetection() {
	_jsii_.InvokeVoid(
		g,
		"resetOutlierDetection",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetPortName() {
	_jsii_.InvokeVoid(
		g,
		"resetPortName",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetProject() {
	_jsii_.InvokeVoid(
		g,
		"resetProject",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetProtocol() {
	_jsii_.InvokeVoid(
		g,
		"resetProtocol",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetSecurityPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetSecurityPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetSecuritySettings() {
	_jsii_.InvokeVoid(
		g,
		"resetSecuritySettings",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetServiceLbPolicy() {
	_jsii_.InvokeVoid(
		g,
		"resetServiceLbPolicy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		g,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetStrongSessionAffinityCookie() {
	_jsii_.InvokeVoid(
		g,
		"resetStrongSessionAffinityCookie",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) ResetTimeoutSec() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeoutSec",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeBackendService) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendService) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendService) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendService) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeBackendService) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

