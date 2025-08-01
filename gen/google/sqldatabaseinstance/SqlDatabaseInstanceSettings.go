package sqldatabaseinstance


type SqlDatabaseInstanceSettings struct {
	// The machine type to use.
	//
	// See tiers for more details and supported versions. Postgres supports only shared-core machine types, and custom machine types such as db-custom-2-13312. See the Custom Machine Type Documentation to learn about specifying custom machine types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#tier SqlDatabaseInstance#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// This specifies when the instance should be active. Can be either ALWAYS, NEVER or ON_DEMAND.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#activation_policy SqlDatabaseInstance#activation_policy}
	ActivationPolicy *string `field:"optional" json:"activationPolicy" yaml:"activationPolicy"`
	// active_directory_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#active_directory_config SqlDatabaseInstance#active_directory_config}
	ActiveDirectoryConfig *SqlDatabaseInstanceSettingsActiveDirectoryConfig `field:"optional" json:"activeDirectoryConfig" yaml:"activeDirectoryConfig"`
	// advanced_machine_features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#advanced_machine_features SqlDatabaseInstance#advanced_machine_features}
	AdvancedMachineFeatures *SqlDatabaseInstanceSettingsAdvancedMachineFeatures `field:"optional" json:"advancedMachineFeatures" yaml:"advancedMachineFeatures"`
	// The availability type of the Cloud SQL instance, high availability (REGIONAL) or single zone (ZONAL).
	//
	// For all instances, ensure that
	// settings.backup_configuration.enabled is set to true.
	// For MySQL instances, ensure that settings.backup_configuration.binary_log_enabled is set to true.
	// For Postgres instances, ensure that settings.backup_configuration.point_in_time_recovery_enabled
	// is set to true. Defaults to ZONAL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#availability_type SqlDatabaseInstance#availability_type}
	AvailabilityType *string `field:"optional" json:"availabilityType" yaml:"availabilityType"`
	// backup_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#backup_configuration SqlDatabaseInstance#backup_configuration}
	BackupConfiguration *SqlDatabaseInstanceSettingsBackupConfiguration `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	// The name of server instance collation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#collation SqlDatabaseInstance#collation}
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// connection_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#connection_pool_config SqlDatabaseInstance#connection_pool_config}
	ConnectionPoolConfig interface{} `field:"optional" json:"connectionPoolConfig" yaml:"connectionPoolConfig"`
	// Enables the enforcement of Cloud SQL Auth Proxy or Cloud SQL connectors for all the connections.
	//
	// If enabled, all the direct connections are rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#connector_enforcement SqlDatabaseInstance#connector_enforcement}
	ConnectorEnforcement *string `field:"optional" json:"connectorEnforcement" yaml:"connectorEnforcement"`
	// database_flags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#database_flags SqlDatabaseInstance#database_flags}
	DatabaseFlags interface{} `field:"optional" json:"databaseFlags" yaml:"databaseFlags"`
	// data_cache_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#data_cache_config SqlDatabaseInstance#data_cache_config}
	DataCacheConfig *SqlDatabaseInstanceSettingsDataCacheConfig `field:"optional" json:"dataCacheConfig" yaml:"dataCacheConfig"`
	// Configuration to protect against accidental instance deletion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#deletion_protection_enabled SqlDatabaseInstance#deletion_protection_enabled}
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// deny_maintenance_period block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#deny_maintenance_period SqlDatabaseInstance#deny_maintenance_period}
	DenyMaintenancePeriod *SqlDatabaseInstanceSettingsDenyMaintenancePeriod `field:"optional" json:"denyMaintenancePeriod" yaml:"denyMaintenancePeriod"`
	// Enables auto-resizing of the storage size. Defaults to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#disk_autoresize SqlDatabaseInstance#disk_autoresize}
	DiskAutoresize interface{} `field:"optional" json:"diskAutoresize" yaml:"diskAutoresize"`
	// The maximum size, in GB, to which storage capacity can be automatically increased.
	//
	// The default value is 0, which specifies that there is no limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#disk_autoresize_limit SqlDatabaseInstance#disk_autoresize_limit}
	DiskAutoresizeLimit *float64 `field:"optional" json:"diskAutoresizeLimit" yaml:"diskAutoresizeLimit"`
	// The size of data disk, in GB.
	//
	// Size of a running instance cannot be reduced but can be increased. The minimum value is 10GB for PD_SSD, PD_HDD and 20GB for HYPERDISK_BALANCED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#disk_size SqlDatabaseInstance#disk_size}
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// The type of supported data disk is tier dependent and can be PD_SSD or PD_HDD or HYPERDISK_BALANCED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#disk_type SqlDatabaseInstance#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// The edition of the instance, can be ENTERPRISE or ENTERPRISE_PLUS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#edition SqlDatabaseInstance#edition}
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Enables Dataplex Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#enable_dataplex_integration SqlDatabaseInstance#enable_dataplex_integration}
	EnableDataplexIntegration interface{} `field:"optional" json:"enableDataplexIntegration" yaml:"enableDataplexIntegration"`
	// Enables Vertex AI Integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#enable_google_ml_integration SqlDatabaseInstance#enable_google_ml_integration}
	EnableGoogleMlIntegration interface{} `field:"optional" json:"enableGoogleMlIntegration" yaml:"enableGoogleMlIntegration"`
	// insights_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#insights_config SqlDatabaseInstance#insights_config}
	InsightsConfig *SqlDatabaseInstanceSettingsInsightsConfig `field:"optional" json:"insightsConfig" yaml:"insightsConfig"`
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#ip_configuration SqlDatabaseInstance#ip_configuration}
	IpConfiguration *SqlDatabaseInstanceSettingsIpConfiguration `field:"optional" json:"ipConfiguration" yaml:"ipConfiguration"`
	// location_preference block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#location_preference SqlDatabaseInstance#location_preference}
	LocationPreference *SqlDatabaseInstanceSettingsLocationPreference `field:"optional" json:"locationPreference" yaml:"locationPreference"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#maintenance_window SqlDatabaseInstance#maintenance_window}
	MaintenanceWindow *SqlDatabaseInstanceSettingsMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// password_validation_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#password_validation_policy SqlDatabaseInstance#password_validation_policy}
	PasswordValidationPolicy *SqlDatabaseInstanceSettingsPasswordValidationPolicy `field:"optional" json:"passwordValidationPolicy" yaml:"passwordValidationPolicy"`
	// Pricing plan for this instance, can only be PER_USE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#pricing_plan SqlDatabaseInstance#pricing_plan}
	PricingPlan *string `field:"optional" json:"pricingPlan" yaml:"pricingPlan"`
	// When this parameter is set to true, Cloud SQL retains backups of the instance even after the instance is deleted.
	//
	// The ON_DEMAND backup will be retained until customer deletes the backup or the project. The AUTOMATED backup will be retained based on the backups retention setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#retain_backups_on_delete SqlDatabaseInstance#retain_backups_on_delete}
	RetainBackupsOnDelete interface{} `field:"optional" json:"retainBackupsOnDelete" yaml:"retainBackupsOnDelete"`
	// sql_server_audit_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#sql_server_audit_config SqlDatabaseInstance#sql_server_audit_config}
	SqlServerAuditConfig *SqlDatabaseInstanceSettingsSqlServerAuditConfig `field:"optional" json:"sqlServerAuditConfig" yaml:"sqlServerAuditConfig"`
	// The time_zone to be used by the database engine (supported only for SQL Server), in SQL Server timezone format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#time_zone SqlDatabaseInstance#time_zone}
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// A set of key/value user label pairs to assign to the instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/sql_database_instance#user_labels SqlDatabaseInstance#user_labels}
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
}

