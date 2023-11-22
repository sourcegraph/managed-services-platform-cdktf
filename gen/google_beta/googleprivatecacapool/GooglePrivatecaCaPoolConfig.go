package googleprivatecacapool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePrivatecaCaPoolConfig struct {
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
	// Location of the CaPool. A full list of valid locations can be found by running 'gcloud privateca locations list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#location GooglePrivatecaCaPool#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name for this CaPool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#name GooglePrivatecaCaPool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Tier of this CaPool. Possible values: ["ENTERPRISE", "DEVOPS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#tier GooglePrivatecaCaPool#tier}
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#id GooglePrivatecaCaPool#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// issuance_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#issuance_policy GooglePrivatecaCaPool#issuance_policy}
	IssuancePolicy *GooglePrivatecaCaPoolIssuancePolicy `field:"optional" json:"issuancePolicy" yaml:"issuancePolicy"`
	// Labels with user-defined metadata.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	//
	// *Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#labels GooglePrivatecaCaPool#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#project GooglePrivatecaCaPool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#publishing_options GooglePrivatecaCaPool#publishing_options}
	PublishingOptions *GooglePrivatecaCaPoolPublishingOptions `field:"optional" json:"publishingOptions" yaml:"publishingOptions"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.7.0/docs/resources/google_privateca_ca_pool#timeouts GooglePrivatecaCaPool#timeouts}
	Timeouts *GooglePrivatecaCaPoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

