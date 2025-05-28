package googleprivatecacertificateauthority

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePrivatecaCertificateAuthorityConfig struct {
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
	// The user provided Resource ID for this Certificate Authority.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#certificate_authority_id GooglePrivatecaCertificateAuthority#certificate_authority_id}
	CertificateAuthorityId *string `field:"required" json:"certificateAuthorityId" yaml:"certificateAuthorityId"`
	// config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#config GooglePrivatecaCertificateAuthority#config}
	Config *GooglePrivatecaCertificateAuthorityConfigA `field:"required" json:"config" yaml:"config"`
	// key_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#key_spec GooglePrivatecaCertificateAuthority#key_spec}
	KeySpec *GooglePrivatecaCertificateAuthorityKeySpec `field:"required" json:"keySpec" yaml:"keySpec"`
	// Location of the CertificateAuthority. A full list of valid locations can be found by running 'gcloud privateca locations list'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#location GooglePrivatecaCertificateAuthority#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The name of the CaPool this Certificate Authority belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#pool GooglePrivatecaCertificateAuthority#pool}
	Pool *string `field:"required" json:"pool" yaml:"pool"`
	// Whether Terraform will be prevented from destroying the CertificateAuthority.
	//
	// When the field is set to true or unset in Terraform state, a 'terraform apply'
	// or 'terraform destroy' that would delete the CertificateAuthority will fail.
	// When the field is set to false, deleting the CertificateAuthority is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#deletion_protection GooglePrivatecaCertificateAuthority#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Desired state of the CertificateAuthority.
	//
	// Set this field to 'STAGED' to create a 'STAGED' root CA.
	// Possible values: ENABLED, DISABLED, STAGED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#desired_state GooglePrivatecaCertificateAuthority#desired_state}
	DesiredState *string `field:"optional" json:"desiredState" yaml:"desiredState"`
	// The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs.
	//
	// This must be a bucket name, without any prefixes
	// (such as 'gs://') or suffixes (such as '.googleapis.com'). For example, to use a bucket named
	// my-bucket, you would simply specify 'my-bucket'. If not specified, a managed bucket will be
	// created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#gcs_bucket GooglePrivatecaCertificateAuthority#gcs_bucket}
	GcsBucket *string `field:"optional" json:"gcsBucket" yaml:"gcsBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#id GooglePrivatecaCertificateAuthority#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// This field allows the CA to be deleted even if the CA has active certs.
	//
	// Active certs include both unrevoked and unexpired certs.
	// Use with care. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#ignore_active_certificates_on_deletion GooglePrivatecaCertificateAuthority#ignore_active_certificates_on_deletion}
	IgnoreActiveCertificatesOnDeletion interface{} `field:"optional" json:"ignoreActiveCertificatesOnDeletion" yaml:"ignoreActiveCertificatesOnDeletion"`
	// Labels with user-defined metadata.
	//
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }.
	//
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#labels GooglePrivatecaCertificateAuthority#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The desired lifetime of the CA certificate.
	//
	// Used to create the "notBeforeTime" and
	// "notAfterTime" fields inside an X.509 certificate. A duration in seconds with up to nine
	// fractional digits, terminated by 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#lifetime GooglePrivatecaCertificateAuthority#lifetime}
	Lifetime *string `field:"optional" json:"lifetime" yaml:"lifetime"`
	// The signed CA certificate issued from the subordinated CA's CSR.
	//
	// This is needed when activating the subordiante CA with a third party issuer.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#pem_ca_certificate GooglePrivatecaCertificateAuthority#pem_ca_certificate}
	PemCaCertificate *string `field:"optional" json:"pemCaCertificate" yaml:"pemCaCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#project GooglePrivatecaCertificateAuthority#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// If this flag is set, the Certificate Authority will be deleted as soon as possible without a 30-day grace period where undeletion would have been allowed.
	//
	// If you proceed, there will be no way to recover this CA.
	// Use with care. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#skip_grace_period GooglePrivatecaCertificateAuthority#skip_grace_period}
	SkipGracePeriod interface{} `field:"optional" json:"skipGracePeriod" yaml:"skipGracePeriod"`
	// subordinate_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#subordinate_config GooglePrivatecaCertificateAuthority#subordinate_config}
	SubordinateConfig *GooglePrivatecaCertificateAuthoritySubordinateConfig `field:"optional" json:"subordinateConfig" yaml:"subordinateConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#timeouts GooglePrivatecaCertificateAuthority#timeouts}
	Timeouts *GooglePrivatecaCertificateAuthorityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The Type of this CertificateAuthority.
	//
	// ~> **Note:** For 'SUBORDINATE' Certificate Authorities, they need to
	// be activated before they can issue certificates. Default value: "SELF_SIGNED" Possible values: ["SELF_SIGNED", "SUBORDINATE"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#type GooglePrivatecaCertificateAuthority#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// user_defined_access_urls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_privateca_certificate_authority#user_defined_access_urls GooglePrivatecaCertificateAuthority#user_defined_access_urls}
	UserDefinedAccessUrls *GooglePrivatecaCertificateAuthorityUserDefinedAccessUrls `field:"optional" json:"userDefinedAccessUrls" yaml:"userDefinedAccessUrls"`
}

