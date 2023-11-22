package googlecertificatemanagercertificatemap

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleCertificateManagerCertificateMapConfig struct {
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
	// A user-defined name of the Certificate Map. Certificate Map names must be unique globally and match the pattern 'projects/*\/locations/*\/certificateMaps/*'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_certificate_manager_certificate_map#name GoogleCertificateManagerCertificateMap#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A human-readable description of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_certificate_manager_certificate_map#description GoogleCertificateManagerCertificateMap#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_certificate_manager_certificate_map#id GoogleCertificateManagerCertificateMap#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set of labels associated with a Certificate Map resource.
	//
	// *Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_certificate_manager_certificate_map#labels GoogleCertificateManagerCertificateMap#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_certificate_manager_certificate_map#project GoogleCertificateManagerCertificateMap#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_certificate_manager_certificate_map#timeouts GoogleCertificateManagerCertificateMap#timeouts}
	Timeouts *GoogleCertificateManagerCertificateMapTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

