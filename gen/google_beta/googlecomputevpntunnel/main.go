package googlecomputevpntunnel

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnel",
		reflect.TypeOf((*GoogleComputeVpnTunnel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cipherSuite", GoGetter: "CipherSuite"},
			_jsii_.MemberProperty{JsiiProperty: "cipherSuiteInput", GoGetter: "CipherSuiteInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationTimestamp", GoGetter: "CreationTimestamp"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "detailedStatus", GoGetter: "DetailedStatus"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveLabels", GoGetter: "EffectiveLabels"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ikeVersion", GoGetter: "IkeVersion"},
			_jsii_.MemberProperty{JsiiProperty: "ikeVersionInput", GoGetter: "IkeVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labelFingerprint", GoGetter: "LabelFingerprint"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "localTrafficSelector", GoGetter: "LocalTrafficSelector"},
			_jsii_.MemberProperty{JsiiProperty: "localTrafficSelectorInput", GoGetter: "LocalTrafficSelectorInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "peerExternalGateway", GoGetter: "PeerExternalGateway"},
			_jsii_.MemberProperty{JsiiProperty: "peerExternalGatewayInput", GoGetter: "PeerExternalGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "peerExternalGatewayInterface", GoGetter: "PeerExternalGatewayInterface"},
			_jsii_.MemberProperty{JsiiProperty: "peerExternalGatewayInterfaceInput", GoGetter: "PeerExternalGatewayInterfaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "peerGcpGateway", GoGetter: "PeerGcpGateway"},
			_jsii_.MemberProperty{JsiiProperty: "peerGcpGatewayInput", GoGetter: "PeerGcpGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "peerIp", GoGetter: "PeerIp"},
			_jsii_.MemberProperty{JsiiProperty: "peerIpInput", GoGetter: "PeerIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putCipherSuite", GoMethod: "PutCipherSuite"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberProperty{JsiiProperty: "remoteTrafficSelector", GoGetter: "RemoteTrafficSelector"},
			_jsii_.MemberProperty{JsiiProperty: "remoteTrafficSelectorInput", GoGetter: "RemoteTrafficSelectorInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCipherSuite", GoMethod: "ResetCipherSuite"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIkeVersion", GoMethod: "ResetIkeVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocalTrafficSelector", GoMethod: "ResetLocalTrafficSelector"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeerExternalGateway", GoMethod: "ResetPeerExternalGateway"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeerExternalGatewayInterface", GoMethod: "ResetPeerExternalGatewayInterface"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeerGcpGateway", GoMethod: "ResetPeerGcpGateway"},
			_jsii_.MemberMethod{JsiiMethod: "resetPeerIp", GoMethod: "ResetPeerIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemoteTrafficSelector", GoMethod: "ResetRemoteTrafficSelector"},
			_jsii_.MemberMethod{JsiiMethod: "resetRouter", GoMethod: "ResetRouter"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetVpnGateway", GoMethod: "ResetTargetVpnGateway"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpnGateway", GoMethod: "ResetVpnGateway"},
			_jsii_.MemberMethod{JsiiMethod: "resetVpnGatewayInterface", GoMethod: "ResetVpnGatewayInterface"},
			_jsii_.MemberProperty{JsiiProperty: "router", GoGetter: "Router"},
			_jsii_.MemberProperty{JsiiProperty: "routerInput", GoGetter: "RouterInput"},
			_jsii_.MemberProperty{JsiiProperty: "selfLink", GoGetter: "SelfLink"},
			_jsii_.MemberProperty{JsiiProperty: "sharedSecret", GoGetter: "SharedSecret"},
			_jsii_.MemberProperty{JsiiProperty: "sharedSecretHash", GoGetter: "SharedSecretHash"},
			_jsii_.MemberProperty{JsiiProperty: "sharedSecretInput", GoGetter: "SharedSecretInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetVpnGateway", GoGetter: "TargetVpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "targetVpnGatewayInput", GoGetter: "TargetVpnGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformLabels", GoGetter: "TerraformLabels"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelId", GoGetter: "TunnelId"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGateway", GoGetter: "VpnGateway"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayInput", GoGetter: "VpnGatewayInput"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayInterface", GoGetter: "VpnGatewayInterface"},
			_jsii_.MemberProperty{JsiiProperty: "vpnGatewayInterfaceInput", GoGetter: "VpnGatewayInterfaceInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleComputeVpnTunnel{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuite",
		reflect.TypeOf((*GoogleComputeVpnTunnelCipherSuite)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuiteOutputReference",
		reflect.TypeOf((*GoogleComputeVpnTunnelCipherSuiteOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "phase1", GoGetter: "Phase1"},
			_jsii_.MemberProperty{JsiiProperty: "phase1Input", GoGetter: "Phase1Input"},
			_jsii_.MemberProperty{JsiiProperty: "phase2", GoGetter: "Phase2"},
			_jsii_.MemberProperty{JsiiProperty: "phase2Input", GoGetter: "Phase2Input"},
			_jsii_.MemberMethod{JsiiMethod: "putPhase1", GoMethod: "PutPhase1"},
			_jsii_.MemberMethod{JsiiMethod: "putPhase2", GoMethod: "PutPhase2"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase1", GoMethod: "ResetPhase1"},
			_jsii_.MemberMethod{JsiiMethod: "resetPhase2", GoMethod: "ResetPhase2"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleComputeVpnTunnelCipherSuiteOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuitePhase1",
		reflect.TypeOf((*GoogleComputeVpnTunnelCipherSuitePhase1)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuitePhase1OutputReference",
		reflect.TypeOf((*GoogleComputeVpnTunnelCipherSuitePhase1OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dh", GoGetter: "Dh"},
			_jsii_.MemberProperty{JsiiProperty: "dhInput", GoGetter: "DhInput"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionInput", GoGetter: "EncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "integrity", GoGetter: "Integrity"},
			_jsii_.MemberProperty{JsiiProperty: "integrityInput", GoGetter: "IntegrityInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "prf", GoGetter: "Prf"},
			_jsii_.MemberProperty{JsiiProperty: "prfInput", GoGetter: "PrfInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDh", GoMethod: "ResetDh"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryption", GoMethod: "ResetEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrity", GoMethod: "ResetIntegrity"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrf", GoMethod: "ResetPrf"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleComputeVpnTunnelCipherSuitePhase1OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuitePhase2",
		reflect.TypeOf((*GoogleComputeVpnTunnelCipherSuitePhase2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelCipherSuitePhase2OutputReference",
		reflect.TypeOf((*GoogleComputeVpnTunnelCipherSuitePhase2OutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryption", GoGetter: "Encryption"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionInput", GoGetter: "EncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "integrity", GoGetter: "Integrity"},
			_jsii_.MemberProperty{JsiiProperty: "integrityInput", GoGetter: "IntegrityInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pfs", GoGetter: "Pfs"},
			_jsii_.MemberProperty{JsiiProperty: "pfsInput", GoGetter: "PfsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryption", GoMethod: "ResetEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegrity", GoMethod: "ResetIntegrity"},
			_jsii_.MemberMethod{JsiiMethod: "resetPfs", GoMethod: "ResetPfs"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleComputeVpnTunnelCipherSuitePhase2OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelConfig",
		reflect.TypeOf((*GoogleComputeVpnTunnelConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelTimeouts",
		reflect.TypeOf((*GoogleComputeVpnTunnelTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google_beta.googleComputeVpnTunnel.GoogleComputeVpnTunnelTimeoutsOutputReference",
		reflect.TypeOf((*GoogleComputeVpnTunnelTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleComputeVpnTunnelTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
