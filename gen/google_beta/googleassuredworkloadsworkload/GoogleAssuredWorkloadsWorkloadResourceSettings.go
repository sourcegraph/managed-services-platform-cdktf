package googleassuredworkloadsworkload


type GoogleAssuredWorkloadsWorkloadResourceSettings struct {
	// User-assigned resource display name. If not empty it will be used to create a resource with the specified name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_assured_workloads_workload#display_name GoogleAssuredWorkloadsWorkload#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Resource identifier.
	//
	// For a project this represents projectId. If the project is already taken, the workload creation will fail. For KeyRing, this represents the keyring_id. For a folder, don't set this value as folder_id is assigned by Google.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_assured_workloads_workload#resource_id GoogleAssuredWorkloadsWorkload#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Indicates the type of resource.
	//
	// This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_assured_workloads_workload#resource_type GoogleAssuredWorkloadsWorkload#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

