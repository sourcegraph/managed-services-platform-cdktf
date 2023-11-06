package googlekmscryptokey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleKmsCryptoKeyConfig struct {
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
	// The KeyRing that this key belongs to. Format: ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#key_ring GoogleKmsCryptoKey#key_ring}
	KeyRing *string `field:"required" json:"keyRing" yaml:"keyRing"`
	// The resource name for the CryptoKey.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#name GoogleKmsCryptoKey#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The period of time that versions of this key spend in the DESTROY_SCHEDULED state before transitioning to DESTROYED.
	//
	// If not specified at creation time, the default duration is 24 hours.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#destroy_scheduled_duration GoogleKmsCryptoKey#destroy_scheduled_duration}
	DestroyScheduledDuration *string `field:"optional" json:"destroyScheduledDuration" yaml:"destroyScheduledDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#id GoogleKmsCryptoKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether this key may contain imported versions only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#import_only GoogleKmsCryptoKey#import_only}
	ImportOnly interface{} `field:"optional" json:"importOnly" yaml:"importOnly"`
	// Labels with user-defined metadata to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#labels GoogleKmsCryptoKey#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The immutable purpose of this CryptoKey. See the [purpose reference](https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.cryptoKeys#CryptoKeyPurpose) for possible inputs. Default value is "ENCRYPT_DECRYPT".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#purpose GoogleKmsCryptoKey#purpose}
	Purpose *string `field:"optional" json:"purpose" yaml:"purpose"`
	// Every time this period passes, generate a new CryptoKeyVersion and set it as the primary.
	//
	// The first rotation will take place after the specified period. The rotation period has
	// the format of a decimal number with up to 9 fractional digits, followed by the
	// letter 's' (seconds). It must be greater than a day (ie, 86400).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#rotation_period GoogleKmsCryptoKey#rotation_period}
	RotationPeriod *string `field:"optional" json:"rotationPeriod" yaml:"rotationPeriod"`
	// If set to true, the request will create a CryptoKey without any CryptoKeyVersions.
	//
	// You must use the 'google_kms_key_ring_import_job' resource to import the CryptoKeyVersion.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#skip_initial_version_creation GoogleKmsCryptoKey#skip_initial_version_creation}
	SkipInitialVersionCreation interface{} `field:"optional" json:"skipInitialVersionCreation" yaml:"skipInitialVersionCreation"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#timeouts GoogleKmsCryptoKey#timeouts}
	Timeouts *GoogleKmsCryptoKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// version_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_kms_crypto_key#version_template GoogleKmsCryptoKey#version_template}
	VersionTemplate *GoogleKmsCryptoKeyVersionTemplate `field:"optional" json:"versionTemplate" yaml:"versionTemplate"`
}

