package googlesecretmanagerregionalsecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleSecretManagerRegionalSecretConfig struct {
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
	// The location of the regional secret. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#location GoogleSecretManagerRegionalSecret#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// This must be unique within the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#secret_id GoogleSecretManagerRegionalSecret#secret_id}
	SecretId *string `field:"required" json:"secretId" yaml:"secretId"`
	// Custom metadata about the regional secret.
	//
	// Annotations are distinct from various forms of labels. Annotations exist to allow
	// client tools to store their own state information without requiring a database.
	//
	// Annotation keys must be between 1 and 63 characters long, have a UTF-8 encoding of
	// maximum 128 bytes, begin and end with an alphanumeric character ([a-z0-9A-Z]), and
	// may have dashes (-), underscores (_), dots (.), and alphanumerics in between these
	// symbols.
	//
	// The total size of annotation keys and values must be less than 16KiB.
	//
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#annotations GoogleSecretManagerRegionalSecret#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// customer_managed_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#customer_managed_encryption GoogleSecretManagerRegionalSecret#customer_managed_encryption}
	CustomerManagedEncryption *GoogleSecretManagerRegionalSecretCustomerManagedEncryption `field:"optional" json:"customerManagedEncryption" yaml:"customerManagedEncryption"`
	// Whether Terraform will be prevented from destroying the regional secret.
	//
	// Defaults to false.
	// When the field is set to true in Terraform state, a 'terraform apply'
	// or 'terraform destroy' that would delete the federation will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#deletion_protection GoogleSecretManagerRegionalSecret#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Timestamp in UTC when the regional secret is scheduled to expire.
	//
	// This is always provided on
	// output, regardless of what was sent on input. A timestamp in RFC3339 UTC "Zulu" format, with
	// nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and
	// "2014-10-02T15:01:23.045123456Z". Only one of 'expire_time' or 'ttl' can be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#expire_time GoogleSecretManagerRegionalSecret#expire_time}
	ExpireTime *string `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#id GoogleSecretManagerRegionalSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The labels assigned to this regional secret.
	//
	// Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	// and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
	//
	// Label values must be between 0 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes,
	// and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
	//
	// No more than 64 labels can be assigned to a given resource.
	//
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#labels GoogleSecretManagerRegionalSecret#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#project GoogleSecretManagerRegionalSecret#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// rotation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#rotation GoogleSecretManagerRegionalSecret#rotation}
	Rotation *GoogleSecretManagerRegionalSecretRotation `field:"optional" json:"rotation" yaml:"rotation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#timeouts GoogleSecretManagerRegionalSecret#timeouts}
	Timeouts *GoogleSecretManagerRegionalSecretTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// topics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#topics GoogleSecretManagerRegionalSecret#topics}
	Topics interface{} `field:"optional" json:"topics" yaml:"topics"`
	// The TTL for the regional secret.
	//
	// A duration in seconds with up to nine fractional digits,
	// terminated by 's'. Example: "3.5s". Only one of 'ttl' or 'expire_time' can be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#ttl GoogleSecretManagerRegionalSecret#ttl}
	Ttl *string `field:"optional" json:"ttl" yaml:"ttl"`
	// Mapping from version alias to version name.
	//
	// A version alias is a string with a maximum length of 63 characters and can contain
	// uppercase and lowercase letters, numerals, and the hyphen (-) and underscore ('_')
	// characters. An alias string must start with a letter and cannot be the string
	// 'latest' or 'NEW'. No more than 50 aliases can be assigned to a given secret.
	//
	// An object containing a list of "key": value pairs. Example:
	// { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#version_aliases GoogleSecretManagerRegionalSecret#version_aliases}
	VersionAliases *map[string]*string `field:"optional" json:"versionAliases" yaml:"versionAliases"`
	// Secret Version TTL after destruction request.
	//
	// This is a part of the delayed delete feature on Secret Version.
	// For secret with versionDestroyTtl>0, version destruction doesn't happen immediately
	// on calling destroy instead the version goes to a disabled state and
	// the actual destruction happens after this TTL expires. It must be atleast 24h.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_secret_manager_regional_secret#version_destroy_ttl GoogleSecretManagerRegionalSecret#version_destroy_ttl}
	VersionDestroyTtl *string `field:"optional" json:"versionDestroyTtl" yaml:"versionDestroyTtl"`
}

