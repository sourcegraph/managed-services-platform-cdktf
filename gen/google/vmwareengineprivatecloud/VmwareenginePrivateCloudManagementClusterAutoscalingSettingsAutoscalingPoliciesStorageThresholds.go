package vmwareengineprivatecloud


type VmwareenginePrivateCloudManagementClusterAutoscalingSettingsAutoscalingPoliciesStorageThresholds struct {
	// The utilization triggering the scale-in operation in percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#scale_in VmwareenginePrivateCloud#scale_in}
	ScaleIn *float64 `field:"required" json:"scaleIn" yaml:"scaleIn"`
	// The utilization triggering the scale-out operation in percent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/vmwareengine_private_cloud#scale_out VmwareenginePrivateCloud#scale_out}
	ScaleOut *float64 `field:"required" json:"scaleOut" yaml:"scaleOut"`
}

