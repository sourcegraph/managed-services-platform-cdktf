package teamnotificationconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamNotificationConfigurationConfig struct {
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
	// The type of notification configuration payload to send.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#destination_type TeamNotificationConfiguration#destination_type}
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Name of the team notification configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#name TeamNotificationConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the team that owns the notification configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#team_id TeamNotificationConfiguration#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// A list of email addresses. This value must not be provided if `destination_type` is `generic`, `microsoft-teams`, or `slack`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#email_addresses TeamNotificationConfiguration#email_addresses}
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// A list of user IDs. This value must not be provided if `destination_type` is `generic`, `microsoft-teams`, or `slack`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#email_user_ids TeamNotificationConfiguration#email_user_ids}
	EmailUserIds *[]*string `field:"optional" json:"emailUserIds" yaml:"emailUserIds"`
	// Whether the team notification configuration should be enabled or not.
	//
	// Disabled configurations will not send any notifications. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#enabled TeamNotificationConfiguration#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// A write-only secure token for the notification configuration, which can be used by the receiving server to verify request authenticity when configured for notification configurations with a destination type of `generic`.
	//
	// Defaults to `null`. This value _must not_ be provided if `destination_type` is `email`, `microsoft-teams`, or `slack`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#token TeamNotificationConfiguration#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// A write-only secure token for the notification configuration, guaranteed not to be written to plan or state artifacts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#token_wo TeamNotificationConfiguration#token_wo}
	TokenWo *string `field:"optional" json:"tokenWo" yaml:"tokenWo"`
	// The array of triggers for which this team notification configuration will send notifications.
	//
	// If omitted, no notification triggers are configured.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#triggers TeamNotificationConfiguration#triggers}
	Triggers *[]*string `field:"optional" json:"triggers" yaml:"triggers"`
	// The HTTP or HTTPS URL where notification requests will be made.
	//
	// This value must not be provided if `email_addresses` or `email_user_ids` is present, or if `destination_type` is `email`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/tfe/0.66.0/docs/resources/team_notification_configuration#url TeamNotificationConfiguration#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
}

