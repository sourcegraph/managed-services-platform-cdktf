package googledataflowflextemplatejob

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google_beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJob",
		reflect.TypeOf((*GoogleDataflowFlexTemplateJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalExperiments", GoGetter: "AdditionalExperiments"},
			_jsii_.MemberProperty{JsiiProperty: "additionalExperimentsInput", GoGetter: "AdditionalExperimentsInput"},
			_jsii_.MemberProperty{JsiiProperty: "additionalPipelineOptions", GoGetter: "AdditionalPipelineOptions"},
			_jsii_.MemberProperty{JsiiProperty: "additionalPipelineOptionsInput", GoGetter: "AdditionalPipelineOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingAlgorithm", GoGetter: "AutoscalingAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingAlgorithmInput", GoGetter: "AutoscalingAlgorithmInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerSpecGcsPath", GoGetter: "ContainerSpecGcsPath"},
			_jsii_.MemberProperty{JsiiProperty: "containerSpecGcsPathInput", GoGetter: "ContainerSpecGcsPathInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveLabels", GoGetter: "EffectiveLabels"},
			_jsii_.MemberProperty{JsiiProperty: "enableStreamingEngine", GoGetter: "EnableStreamingEngine"},
			_jsii_.MemberProperty{JsiiProperty: "enableStreamingEngineInput", GoGetter: "EnableStreamingEngineInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipConfiguration", GoGetter: "IpConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "ipConfigurationInput", GoGetter: "IpConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "jobId", GoGetter: "JobId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyName", GoGetter: "KmsKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyNameInput", GoGetter: "KmsKeyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "launcherMachineType", GoGetter: "LauncherMachineType"},
			_jsii_.MemberProperty{JsiiProperty: "launcherMachineTypeInput", GoGetter: "LauncherMachineTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "machineType", GoGetter: "MachineType"},
			_jsii_.MemberProperty{JsiiProperty: "machineTypeInput", GoGetter: "MachineTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkers", GoGetter: "MaxWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "maxWorkersInput", GoGetter: "MaxWorkersInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "network", GoGetter: "Network"},
			_jsii_.MemberProperty{JsiiProperty: "networkInput", GoGetter: "NetworkInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "numWorkers", GoGetter: "NumWorkers"},
			_jsii_.MemberProperty{JsiiProperty: "numWorkersInput", GoGetter: "NumWorkersInput"},
			_jsii_.MemberProperty{JsiiProperty: "onDelete", GoGetter: "OnDelete"},
			_jsii_.MemberProperty{JsiiProperty: "onDeleteInput", GoGetter: "OnDeleteInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalExperiments", GoMethod: "ResetAdditionalExperiments"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalPipelineOptions", GoMethod: "ResetAdditionalPipelineOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoscalingAlgorithm", GoMethod: "ResetAutoscalingAlgorithm"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableStreamingEngine", GoMethod: "ResetEnableStreamingEngine"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpConfiguration", GoMethod: "ResetIpConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyName", GoMethod: "ResetKmsKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetLauncherMachineType", GoMethod: "ResetLauncherMachineType"},
			_jsii_.MemberMethod{JsiiMethod: "resetMachineType", GoMethod: "ResetMachineType"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxWorkers", GoMethod: "ResetMaxWorkers"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetwork", GoMethod: "ResetNetwork"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumWorkers", GoMethod: "ResetNumWorkers"},
			_jsii_.MemberMethod{JsiiMethod: "resetOnDelete", GoMethod: "ResetOnDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSdkContainerImage", GoMethod: "ResetSdkContainerImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountEmail", GoMethod: "ResetServiceAccountEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipWaitOnJobTermination", GoMethod: "ResetSkipWaitOnJobTermination"},
			_jsii_.MemberMethod{JsiiMethod: "resetStagingLocation", GoMethod: "ResetStagingLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetwork", GoMethod: "ResetSubnetwork"},
			_jsii_.MemberMethod{JsiiMethod: "resetTempLocation", GoMethod: "ResetTempLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransformNameMapping", GoMethod: "ResetTransformNameMapping"},
			_jsii_.MemberProperty{JsiiProperty: "sdkContainerImage", GoGetter: "SdkContainerImage"},
			_jsii_.MemberProperty{JsiiProperty: "sdkContainerImageInput", GoGetter: "SdkContainerImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountEmail", GoGetter: "ServiceAccountEmail"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountEmailInput", GoGetter: "ServiceAccountEmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipWaitOnJobTermination", GoGetter: "SkipWaitOnJobTermination"},
			_jsii_.MemberProperty{JsiiProperty: "skipWaitOnJobTerminationInput", GoGetter: "SkipWaitOnJobTerminationInput"},
			_jsii_.MemberProperty{JsiiProperty: "stagingLocation", GoGetter: "StagingLocation"},
			_jsii_.MemberProperty{JsiiProperty: "stagingLocationInput", GoGetter: "StagingLocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "subnetwork", GoGetter: "Subnetwork"},
			_jsii_.MemberProperty{JsiiProperty: "subnetworkInput", GoGetter: "SubnetworkInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tempLocation", GoGetter: "TempLocation"},
			_jsii_.MemberProperty{JsiiProperty: "tempLocationInput", GoGetter: "TempLocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformLabels", GoGetter: "TerraformLabels"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "transformNameMapping", GoGetter: "TransformNameMapping"},
			_jsii_.MemberProperty{JsiiProperty: "transformNameMappingInput", GoGetter: "TransformNameMappingInput"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataflowFlexTemplateJob{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google_beta.googleDataflowFlexTemplateJob.GoogleDataflowFlexTemplateJobConfig",
		reflect.TypeOf((*GoogleDataflowFlexTemplateJobConfig)(nil)).Elem(),
	)
}
