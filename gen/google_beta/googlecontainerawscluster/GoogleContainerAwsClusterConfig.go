package googlecontainerawscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleContainerAwsClusterConfig struct {
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
	// authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#authorization GoogleContainerAwsCluster#authorization}
	Authorization *GoogleContainerAwsClusterAuthorization `field:"required" json:"authorization" yaml:"authorization"`
	// The AWS region where the cluster runs.
	//
	// Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#aws_region GoogleContainerAwsCluster#aws_region}
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// control_plane block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#control_plane GoogleContainerAwsCluster#control_plane}
	ControlPlane *GoogleContainerAwsClusterControlPlane `field:"required" json:"controlPlane" yaml:"controlPlane"`
	// fleet block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#fleet GoogleContainerAwsCluster#fleet}
	Fleet *GoogleContainerAwsClusterFleet `field:"required" json:"fleet" yaml:"fleet"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#location GoogleContainerAwsCluster#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#name GoogleContainerAwsCluster#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// networking block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#networking GoogleContainerAwsCluster#networking}
	Networking *GoogleContainerAwsClusterNetworking `field:"required" json:"networking" yaml:"networking"`
	// Optional.
	//
	// Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field `effective_annotations` for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#annotations GoogleContainerAwsCluster#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// binary_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#binary_authorization GoogleContainerAwsCluster#binary_authorization}
	BinaryAuthorization *GoogleContainerAwsClusterBinaryAuthorization `field:"optional" json:"binaryAuthorization" yaml:"binaryAuthorization"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#description GoogleContainerAwsCluster#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#id GoogleContainerAwsCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#logging_config GoogleContainerAwsCluster#logging_config}
	LoggingConfig *GoogleContainerAwsClusterLoggingConfig `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The project for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#project GoogleContainerAwsCluster#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_container_aws_cluster#timeouts GoogleContainerAwsCluster#timeouts}
	Timeouts *GoogleContainerAwsClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

