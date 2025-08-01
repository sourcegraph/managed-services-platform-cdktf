package zerotrustdeviceposturerule

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRule",
		reflect.TypeOf((*ZeroTrustDevicePostureRule)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "expiration", GoGetter: "Expiration"},
			_jsii_.MemberProperty{JsiiProperty: "expirationInput", GoGetter: "ExpirationInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "input", GoGetter: "Input"},
			_jsii_.MemberProperty{JsiiProperty: "inputInput", GoGetter: "InputInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "match", GoGetter: "Match"},
			_jsii_.MemberProperty{JsiiProperty: "matchInput", GoGetter: "MatchInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putInput", GoMethod: "PutInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMatch", GoMethod: "PutMatch"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpiration", GoMethod: "ResetExpiration"},
			_jsii_.MemberMethod{JsiiMethod: "resetInput", GoMethod: "ResetInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMatch", GoMethod: "ResetMatch"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchedule", GoMethod: "ResetSchedule"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberProperty{JsiiProperty: "scheduleInput", GoGetter: "ScheduleInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_ZeroTrustDevicePostureRule{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleConfig",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleInput",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleInput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleInputLocations",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleInputLocations)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleInputLocationsOutputReference",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleInputLocationsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "paths", GoGetter: "Paths"},
			_jsii_.MemberProperty{JsiiProperty: "pathsInput", GoGetter: "PathsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPaths", GoMethod: "ResetPaths"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrustStores", GoMethod: "ResetTrustStores"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trustStores", GoGetter: "TrustStores"},
			_jsii_.MemberProperty{JsiiProperty: "trustStoresInput", GoGetter: "TrustStoresInput"},
		},
		func() interface{} {
			j := jsiiProxy_ZeroTrustDevicePostureRuleInputLocationsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleInputOutputReference",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleInputOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "activeThreats", GoGetter: "ActiveThreats"},
			_jsii_.MemberProperty{JsiiProperty: "activeThreatsInput", GoGetter: "ActiveThreatsInput"},
			_jsii_.MemberProperty{JsiiProperty: "certificateId", GoGetter: "CertificateId"},
			_jsii_.MemberProperty{JsiiProperty: "certificateIdInput", GoGetter: "CertificateIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "checkDisks", GoGetter: "CheckDisks"},
			_jsii_.MemberProperty{JsiiProperty: "checkDisksInput", GoGetter: "CheckDisksInput"},
			_jsii_.MemberProperty{JsiiProperty: "checkPrivateKey", GoGetter: "CheckPrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "checkPrivateKeyInput", GoGetter: "CheckPrivateKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "cn", GoGetter: "Cn"},
			_jsii_.MemberProperty{JsiiProperty: "cnInput", GoGetter: "CnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "complianceStatus", GoGetter: "ComplianceStatus"},
			_jsii_.MemberProperty{JsiiProperty: "complianceStatusInput", GoGetter: "ComplianceStatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "connectionId", GoGetter: "ConnectionId"},
			_jsii_.MemberProperty{JsiiProperty: "connectionIdInput", GoGetter: "ConnectionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "countOperator", GoGetter: "CountOperator"},
			_jsii_.MemberProperty{JsiiProperty: "countOperatorInput", GoGetter: "CountOperatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domain", GoGetter: "Domain"},
			_jsii_.MemberProperty{JsiiProperty: "domainInput", GoGetter: "DomainInput"},
			_jsii_.MemberProperty{JsiiProperty: "eidLastSeen", GoGetter: "EidLastSeen"},
			_jsii_.MemberProperty{JsiiProperty: "eidLastSeenInput", GoGetter: "EidLastSeenInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "exists", GoGetter: "Exists"},
			_jsii_.MemberProperty{JsiiProperty: "existsInput", GoGetter: "ExistsInput"},
			_jsii_.MemberProperty{JsiiProperty: "extendedKeyUsage", GoGetter: "ExtendedKeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "extendedKeyUsageInput", GoGetter: "ExtendedKeyUsageInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "infected", GoGetter: "Infected"},
			_jsii_.MemberProperty{JsiiProperty: "infectedInput", GoGetter: "InfectedInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isActive", GoGetter: "IsActive"},
			_jsii_.MemberProperty{JsiiProperty: "isActiveInput", GoGetter: "IsActiveInput"},
			_jsii_.MemberProperty{JsiiProperty: "issueCount", GoGetter: "IssueCount"},
			_jsii_.MemberProperty{JsiiProperty: "issueCountInput", GoGetter: "IssueCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "lastSeen", GoGetter: "LastSeen"},
			_jsii_.MemberProperty{JsiiProperty: "lastSeenInput", GoGetter: "LastSeenInput"},
			_jsii_.MemberProperty{JsiiProperty: "locations", GoGetter: "Locations"},
			_jsii_.MemberProperty{JsiiProperty: "locationsInput", GoGetter: "LocationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkStatus", GoGetter: "NetworkStatus"},
			_jsii_.MemberProperty{JsiiProperty: "networkStatusInput", GoGetter: "NetworkStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "operatingSystem", GoGetter: "OperatingSystem"},
			_jsii_.MemberProperty{JsiiProperty: "operatingSystemInput", GoGetter: "OperatingSystemInput"},
			_jsii_.MemberProperty{JsiiProperty: "operationalState", GoGetter: "OperationalState"},
			_jsii_.MemberProperty{JsiiProperty: "operationalStateInput", GoGetter: "OperationalStateInput"},
			_jsii_.MemberProperty{JsiiProperty: "operator", GoGetter: "Operator"},
			_jsii_.MemberProperty{JsiiProperty: "operatorInput", GoGetter: "OperatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "os", GoGetter: "Os"},
			_jsii_.MemberProperty{JsiiProperty: "osDistroName", GoGetter: "OsDistroName"},
			_jsii_.MemberProperty{JsiiProperty: "osDistroNameInput", GoGetter: "OsDistroNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "osDistroRevision", GoGetter: "OsDistroRevision"},
			_jsii_.MemberProperty{JsiiProperty: "osDistroRevisionInput", GoGetter: "OsDistroRevisionInput"},
			_jsii_.MemberProperty{JsiiProperty: "osInput", GoGetter: "OsInput"},
			_jsii_.MemberProperty{JsiiProperty: "osVersionExtra", GoGetter: "OsVersionExtra"},
			_jsii_.MemberProperty{JsiiProperty: "osVersionExtraInput", GoGetter: "OsVersionExtraInput"},
			_jsii_.MemberProperty{JsiiProperty: "overall", GoGetter: "Overall"},
			_jsii_.MemberProperty{JsiiProperty: "overallInput", GoGetter: "OverallInput"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLocations", GoMethod: "PutLocations"},
			_jsii_.MemberProperty{JsiiProperty: "requireAll", GoGetter: "RequireAll"},
			_jsii_.MemberProperty{JsiiProperty: "requireAllInput", GoGetter: "RequireAllInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetActiveThreats", GoMethod: "ResetActiveThreats"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertificateId", GoMethod: "ResetCertificateId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCheckDisks", GoMethod: "ResetCheckDisks"},
			_jsii_.MemberMethod{JsiiMethod: "resetCheckPrivateKey", GoMethod: "ResetCheckPrivateKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetCn", GoMethod: "ResetCn"},
			_jsii_.MemberMethod{JsiiMethod: "resetComplianceStatus", GoMethod: "ResetComplianceStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetConnectionId", GoMethod: "ResetConnectionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCountOperator", GoMethod: "ResetCountOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomain", GoMethod: "ResetDomain"},
			_jsii_.MemberMethod{JsiiMethod: "resetEidLastSeen", GoMethod: "ResetEidLastSeen"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetExists", GoMethod: "ResetExists"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtendedKeyUsage", GoMethod: "ResetExtendedKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInfected", GoMethod: "ResetInfected"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsActive", GoMethod: "ResetIsActive"},
			_jsii_.MemberMethod{JsiiMethod: "resetIssueCount", GoMethod: "ResetIssueCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetLastSeen", GoMethod: "ResetLastSeen"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocations", GoMethod: "ResetLocations"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkStatus", GoMethod: "ResetNetworkStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperatingSystem", GoMethod: "ResetOperatingSystem"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperationalState", GoMethod: "ResetOperationalState"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperator", GoMethod: "ResetOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetOs", GoMethod: "ResetOs"},
			_jsii_.MemberMethod{JsiiMethod: "resetOsDistroName", GoMethod: "ResetOsDistroName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOsDistroRevision", GoMethod: "ResetOsDistroRevision"},
			_jsii_.MemberMethod{JsiiMethod: "resetOsVersionExtra", GoMethod: "ResetOsVersionExtra"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverall", GoMethod: "ResetOverall"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireAll", GoMethod: "ResetRequireAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetRiskLevel", GoMethod: "ResetRiskLevel"},
			_jsii_.MemberMethod{JsiiMethod: "resetScore", GoMethod: "ResetScore"},
			_jsii_.MemberMethod{JsiiMethod: "resetScoreOperator", GoMethod: "ResetScoreOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resetSensorConfig", GoMethod: "ResetSensorConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSha256", GoMethod: "ResetSha256"},
			_jsii_.MemberMethod{JsiiMethod: "resetState", GoMethod: "ResetState"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubjectAlternativeNames", GoMethod: "ResetSubjectAlternativeNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetThumbprint", GoMethod: "ResetThumbprint"},
			_jsii_.MemberMethod{JsiiMethod: "resetTotalScore", GoMethod: "ResetTotalScore"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersion", GoMethod: "ResetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersionOperator", GoMethod: "ResetVersionOperator"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "riskLevel", GoGetter: "RiskLevel"},
			_jsii_.MemberProperty{JsiiProperty: "riskLevelInput", GoGetter: "RiskLevelInput"},
			_jsii_.MemberProperty{JsiiProperty: "score", GoGetter: "Score"},
			_jsii_.MemberProperty{JsiiProperty: "scoreInput", GoGetter: "ScoreInput"},
			_jsii_.MemberProperty{JsiiProperty: "scoreOperator", GoGetter: "ScoreOperator"},
			_jsii_.MemberProperty{JsiiProperty: "scoreOperatorInput", GoGetter: "ScoreOperatorInput"},
			_jsii_.MemberProperty{JsiiProperty: "sensorConfig", GoGetter: "SensorConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sensorConfigInput", GoGetter: "SensorConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "sha256", GoGetter: "Sha256"},
			_jsii_.MemberProperty{JsiiProperty: "sha256Input", GoGetter: "Sha256Input"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "stateInput", GoGetter: "StateInput"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNames", GoGetter: "SubjectAlternativeNames"},
			_jsii_.MemberProperty{JsiiProperty: "subjectAlternativeNamesInput", GoGetter: "SubjectAlternativeNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "thumbprint", GoGetter: "Thumbprint"},
			_jsii_.MemberProperty{JsiiProperty: "thumbprintInput", GoGetter: "ThumbprintInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "totalScore", GoGetter: "TotalScore"},
			_jsii_.MemberProperty{JsiiProperty: "totalScoreInput", GoGetter: "TotalScoreInput"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "versionOperator", GoGetter: "VersionOperator"},
			_jsii_.MemberProperty{JsiiProperty: "versionOperatorInput", GoGetter: "VersionOperatorInput"},
		},
		func() interface{} {
			j := jsiiProxy_ZeroTrustDevicePostureRuleInputOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleMatch",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleMatch)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleMatchList",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleMatchList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_ZeroTrustDevicePostureRuleMatchList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.zeroTrustDevicePostureRule.ZeroTrustDevicePostureRuleMatchOutputReference",
		reflect.TypeOf((*ZeroTrustDevicePostureRuleMatchOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "platform", GoGetter: "Platform"},
			_jsii_.MemberProperty{JsiiProperty: "platformInput", GoGetter: "PlatformInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlatform", GoMethod: "ResetPlatform"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ZeroTrustDevicePostureRuleMatchOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
