package dataproccluster


type DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy struct {
	// instance_selection_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_cluster#instance_selection_list DataprocCluster#instance_selection_list}
	InstanceSelectionList interface{} `field:"optional" json:"instanceSelectionList" yaml:"instanceSelectionList"`
	// provisioning_model_mix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_cluster#provisioning_model_mix DataprocCluster#provisioning_model_mix}
	ProvisioningModelMix *DataprocClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix `field:"optional" json:"provisioningModelMix" yaml:"provisioningModelMix"`
}

