package googlenotebooksruntime


type GoogleNotebooksRuntimeSoftwareConfig struct {
	// Specify a custom Cloud Storage path where the GPU driver is stored.
	//
	// If not specified, we'll automatically choose from official GPU drivers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#custom_gpu_driver_path GoogleNotebooksRuntime#custom_gpu_driver_path}
	CustomGpuDriverPath *string `field:"optional" json:"customGpuDriverPath" yaml:"customGpuDriverPath"`
	// Verifies core internal services are running. Default: True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#enable_health_monitoring GoogleNotebooksRuntime#enable_health_monitoring}
	EnableHealthMonitoring interface{} `field:"optional" json:"enableHealthMonitoring" yaml:"enableHealthMonitoring"`
	// Runtime will automatically shutdown after idle_shutdown_time. Default: True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#idle_shutdown GoogleNotebooksRuntime#idle_shutdown}
	IdleShutdown interface{} `field:"optional" json:"idleShutdown" yaml:"idleShutdown"`
	// Time in minutes to wait before shuting down runtime. Default: 180 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#idle_shutdown_timeout GoogleNotebooksRuntime#idle_shutdown_timeout}
	IdleShutdownTimeout *float64 `field:"optional" json:"idleShutdownTimeout" yaml:"idleShutdownTimeout"`
	// Install Nvidia Driver automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#install_gpu_driver GoogleNotebooksRuntime#install_gpu_driver}
	InstallGpuDriver interface{} `field:"optional" json:"installGpuDriver" yaml:"installGpuDriver"`
	// kernels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#kernels GoogleNotebooksRuntime#kernels}
	Kernels interface{} `field:"optional" json:"kernels" yaml:"kernels"`
	// Cron expression in UTC timezone for schedule instance auto upgrade. Please follow the [cron format](https://en.wikipedia.org/wiki/Cron).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#notebook_upgrade_schedule GoogleNotebooksRuntime#notebook_upgrade_schedule}
	NotebookUpgradeSchedule *string `field:"optional" json:"notebookUpgradeSchedule" yaml:"notebookUpgradeSchedule"`
	// Path to a Bash script that automatically runs after a notebook instance fully boots up.
	//
	// The path must be a URL or
	// Cloud Storage path (gs://path-to-file/file-name).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#post_startup_script GoogleNotebooksRuntime#post_startup_script}
	PostStartupScript *string `field:"optional" json:"postStartupScript" yaml:"postStartupScript"`
	// Behavior for the post startup script. Possible values: ["POST_STARTUP_SCRIPT_BEHAVIOR_UNSPECIFIED", "RUN_EVERY_START", "DOWNLOAD_AND_RUN_EVERY_START"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_notebooks_runtime#post_startup_script_behavior GoogleNotebooksRuntime#post_startup_script_behavior}
	PostStartupScriptBehavior *string `field:"optional" json:"postStartupScriptBehavior" yaml:"postStartupScriptBehavior"`
}

