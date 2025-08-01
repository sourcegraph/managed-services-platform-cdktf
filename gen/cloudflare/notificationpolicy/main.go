package notificationpolicy

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicy",
		reflect.TypeOf((*NotificationPolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alertInterval", GoGetter: "AlertInterval"},
			_jsii_.MemberProperty{JsiiProperty: "alertIntervalInput", GoGetter: "AlertIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "alertType", GoGetter: "AlertType"},
			_jsii_.MemberProperty{JsiiProperty: "alertTypeInput", GoGetter: "AlertTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "created", GoGetter: "Created"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "filters", GoGetter: "Filters"},
			_jsii_.MemberProperty{JsiiProperty: "filtersInput", GoGetter: "FiltersInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mechanisms", GoGetter: "Mechanisms"},
			_jsii_.MemberProperty{JsiiProperty: "mechanismsInput", GoGetter: "MechanismsInput"},
			_jsii_.MemberProperty{JsiiProperty: "modified", GoGetter: "Modified"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putFilters", GoMethod: "PutFilters"},
			_jsii_.MemberMethod{JsiiMethod: "putMechanisms", GoMethod: "PutMechanisms"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlertInterval", GoMethod: "ResetAlertInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilters", GoMethod: "ResetFilters"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_NotificationPolicy{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyConfig",
		reflect.TypeOf((*NotificationPolicyConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyFilters",
		reflect.TypeOf((*NotificationPolicyFilters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyFiltersOutputReference",
		reflect.TypeOf((*NotificationPolicyFiltersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberProperty{JsiiProperty: "actionsInput", GoGetter: "ActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "affectedAsns", GoGetter: "AffectedAsns"},
			_jsii_.MemberProperty{JsiiProperty: "affectedAsnsInput", GoGetter: "AffectedAsnsInput"},
			_jsii_.MemberProperty{JsiiProperty: "affectedComponents", GoGetter: "AffectedComponents"},
			_jsii_.MemberProperty{JsiiProperty: "affectedComponentsInput", GoGetter: "AffectedComponentsInput"},
			_jsii_.MemberProperty{JsiiProperty: "affectedLocations", GoGetter: "AffectedLocations"},
			_jsii_.MemberProperty{JsiiProperty: "affectedLocationsInput", GoGetter: "AffectedLocationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "airportCode", GoGetter: "AirportCode"},
			_jsii_.MemberProperty{JsiiProperty: "airportCodeInput", GoGetter: "AirportCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "alertTriggerPreferences", GoGetter: "AlertTriggerPreferences"},
			_jsii_.MemberProperty{JsiiProperty: "alertTriggerPreferencesInput", GoGetter: "AlertTriggerPreferencesInput"},
			_jsii_.MemberProperty{JsiiProperty: "alertTriggerPreferencesValue", GoGetter: "AlertTriggerPreferencesValue"},
			_jsii_.MemberProperty{JsiiProperty: "alertTriggerPreferencesValueInput", GoGetter: "AlertTriggerPreferencesValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberProperty{JsiiProperty: "environmentInput", GoGetter: "EnvironmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "event", GoGetter: "Event"},
			_jsii_.MemberProperty{JsiiProperty: "eventInput", GoGetter: "EventInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventSource", GoGetter: "EventSource"},
			_jsii_.MemberProperty{JsiiProperty: "eventSourceInput", GoGetter: "EventSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypeInput", GoGetter: "EventTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "groupBy", GoGetter: "GroupBy"},
			_jsii_.MemberProperty{JsiiProperty: "groupByInput", GoGetter: "GroupByInput"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckId", GoGetter: "HealthCheckId"},
			_jsii_.MemberProperty{JsiiProperty: "healthCheckIdInput", GoGetter: "HealthCheckIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "incidentImpact", GoGetter: "IncidentImpact"},
			_jsii_.MemberProperty{JsiiProperty: "incidentImpactInput", GoGetter: "IncidentImpactInput"},
			_jsii_.MemberProperty{JsiiProperty: "inputId", GoGetter: "InputId"},
			_jsii_.MemberProperty{JsiiProperty: "inputIdInput", GoGetter: "InputIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "insightClass", GoGetter: "InsightClass"},
			_jsii_.MemberProperty{JsiiProperty: "insightClassInput", GoGetter: "InsightClassInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "limit", GoGetter: "Limit"},
			_jsii_.MemberProperty{JsiiProperty: "limitInput", GoGetter: "LimitInput"},
			_jsii_.MemberProperty{JsiiProperty: "logoTag", GoGetter: "LogoTag"},
			_jsii_.MemberProperty{JsiiProperty: "logoTagInput", GoGetter: "LogoTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "megabitsPerSecond", GoGetter: "MegabitsPerSecond"},
			_jsii_.MemberProperty{JsiiProperty: "megabitsPerSecondInput", GoGetter: "MegabitsPerSecondInput"},
			_jsii_.MemberProperty{JsiiProperty: "newHealth", GoGetter: "NewHealth"},
			_jsii_.MemberProperty{JsiiProperty: "newHealthInput", GoGetter: "NewHealthInput"},
			_jsii_.MemberProperty{JsiiProperty: "newStatus", GoGetter: "NewStatus"},
			_jsii_.MemberProperty{JsiiProperty: "newStatusInput", GoGetter: "NewStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "packetsPerSecond", GoGetter: "PacketsPerSecond"},
			_jsii_.MemberProperty{JsiiProperty: "packetsPerSecondInput", GoGetter: "PacketsPerSecondInput"},
			_jsii_.MemberProperty{JsiiProperty: "poolId", GoGetter: "PoolId"},
			_jsii_.MemberProperty{JsiiProperty: "poolIdInput", GoGetter: "PoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "popNames", GoGetter: "PopNames"},
			_jsii_.MemberProperty{JsiiProperty: "popNamesInput", GoGetter: "PopNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "product", GoGetter: "Product"},
			_jsii_.MemberProperty{JsiiProperty: "productInput", GoGetter: "ProductInput"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryTag", GoGetter: "QueryTag"},
			_jsii_.MemberProperty{JsiiProperty: "queryTagInput", GoGetter: "QueryTagInput"},
			_jsii_.MemberProperty{JsiiProperty: "requestsPerSecond", GoGetter: "RequestsPerSecond"},
			_jsii_.MemberProperty{JsiiProperty: "requestsPerSecondInput", GoGetter: "RequestsPerSecondInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetActions", GoMethod: "ResetActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAffectedAsns", GoMethod: "ResetAffectedAsns"},
			_jsii_.MemberMethod{JsiiMethod: "resetAffectedComponents", GoMethod: "ResetAffectedComponents"},
			_jsii_.MemberMethod{JsiiMethod: "resetAffectedLocations", GoMethod: "ResetAffectedLocations"},
			_jsii_.MemberMethod{JsiiMethod: "resetAirportCode", GoMethod: "ResetAirportCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlertTriggerPreferences", GoMethod: "ResetAlertTriggerPreferences"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlertTriggerPreferencesValue", GoMethod: "ResetAlertTriggerPreferencesValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnvironment", GoMethod: "ResetEnvironment"},
			_jsii_.MemberMethod{JsiiMethod: "resetEvent", GoMethod: "ResetEvent"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventSource", GoMethod: "ResetEventSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventType", GoMethod: "ResetEventType"},
			_jsii_.MemberMethod{JsiiMethod: "resetGroupBy", GoMethod: "ResetGroupBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetHealthCheckId", GoMethod: "ResetHealthCheckId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncidentImpact", GoMethod: "ResetIncidentImpact"},
			_jsii_.MemberMethod{JsiiMethod: "resetInputId", GoMethod: "ResetInputId"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsightClass", GoMethod: "ResetInsightClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetLimit", GoMethod: "ResetLimit"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogoTag", GoMethod: "ResetLogoTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetMegabitsPerSecond", GoMethod: "ResetMegabitsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewHealth", GoMethod: "ResetNewHealth"},
			_jsii_.MemberMethod{JsiiMethod: "resetNewStatus", GoMethod: "ResetNewStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetPacketsPerSecond", GoMethod: "ResetPacketsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "resetPoolId", GoMethod: "ResetPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPopNames", GoMethod: "ResetPopNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetProduct", GoMethod: "ResetProduct"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectId", GoMethod: "ResetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocol", GoMethod: "ResetProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryTag", GoMethod: "ResetQueryTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestsPerSecond", GoMethod: "ResetRequestsPerSecond"},
			_jsii_.MemberMethod{JsiiMethod: "resetSelectors", GoMethod: "ResetSelectors"},
			_jsii_.MemberMethod{JsiiMethod: "resetServices", GoMethod: "ResetServices"},
			_jsii_.MemberMethod{JsiiMethod: "resetSlo", GoMethod: "ResetSlo"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetHostname", GoMethod: "ResetTargetHostname"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetIp", GoMethod: "ResetTargetIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetZoneName", GoMethod: "ResetTargetZoneName"},
			_jsii_.MemberMethod{JsiiMethod: "resetTrafficExclusions", GoMethod: "ResetTrafficExclusions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunnelId", GoMethod: "ResetTunnelId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTunnelName", GoMethod: "ResetTunnelName"},
			_jsii_.MemberMethod{JsiiMethod: "resetWhere", GoMethod: "ResetWhere"},
			_jsii_.MemberMethod{JsiiMethod: "resetZones", GoMethod: "ResetZones"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "selectors", GoGetter: "Selectors"},
			_jsii_.MemberProperty{JsiiProperty: "selectorsInput", GoGetter: "SelectorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "services", GoGetter: "Services"},
			_jsii_.MemberProperty{JsiiProperty: "servicesInput", GoGetter: "ServicesInput"},
			_jsii_.MemberProperty{JsiiProperty: "slo", GoGetter: "Slo"},
			_jsii_.MemberProperty{JsiiProperty: "sloInput", GoGetter: "SloInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetHostname", GoGetter: "TargetHostname"},
			_jsii_.MemberProperty{JsiiProperty: "targetHostnameInput", GoGetter: "TargetHostnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetIp", GoGetter: "TargetIp"},
			_jsii_.MemberProperty{JsiiProperty: "targetIpInput", GoGetter: "TargetIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetZoneName", GoGetter: "TargetZoneName"},
			_jsii_.MemberProperty{JsiiProperty: "targetZoneNameInput", GoGetter: "TargetZoneNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "trafficExclusions", GoGetter: "TrafficExclusions"},
			_jsii_.MemberProperty{JsiiProperty: "trafficExclusionsInput", GoGetter: "TrafficExclusionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelId", GoGetter: "TunnelId"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelIdInput", GoGetter: "TunnelIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelName", GoGetter: "TunnelName"},
			_jsii_.MemberProperty{JsiiProperty: "tunnelNameInput", GoGetter: "TunnelNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "where", GoGetter: "Where"},
			_jsii_.MemberProperty{JsiiProperty: "whereInput", GoGetter: "WhereInput"},
			_jsii_.MemberProperty{JsiiProperty: "zones", GoGetter: "Zones"},
			_jsii_.MemberProperty{JsiiProperty: "zonesInput", GoGetter: "ZonesInput"},
		},
		func() interface{} {
			j := jsiiProxy_NotificationPolicyFiltersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanisms",
		reflect.TypeOf((*NotificationPolicyMechanisms)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsEmail",
		reflect.TypeOf((*NotificationPolicyMechanismsEmail)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsEmailList",
		reflect.TypeOf((*NotificationPolicyMechanismsEmailList)(nil)).Elem(),
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
			j := jsiiProxy_NotificationPolicyMechanismsEmailList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsEmailOutputReference",
		reflect.TypeOf((*NotificationPolicyMechanismsEmailOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NotificationPolicyMechanismsEmailOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsOutputReference",
		reflect.TypeOf((*NotificationPolicyMechanismsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "pagerduty", GoGetter: "Pagerduty"},
			_jsii_.MemberProperty{JsiiProperty: "pagerdutyInput", GoGetter: "PagerdutyInput"},
			_jsii_.MemberMethod{JsiiMethod: "putEmail", GoMethod: "PutEmail"},
			_jsii_.MemberMethod{JsiiMethod: "putPagerduty", GoMethod: "PutPagerduty"},
			_jsii_.MemberMethod{JsiiMethod: "putWebhooks", GoMethod: "PutWebhooks"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetPagerduty", GoMethod: "ResetPagerduty"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebhooks", GoMethod: "ResetWebhooks"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webhooks", GoGetter: "Webhooks"},
			_jsii_.MemberProperty{JsiiProperty: "webhooksInput", GoGetter: "WebhooksInput"},
		},
		func() interface{} {
			j := jsiiProxy_NotificationPolicyMechanismsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsPagerduty",
		reflect.TypeOf((*NotificationPolicyMechanismsPagerduty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsPagerdutyList",
		reflect.TypeOf((*NotificationPolicyMechanismsPagerdutyList)(nil)).Elem(),
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
			j := jsiiProxy_NotificationPolicyMechanismsPagerdutyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsPagerdutyOutputReference",
		reflect.TypeOf((*NotificationPolicyMechanismsPagerdutyOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NotificationPolicyMechanismsPagerdutyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsWebhooks",
		reflect.TypeOf((*NotificationPolicyMechanismsWebhooks)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsWebhooksList",
		reflect.TypeOf((*NotificationPolicyMechanismsWebhooksList)(nil)).Elem(),
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
			j := jsiiProxy_NotificationPolicyMechanismsWebhooksList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.notificationPolicy.NotificationPolicyMechanismsWebhooksOutputReference",
		reflect.TypeOf((*NotificationPolicyMechanismsWebhooksOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_NotificationPolicyMechanismsWebhooksOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
