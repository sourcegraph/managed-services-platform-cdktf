package cloudtasksqueue


type CloudTasksQueueHttpTargetUriOverrideQueryOverride struct {
	// The query parameters (e.g., qparam1=123&qparam2=456). Default is an empty string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_tasks_queue#query_params CloudTasksQueue#query_params}
	QueryParams *string `field:"optional" json:"queryParams" yaml:"queryParams"`
}

