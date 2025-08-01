package googleoracledatabasecloudexadatainfrastructure


type GoogleOracleDatabaseCloudExadataInfrastructureProperties struct {
	// The shape of the Exadata Infrastructure.
	//
	// The shape determines the
	// amount of CPU, storage, and memory resources allocated to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure#shape GoogleOracleDatabaseCloudExadataInfrastructure#shape}
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// The number of compute servers for the Exadata Infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure#compute_count GoogleOracleDatabaseCloudExadataInfrastructure#compute_count}
	ComputeCount *float64 `field:"optional" json:"computeCount" yaml:"computeCount"`
	// customer_contacts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure#customer_contacts GoogleOracleDatabaseCloudExadataInfrastructure#customer_contacts}
	CustomerContacts interface{} `field:"optional" json:"customerContacts" yaml:"customerContacts"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure#maintenance_window GoogleOracleDatabaseCloudExadataInfrastructure#maintenance_window}
	MaintenanceWindow *GoogleOracleDatabaseCloudExadataInfrastructurePropertiesMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The number of Cloud Exadata storage servers for the Exadata Infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure#storage_count GoogleOracleDatabaseCloudExadataInfrastructure#storage_count}
	StorageCount *float64 `field:"optional" json:"storageCount" yaml:"storageCount"`
	// The total storage allocated to the Exadata Infrastructure resource, in gigabytes (GB).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_oracle_database_cloud_exadata_infrastructure#total_storage_size_gb GoogleOracleDatabaseCloudExadataInfrastructure#total_storage_size_gb}
	TotalStorageSizeGb *float64 `field:"optional" json:"totalStorageSizeGb" yaml:"totalStorageSizeGb"`
}

