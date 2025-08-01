package googleclouddeploydeliverypipeline


type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryRuntimeConfigKubernetesServiceNetworking struct {
	// Required. Name of the Kubernetes Deployment whose traffic is managed by the specified Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_delivery_pipeline#deployment GoogleClouddeployDeliveryPipeline#deployment}
	Deployment *string `field:"required" json:"deployment" yaml:"deployment"`
	// Required. Name of the Kubernetes Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_delivery_pipeline#service GoogleClouddeployDeliveryPipeline#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// Optional.
	//
	// Whether to disable Pod overprovisioning. If Pod overprovisioning is disabled then Cloud Deploy will limit the number of total Pods used for the deployment strategy to the number of Pods the Deployment has on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_delivery_pipeline#disable_pod_overprovisioning GoogleClouddeployDeliveryPipeline#disable_pod_overprovisioning}
	DisablePodOverprovisioning interface{} `field:"optional" json:"disablePodOverprovisioning" yaml:"disablePodOverprovisioning"`
	// Optional.
	//
	// The label to use when selecting Pods for the Deployment resource. This label must already be present in the Deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_clouddeploy_delivery_pipeline#pod_selector_label GoogleClouddeployDeliveryPipeline#pod_selector_label}
	PodSelectorLabel *string `field:"optional" json:"podSelectorLabel" yaml:"podSelectorLabel"`
}

