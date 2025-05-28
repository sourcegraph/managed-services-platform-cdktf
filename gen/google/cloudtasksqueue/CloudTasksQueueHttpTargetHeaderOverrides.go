package cloudtasksqueue


type CloudTasksQueueHttpTargetHeaderOverrides struct {
	// header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.36.1/docs/resources/cloud_tasks_queue#header CloudTasksQueue#header}
	Header *CloudTasksQueueHttpTargetHeaderOverridesHeader `field:"required" json:"header" yaml:"header"`
}

