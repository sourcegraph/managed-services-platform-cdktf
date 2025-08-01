package googlenetappbackupvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetappBackupVaultConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Location (region) of the backup vault.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#location GoogleNetappBackupVault#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the backup vault. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#name GoogleNetappBackupVault#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Region in which backup is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#backup_region GoogleNetappBackupVault#backup_region}
	BackupRegion *string `field:"optional" json:"backupRegion" yaml:"backupRegion"`
	// backup_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#backup_retention_policy GoogleNetappBackupVault#backup_retention_policy}
	BackupRetentionPolicy *GoogleNetappBackupVaultBackupRetentionPolicy `field:"optional" json:"backupRetentionPolicy" yaml:"backupRetentionPolicy"`
	// Type of the backup vault to be created. Default is IN_REGION. Possible values: ["BACKUP_VAULT_TYPE_UNSPECIFIED", "IN_REGION", "CROSS_REGION"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#backup_vault_type GoogleNetappBackupVault#backup_vault_type}
	BackupVaultType *string `field:"optional" json:"backupVaultType" yaml:"backupVaultType"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#description GoogleNetappBackupVault#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#id GoogleNetappBackupVault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#labels GoogleNetappBackupVault#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#project GoogleNetappBackupVault#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_netapp_backup_vault#timeouts GoogleNetappBackupVault#timeouts}
	Timeouts *GoogleNetappBackupVaultTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

