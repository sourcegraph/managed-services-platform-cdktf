package colabschedule


type ColabScheduleCreateNotebookExecutionJobRequest struct {
	// notebook_execution_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/colab_schedule#notebook_execution_job ColabSchedule#notebook_execution_job}
	NotebookExecutionJob *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob `field:"required" json:"notebookExecutionJob" yaml:"notebookExecutionJob"`
}

