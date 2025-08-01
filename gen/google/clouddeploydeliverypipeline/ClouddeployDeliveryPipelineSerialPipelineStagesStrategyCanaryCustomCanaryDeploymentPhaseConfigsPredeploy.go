package clouddeploydeliverypipeline


type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCustomCanaryDeploymentPhaseConfigsPredeploy struct {
	// Optional. A sequence of skaffold custom actions to invoke during execution of the predeploy job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/clouddeploy_delivery_pipeline#actions ClouddeployDeliveryPipeline#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
}

