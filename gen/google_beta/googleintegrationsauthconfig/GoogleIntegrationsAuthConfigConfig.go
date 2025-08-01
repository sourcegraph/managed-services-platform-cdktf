package googleintegrationsauthconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleIntegrationsAuthConfigConfig struct {
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
	// The name of the auth config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#display_name GoogleIntegrationsAuthConfig#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Location in which client needs to be provisioned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#location GoogleIntegrationsAuthConfig#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// client_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#client_certificate GoogleIntegrationsAuthConfig#client_certificate}
	ClientCertificate *GoogleIntegrationsAuthConfigClientCertificate `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// decrypted_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#decrypted_credential GoogleIntegrationsAuthConfig#decrypted_credential}
	DecryptedCredential *GoogleIntegrationsAuthConfigDecryptedCredential `field:"optional" json:"decryptedCredential" yaml:"decryptedCredential"`
	// A description of the auth config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#description GoogleIntegrationsAuthConfig#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User can define the time to receive notification after which the auth config becomes invalid.
	//
	// Support up to 30 days. Support granularity in hours.
	//
	// A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#expiry_notification_duration GoogleIntegrationsAuthConfig#expiry_notification_duration}
	ExpiryNotificationDuration *[]*string `field:"optional" json:"expiryNotificationDuration" yaml:"expiryNotificationDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#id GoogleIntegrationsAuthConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// User provided expiry time to override.
	//
	// For the example of Salesforce, username/password credentials can be valid for 6 months depending on the instance settings.
	//
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#override_valid_time GoogleIntegrationsAuthConfig#override_valid_time}
	OverrideValidTime *string `field:"optional" json:"overrideValidTime" yaml:"overrideValidTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#project GoogleIntegrationsAuthConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#timeouts GoogleIntegrationsAuthConfig#timeouts}
	Timeouts *GoogleIntegrationsAuthConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The visibility of the auth config. Possible values: ["PRIVATE", "CLIENT_VISIBLE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_integrations_auth_config#visibility GoogleIntegrationsAuthConfig#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

