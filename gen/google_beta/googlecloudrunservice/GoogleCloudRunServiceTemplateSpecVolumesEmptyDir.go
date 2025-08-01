package googlecloudrunservice


type GoogleCloudRunServiceTemplateSpecVolumesEmptyDir struct {
	// The medium on which the data is stored.
	//
	// The default is "" which means to use the node's default medium. Must be an empty string (default) or Memory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_service#medium GoogleCloudRunService#medium}
	Medium *string `field:"optional" json:"medium" yaml:"medium"`
	// Limit on the storage usable by this EmptyDir volume.
	//
	// The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. This field's values are of the 'Quantity' k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_service#size_limit GoogleCloudRunService#size_limit}
	SizeLimit *string `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

