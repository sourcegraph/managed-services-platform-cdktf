package googleapigeekeystoresaliasespkcs12

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleApigeeKeystoresAliasesPkcs12Config struct {
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
	// Alias Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#alias GoogleApigeeKeystoresAliasesPkcs12#alias}
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// Environment associated with the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#environment GoogleApigeeKeystoresAliasesPkcs12#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Cert content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#file GoogleApigeeKeystoresAliasesPkcs12#file}
	File *string `field:"required" json:"file" yaml:"file"`
	// Hash of the pkcs file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#filehash GoogleApigeeKeystoresAliasesPkcs12#filehash}
	Filehash *string `field:"required" json:"filehash" yaml:"filehash"`
	// Keystore Name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#keystore GoogleApigeeKeystoresAliasesPkcs12#keystore}
	Keystore *string `field:"required" json:"keystore" yaml:"keystore"`
	// Organization ID associated with the alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#org_id GoogleApigeeKeystoresAliasesPkcs12#org_id}
	OrgId *string `field:"required" json:"orgId" yaml:"orgId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#id GoogleApigeeKeystoresAliasesPkcs12#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Password for the Private Key if it's encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#password GoogleApigeeKeystoresAliasesPkcs12#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.26.0/docs/resources/google_apigee_keystores_aliases_pkcs12#timeouts GoogleApigeeKeystoresAliasesPkcs12#timeouts}
	Timeouts *GoogleApigeeKeystoresAliasesPkcs12Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

