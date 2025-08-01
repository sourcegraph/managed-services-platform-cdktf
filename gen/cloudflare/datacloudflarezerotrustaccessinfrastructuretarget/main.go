package datacloudflarezerotrustaccessinfrastructuretarget

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTarget",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accountId", GoGetter: "AccountId"},
			_jsii_.MemberProperty{JsiiProperty: "accountIdInput", GoGetter: "AccountIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "filter", GoGetter: "Filter"},
			_jsii_.MemberProperty{JsiiProperty: "filterInput", GoGetter: "FilterInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ip", GoGetter: "Ip"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAt", GoGetter: "ModifiedAt"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberMethod{JsiiMethod: "putFilter", GoMethod: "PutFilter"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetFilter", GoMethod: "ResetFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetId", GoMethod: "ResetTargetId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetId", GoGetter: "TargetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetIdInput", GoGetter: "TargetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTarget{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetConfig",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetFilter",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetFilter)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "createdAfter", GoGetter: "CreatedAfter"},
			_jsii_.MemberProperty{JsiiProperty: "createdAfterInput", GoGetter: "CreatedAfterInput"},
			_jsii_.MemberProperty{JsiiProperty: "createdBefore", GoGetter: "CreatedBefore"},
			_jsii_.MemberProperty{JsiiProperty: "createdBeforeInput", GoGetter: "CreatedBeforeInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "direction", GoGetter: "Direction"},
			_jsii_.MemberProperty{JsiiProperty: "directionInput", GoGetter: "DirectionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostname", GoGetter: "Hostname"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameContains", GoGetter: "HostnameContains"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameContainsInput", GoGetter: "HostnameContainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hostnameInput", GoGetter: "HostnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipLike", GoGetter: "IpLike"},
			_jsii_.MemberProperty{JsiiProperty: "ipLikeInput", GoGetter: "IpLikeInput"},
			_jsii_.MemberProperty{JsiiProperty: "ips", GoGetter: "Ips"},
			_jsii_.MemberProperty{JsiiProperty: "ipsInput", GoGetter: "IpsInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipV4", GoGetter: "IpV4"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4End", GoGetter: "Ipv4End"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4EndInput", GoGetter: "Ipv4EndInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipV4Input", GoGetter: "IpV4Input"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4Start", GoGetter: "Ipv4Start"},
			_jsii_.MemberProperty{JsiiProperty: "ipv4StartInput", GoGetter: "Ipv4StartInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipV6", GoGetter: "IpV6"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6End", GoGetter: "Ipv6End"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6EndInput", GoGetter: "Ipv6EndInput"},
			_jsii_.MemberProperty{JsiiProperty: "ipV6Input", GoGetter: "IpV6Input"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6Start", GoGetter: "Ipv6Start"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6StartInput", GoGetter: "Ipv6StartInput"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAfter", GoGetter: "ModifiedAfter"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedAfterInput", GoGetter: "ModifiedAfterInput"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedBefore", GoGetter: "ModifiedBefore"},
			_jsii_.MemberProperty{JsiiProperty: "modifiedBeforeInput", GoGetter: "ModifiedBeforeInput"},
			_jsii_.MemberProperty{JsiiProperty: "order", GoGetter: "Order"},
			_jsii_.MemberProperty{JsiiProperty: "orderInput", GoGetter: "OrderInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreatedAfter", GoMethod: "ResetCreatedAfter"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreatedBefore", GoMethod: "ResetCreatedBefore"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirection", GoMethod: "ResetDirection"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostname", GoMethod: "ResetHostname"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostnameContains", GoMethod: "ResetHostnameContains"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpLike", GoMethod: "ResetIpLike"},
			_jsii_.MemberMethod{JsiiMethod: "resetIps", GoMethod: "ResetIps"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpV4", GoMethod: "ResetIpV4"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4End", GoMethod: "ResetIpv4End"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv4Start", GoMethod: "ResetIpv4Start"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpV6", GoMethod: "ResetIpV6"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6End", GoMethod: "ResetIpv6End"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpv6Start", GoMethod: "ResetIpv6Start"},
			_jsii_.MemberMethod{JsiiMethod: "resetModifiedAfter", GoMethod: "ResetModifiedAfter"},
			_jsii_.MemberMethod{JsiiMethod: "resetModifiedBefore", GoMethod: "ResetModifiedBefore"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrder", GoMethod: "ResetOrder"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetIds", GoMethod: "ResetTargetIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetVirtualNetworkId", GoMethod: "ResetVirtualNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetIds", GoGetter: "TargetIds"},
			_jsii_.MemberProperty{JsiiProperty: "targetIdsInput", GoGetter: "TargetIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkIdInput", GoGetter: "VirtualNetworkIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetFilterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetIp",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetIp)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv4",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv4)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv4OutputReference",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv4OutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipAddr", GoGetter: "IpAddr"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv4OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv6",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv6)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv6OutputReference",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv6OutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipAddr", GoGetter: "IpAddr"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNetworkId", GoGetter: "VirtualNetworkId"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetIpIpv6OutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.dataCloudflareZeroTrustAccessInfrastructureTarget.DataCloudflareZeroTrustAccessInfrastructureTargetIpOutputReference",
		reflect.TypeOf((*DataCloudflareZeroTrustAccessInfrastructureTargetIpOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "ipv4", GoGetter: "Ipv4"},
			_jsii_.MemberProperty{JsiiProperty: "ipv6", GoGetter: "Ipv6"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataCloudflareZeroTrustAccessInfrastructureTargetIpOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
