package googlecomputesnapshot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeSnapshotConfig struct {
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
	// Name of the resource;
	//
	// provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#name GoogleComputeSnapshot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A reference to the disk used to create this snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#source_disk GoogleComputeSnapshot#source_disk}
	SourceDisk *string `field:"required" json:"sourceDisk" yaml:"sourceDisk"`
	// Creates the new snapshot in the snapshot chain labeled with the specified name.
	//
	// The chain name must be 1-63 characters long and
	// comply with RFC1035. This is an uncommon option only for advanced
	// service owners who needs to create separate snapshot chains, for
	// example, for chargeback tracking.  When you describe your snapshot
	// resource, this field is visible only if it has a non-empty value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#chain_name GoogleComputeSnapshot#chain_name}
	ChainName *string `field:"optional" json:"chainName" yaml:"chainName"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#description GoogleComputeSnapshot#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#id GoogleComputeSnapshot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels to apply to this Snapshot.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#labels GoogleComputeSnapshot#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#project GoogleComputeSnapshot#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// snapshot_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#snapshot_encryption_key GoogleComputeSnapshot#snapshot_encryption_key}
	SnapshotEncryptionKey *GoogleComputeSnapshotSnapshotEncryptionKey `field:"optional" json:"snapshotEncryptionKey" yaml:"snapshotEncryptionKey"`
	// source_disk_encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#source_disk_encryption_key GoogleComputeSnapshot#source_disk_encryption_key}
	SourceDiskEncryptionKey *GoogleComputeSnapshotSourceDiskEncryptionKey `field:"optional" json:"sourceDiskEncryptionKey" yaml:"sourceDiskEncryptionKey"`
	// Cloud Storage bucket storage location of the snapshot (regional or multi-regional).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#storage_locations GoogleComputeSnapshot#storage_locations}
	StorageLocations *[]*string `field:"optional" json:"storageLocations" yaml:"storageLocations"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#timeouts GoogleComputeSnapshot#timeouts}
	Timeouts *GoogleComputeSnapshotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A reference to the zone where the disk is hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_snapshot#zone GoogleComputeSnapshot#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

