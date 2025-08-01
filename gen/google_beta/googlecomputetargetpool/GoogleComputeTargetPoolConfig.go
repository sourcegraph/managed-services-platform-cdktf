package googlecomputetargetpool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeTargetPoolConfig struct {
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
	// A unique name for the resource, required by GCE. Changing this forces a new resource to be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#name GoogleComputeTargetPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// URL to the backup target pool. Must also set failover_ratio.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#backup_pool GoogleComputeTargetPool#backup_pool}
	BackupPool *string `field:"optional" json:"backupPool" yaml:"backupPool"`
	// Textual description field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#description GoogleComputeTargetPool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Ratio (0 to 1) of failed nodes before using the backup pool (which must also be set).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#failover_ratio GoogleComputeTargetPool#failover_ratio}
	FailoverRatio *float64 `field:"optional" json:"failoverRatio" yaml:"failoverRatio"`
	// List of zero or one health check name or self_link. Only legacy google_compute_http_health_check is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#health_checks GoogleComputeTargetPool#health_checks}
	HealthChecks *[]*string `field:"optional" json:"healthChecks" yaml:"healthChecks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#id GoogleComputeTargetPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of instances in the pool.
	//
	// They can be given as URLs, or in the form of "zone/name". Note that the instances need not exist at the time of target pool creation, so there is no need to use the Terraform interpolators to create a dependency on the instances from the target pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#instances GoogleComputeTargetPool#instances}
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// The ID of the project in which the resource belongs.
	//
	// If it is not provided, the provider project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#project GoogleComputeTargetPool#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Where the target pool resides. Defaults to project region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#region GoogleComputeTargetPool#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The resource URL for the security policy associated with this target pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#security_policy GoogleComputeTargetPool#security_policy}
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
	// How to distribute load.
	//
	// Options are "NONE" (no affinity). "CLIENT_IP" (hash of the source/dest addresses / ports), and "CLIENT_IP_PROTO" also includes the protocol (default "NONE").
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#session_affinity GoogleComputeTargetPool#session_affinity}
	SessionAffinity *string `field:"optional" json:"sessionAffinity" yaml:"sessionAffinity"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_target_pool#timeouts GoogleComputeTargetPool#timeouts}
	Timeouts *GoogleComputeTargetPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

