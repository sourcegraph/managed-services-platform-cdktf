package googlesecretmanagerregionalsecretversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSecretManagerRegionalSecretVersionConfig struct {
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
	// Secret Manager regional secret resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret_version#secret GoogleSecretManagerRegionalSecretVersion#secret}
	Secret *string `field:"required" json:"secret" yaml:"secret"`
	// The secret data. Must be no larger than 64KiB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret_version#secret_data GoogleSecretManagerRegionalSecretVersion#secret_data}
	SecretData *string `field:"required" json:"secretData" yaml:"secretData"`
	// The deletion policy for the regional secret version.
	//
	// Setting 'ABANDON' allows the resource
	// to be abandoned rather than deleted. Setting 'DISABLE' allows the resource to be
	// disabled rather than deleted. Default is 'DELETE'. Possible values are:
	//   * DELETE
	//   * DISABLE
	//   * ABANDON
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret_version#deletion_policy GoogleSecretManagerRegionalSecretVersion#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// The current state of the regional secret version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret_version#enabled GoogleSecretManagerRegionalSecretVersion#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret_version#id GoogleSecretManagerRegionalSecretVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// If set to 'true', the secret data is expected to be base64-encoded string and would be sent as is.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret_version#is_secret_data_base64 GoogleSecretManagerRegionalSecretVersion#is_secret_data_base64}
	IsSecretDataBase64 interface{} `field:"optional" json:"isSecretDataBase64" yaml:"isSecretDataBase64"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret_version#timeouts GoogleSecretManagerRegionalSecretVersion#timeouts}
	Timeouts *GoogleSecretManagerRegionalSecretVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

