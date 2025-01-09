package googlenetappbackup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleNetappBackupConfig struct {
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
	// Location of the backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#location GoogleNetappBackup#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The resource name of the backup. Needs to be unique per location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#name GoogleNetappBackup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Name of the backup vault to store the backup in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#vault_name GoogleNetappBackup#vault_name}
	VaultName *string `field:"required" json:"vaultName" yaml:"vaultName"`
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#description GoogleNetappBackup#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#id GoogleNetappBackup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#labels GoogleNetappBackup#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#project GoogleNetappBackup#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// If specified, backup will be created from the given snapshot.
	//
	// If not specified,
	// there will be a new snapshot taken to initiate the backup creation.
	// Format: 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}''
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#source_snapshot GoogleNetappBackup#source_snapshot}
	SourceSnapshot *string `field:"optional" json:"sourceSnapshot" yaml:"sourceSnapshot"`
	// ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}''.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#source_volume GoogleNetappBackup#source_volume}
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_netapp_backup#timeouts GoogleNetappBackup#timeouts}
	Timeouts *GoogleNetappBackupTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

