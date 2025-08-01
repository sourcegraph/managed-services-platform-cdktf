package workspace

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-tfe.workspace.Workspace",
		reflect.TypeOf((*Workspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agentPoolId", GoGetter: "AgentPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "agentPoolIdInput", GoGetter: "AgentPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowDestroyPlan", GoGetter: "AllowDestroyPlan"},
			_jsii_.MemberProperty{JsiiProperty: "allowDestroyPlanInput", GoGetter: "AllowDestroyPlanInput"},
			_jsii_.MemberProperty{JsiiProperty: "assessmentsEnabled", GoGetter: "AssessmentsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "assessmentsEnabledInput", GoGetter: "AssessmentsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoApply", GoGetter: "AutoApply"},
			_jsii_.MemberProperty{JsiiProperty: "autoApplyInput", GoGetter: "AutoApplyInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoApplyRunTrigger", GoGetter: "AutoApplyRunTrigger"},
			_jsii_.MemberProperty{JsiiProperty: "autoApplyRunTriggerInput", GoGetter: "AutoApplyRunTriggerInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoDestroyActivityDuration", GoGetter: "AutoDestroyActivityDuration"},
			_jsii_.MemberProperty{JsiiProperty: "autoDestroyActivityDurationInput", GoGetter: "AutoDestroyActivityDurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoDestroyAt", GoGetter: "AutoDestroyAt"},
			_jsii_.MemberProperty{JsiiProperty: "autoDestroyAtInput", GoGetter: "AutoDestroyAtInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveTags", GoGetter: "EffectiveTags"},
			_jsii_.MemberProperty{JsiiProperty: "executionMode", GoGetter: "ExecutionMode"},
			_jsii_.MemberProperty{JsiiProperty: "executionModeInput", GoGetter: "ExecutionModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileTriggersEnabled", GoGetter: "FileTriggersEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "fileTriggersEnabledInput", GoGetter: "FileTriggersEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "forceDelete", GoGetter: "ForceDelete"},
			_jsii_.MemberProperty{JsiiProperty: "forceDeleteInput", GoGetter: "ForceDeleteInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "globalRemoteState", GoGetter: "GlobalRemoteState"},
			_jsii_.MemberProperty{JsiiProperty: "globalRemoteStateInput", GoGetter: "GlobalRemoteStateInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "htmlUrl", GoGetter: "HtmlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreAdditionalTagNames", GoGetter: "IgnoreAdditionalTagNames"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreAdditionalTagNamesInput", GoGetter: "IgnoreAdditionalTagNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreAdditionalTags", GoGetter: "IgnoreAdditionalTags"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreAdditionalTagsInput", GoGetter: "IgnoreAdditionalTagsInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "inheritsProjectAutoDestroy", GoGetter: "InheritsProjectAutoDestroy"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "operations", GoGetter: "Operations"},
			_jsii_.MemberProperty{JsiiProperty: "operationsInput", GoGetter: "OperationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberProperty{JsiiProperty: "organizationInput", GoGetter: "OrganizationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "projectIdInput", GoGetter: "ProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putVcsRepo", GoMethod: "PutVcsRepo"},
			_jsii_.MemberProperty{JsiiProperty: "queueAllRuns", GoGetter: "QueueAllRuns"},
			_jsii_.MemberProperty{JsiiProperty: "queueAllRunsInput", GoGetter: "QueueAllRunsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "remoteStateConsumerIds", GoGetter: "RemoteStateConsumerIds"},
			_jsii_.MemberProperty{JsiiProperty: "remoteStateConsumerIdsInput", GoGetter: "RemoteStateConsumerIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAgentPoolId", GoMethod: "ResetAgentPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowDestroyPlan", GoMethod: "ResetAllowDestroyPlan"},
			_jsii_.MemberMethod{JsiiMethod: "resetAssessmentsEnabled", GoMethod: "ResetAssessmentsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoApply", GoMethod: "ResetAutoApply"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoApplyRunTrigger", GoMethod: "ResetAutoApplyRunTrigger"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDestroyActivityDuration", GoMethod: "ResetAutoDestroyActivityDuration"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDestroyAt", GoMethod: "ResetAutoDestroyAt"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetExecutionMode", GoMethod: "ResetExecutionMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileTriggersEnabled", GoMethod: "ResetFileTriggersEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetForceDelete", GoMethod: "ResetForceDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetGlobalRemoteState", GoMethod: "ResetGlobalRemoteState"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreAdditionalTagNames", GoMethod: "ResetIgnoreAdditionalTagNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreAdditionalTags", GoMethod: "ResetIgnoreAdditionalTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetOperations", GoMethod: "ResetOperations"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganization", GoMethod: "ResetOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProjectId", GoMethod: "ResetProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueueAllRuns", GoMethod: "ResetQueueAllRuns"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemoteStateConsumerIds", GoMethod: "ResetRemoteStateConsumerIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceName", GoMethod: "ResetSourceName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceUrl", GoMethod: "ResetSourceUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetSpeculativeEnabled", GoMethod: "ResetSpeculativeEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSshKeyId", GoMethod: "ResetSshKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStructuredRunOutputEnabled", GoMethod: "ResetStructuredRunOutputEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagNames", GoMethod: "ResetTagNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetTerraformVersion", GoMethod: "ResetTerraformVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggerPatterns", GoMethod: "ResetTriggerPatterns"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggerPrefixes", GoMethod: "ResetTriggerPrefixes"},
			_jsii_.MemberMethod{JsiiMethod: "resetVcsRepo", GoMethod: "ResetVcsRepo"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkingDirectory", GoMethod: "ResetWorkingDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "resourceCount", GoGetter: "ResourceCount"},
			_jsii_.MemberProperty{JsiiProperty: "sourceName", GoGetter: "SourceName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceNameInput", GoGetter: "SourceNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceUrl", GoGetter: "SourceUrl"},
			_jsii_.MemberProperty{JsiiProperty: "sourceUrlInput", GoGetter: "SourceUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "speculativeEnabled", GoGetter: "SpeculativeEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "speculativeEnabledInput", GoGetter: "SpeculativeEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "sshKeyId", GoGetter: "SshKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "sshKeyIdInput", GoGetter: "SshKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "structuredRunOutputEnabled", GoGetter: "StructuredRunOutputEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "structuredRunOutputEnabledInput", GoGetter: "StructuredRunOutputEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tagNames", GoGetter: "TagNames"},
			_jsii_.MemberProperty{JsiiProperty: "tagNamesInput", GoGetter: "TagNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "terraformVersion", GoGetter: "TerraformVersion"},
			_jsii_.MemberProperty{JsiiProperty: "terraformVersionInput", GoGetter: "TerraformVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "triggerPatterns", GoGetter: "TriggerPatterns"},
			_jsii_.MemberProperty{JsiiProperty: "triggerPatternsInput", GoGetter: "TriggerPatternsInput"},
			_jsii_.MemberProperty{JsiiProperty: "triggerPrefixes", GoGetter: "TriggerPrefixes"},
			_jsii_.MemberProperty{JsiiProperty: "triggerPrefixesInput", GoGetter: "TriggerPrefixesInput"},
			_jsii_.MemberProperty{JsiiProperty: "vcsRepo", GoGetter: "VcsRepo"},
			_jsii_.MemberProperty{JsiiProperty: "vcsRepoInput", GoGetter: "VcsRepoInput"},
			_jsii_.MemberProperty{JsiiProperty: "workingDirectory", GoGetter: "WorkingDirectory"},
			_jsii_.MemberProperty{JsiiProperty: "workingDirectoryInput", GoGetter: "WorkingDirectoryInput"},
		},
		func() interface{} {
			j := jsiiProxy_Workspace{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tfe.workspace.WorkspaceConfig",
		reflect.TypeOf((*WorkspaceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tfe.workspace.WorkspaceVcsRepo",
		reflect.TypeOf((*WorkspaceVcsRepo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-tfe.workspace.WorkspaceVcsRepoOutputReference",
		reflect.TypeOf((*WorkspaceVcsRepoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "branch", GoGetter: "Branch"},
			_jsii_.MemberProperty{JsiiProperty: "branchInput", GoGetter: "BranchInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "githubAppInstallationId", GoGetter: "GithubAppInstallationId"},
			_jsii_.MemberProperty{JsiiProperty: "githubAppInstallationIdInput", GoGetter: "GithubAppInstallationIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "identifier", GoGetter: "Identifier"},
			_jsii_.MemberProperty{JsiiProperty: "identifierInput", GoGetter: "IdentifierInput"},
			_jsii_.MemberProperty{JsiiProperty: "ingressSubmodules", GoGetter: "IngressSubmodules"},
			_jsii_.MemberProperty{JsiiProperty: "ingressSubmodulesInput", GoGetter: "IngressSubmodulesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "oauthTokenId", GoGetter: "OauthTokenId"},
			_jsii_.MemberProperty{JsiiProperty: "oauthTokenIdInput", GoGetter: "OauthTokenIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBranch", GoMethod: "ResetBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetGithubAppInstallationId", GoMethod: "ResetGithubAppInstallationId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIngressSubmodules", GoMethod: "ResetIngressSubmodules"},
			_jsii_.MemberMethod{JsiiMethod: "resetOauthTokenId", GoMethod: "ResetOauthTokenId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagsRegex", GoMethod: "ResetTagsRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRegex", GoGetter: "TagsRegex"},
			_jsii_.MemberProperty{JsiiProperty: "tagsRegexInput", GoGetter: "TagsRegexInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspaceVcsRepoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
