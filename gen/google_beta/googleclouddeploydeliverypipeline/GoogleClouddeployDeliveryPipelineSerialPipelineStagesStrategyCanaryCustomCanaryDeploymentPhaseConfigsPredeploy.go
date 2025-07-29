package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy struct {
	// Optional. A sequence of skaffold custom actions to invoke during execution of the predeploy job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_delivery_pipeline#actions GoogleClouddeployDeliveryPipeline#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
}

