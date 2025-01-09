package samlsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SamlSettingsConfig struct {
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
	// Identity Provider Certificate specifies the PEM encoded X.509 Certificate as provided by the IdP configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#idp_cert SamlSettings#idp_cert}
	IdpCert *string `field:"required" json:"idpCert" yaml:"idpCert"`
	// Single Log Out URL specifies the HTTPS endpoint on your IdP for single logout requests.
	//
	// This value is provided by the IdP configuration
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#slo_endpoint_url SamlSettings#slo_endpoint_url}
	SloEndpointUrl *string `field:"required" json:"sloEndpointUrl" yaml:"sloEndpointUrl"`
	// Single Sign On URL specifies the HTTPS endpoint on your IdP for single sign-on requests.
	//
	// This value is provided by the IdP configuration
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#sso_endpoint_url SamlSettings#sso_endpoint_url}
	SsoEndpointUrl *string `field:"required" json:"ssoEndpointUrl" yaml:"ssoEndpointUrl"`
	// Team Attribute Name specifies the name of the SAML attribute that determines team membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#attr_groups SamlSettings#attr_groups}
	AttrGroups *string `field:"optional" json:"attrGroups" yaml:"attrGroups"`
	// Specifies the role for site admin access. Overrides the "Site Admin Role" method.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#attr_site_admin SamlSettings#attr_site_admin}
	AttrSiteAdmin *string `field:"optional" json:"attrSiteAdmin" yaml:"attrSiteAdmin"`
	// Username Attribute Name specifies the name of the SAML attribute that determines the user's username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#attr_username SamlSettings#attr_username}
	AttrUsername *string `field:"optional" json:"attrUsername" yaml:"attrUsername"`
	// Ensure that <samlp:AuthnRequest> messages are signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#authn_requests_signed SamlSettings#authn_requests_signed}
	AuthnRequestsSigned interface{} `field:"optional" json:"authnRequestsSigned" yaml:"authnRequestsSigned"`
	// The certificate used for request and assertion signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#certificate SamlSettings#certificate}
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// When sign-on fails and this is enabled, the SAMLResponse XML will be displayed on the login page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#debug SamlSettings#debug}
	Debug interface{} `field:"optional" json:"debug" yaml:"debug"`
	// The private key used for request and assertion signing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#private_key SamlSettings#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Signature Digest Method. Must be either `SHA1` or `SHA256`. Defaults to `SHA256`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#signature_digest_method SamlSettings#signature_digest_method}
	SignatureDigestMethod *string `field:"optional" json:"signatureDigestMethod" yaml:"signatureDigestMethod"`
	// Signature Signing Method. Must be either `SHA1` or `SHA256`. Defaults to `SHA256`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#signature_signing_method SamlSettings#signature_signing_method}
	SignatureSigningMethod *string `field:"optional" json:"signatureSigningMethod" yaml:"signatureSigningMethod"`
	// Specifies the role for site admin access, provided in the list of roles sent in the Team Attribute Name attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#site_admin_role SamlSettings#site_admin_role}
	SiteAdminRole *string `field:"optional" json:"siteAdminRole" yaml:"siteAdminRole"`
	// Specifies the Single Sign On session timeout in seconds. Defaults to 14 days.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#sso_api_token_session_timeout SamlSettings#sso_api_token_session_timeout}
	SsoApiTokenSessionTimeout *float64 `field:"optional" json:"ssoApiTokenSessionTimeout" yaml:"ssoApiTokenSessionTimeout"`
	// Set it to false if you would rather use Terraform Enterprise to manage team membership.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#team_management_enabled SamlSettings#team_management_enabled}
	TeamManagementEnabled interface{} `field:"optional" json:"teamManagementEnabled" yaml:"teamManagementEnabled"`
	// Ensure that <saml:Assertion> elements are signed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.62.0/docs/resources/saml_settings#want_assertions_signed SamlSettings#want_assertions_signed}
	WantAssertionsSigned interface{} `field:"optional" json:"wantAssertionsSigned" yaml:"wantAssertionsSigned"`
}

