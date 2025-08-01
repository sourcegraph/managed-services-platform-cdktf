package cloudrunv2job


type CloudRunV2JobTemplateTemplateVolumesEmptyDir struct {
	// The different types of medium supported for EmptyDir. Default value: "MEMORY" Possible values: ["MEMORY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_run_v2_job#medium CloudRunV2Job#medium}
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// Limit on the storage usable by this EmptyDir volume.
	//
	// The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field's values are of the 'Quantity' k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/cloud_run_v2_job#size_limit CloudRunV2Job#size_limit}
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

