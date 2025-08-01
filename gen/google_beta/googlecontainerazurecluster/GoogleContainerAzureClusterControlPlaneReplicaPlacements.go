package googlecontainerazurecluster


type GoogleContainerAzureClusterControlPlaneReplicaPlacements struct {
	// For a given replica, the Azure availability zone where to provision the control plane VM and the ETCD disk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_azure_cluster#azure_availability_zone GoogleContainerAzureCluster#azure_availability_zone}
	AzureAvailabilityZone *string `field:"required" json:"azureAvailabilityZone" yaml:"azureAvailabilityZone"`
	// For a given replica, the ARM ID of the subnet where the control plane VM is deployed.
	//
	// Make sure it's a subnet under the virtual network in the cluster configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_azure_cluster#subnet_id GoogleContainerAzureCluster#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

