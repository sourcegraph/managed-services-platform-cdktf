package googlecloudbuildv2repository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCloudbuildv2RepositoryConfig struct {
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
	// Name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#name GoogleCloudbuildv2Repository#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The connection for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#parent_connection GoogleCloudbuildv2Repository#parent_connection}
	ParentConnection *string `field:"required" json:"parentConnection" yaml:"parentConnection"`
	// Required. Git Clone HTTPS URI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#remote_uri GoogleCloudbuildv2Repository#remote_uri}
	RemoteUri *string `field:"required" json:"remoteUri" yaml:"remoteUri"`
	// Allows clients to store small amounts of arbitrary data.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#annotations GoogleCloudbuildv2Repository#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#id GoogleCloudbuildv2Repository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location for the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#location GoogleCloudbuildv2Repository#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#project GoogleCloudbuildv2Repository#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloudbuildv2_repository#timeouts GoogleCloudbuildv2Repository#timeouts}
	Timeouts *GoogleCloudbuildv2RepositoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

