package googleessentialcontactscontact

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleEssentialContactsContactConfig struct {
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
	// The email address to send notifications to. This does not need to be a Google account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_essential_contacts_contact#email GoogleEssentialContactsContact#email}
	Email *string `field:"required" json:"email" yaml:"email"`
	// The preferred language for notifications, as a ISO 639-1 language code.
	//
	// See Supported languages for a list of supported languages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_essential_contacts_contact#language_tag GoogleEssentialContactsContact#language_tag}
	LanguageTag *string `field:"required" json:"languageTag" yaml:"languageTag"`
	// The categories of notifications that the contact will receive communications for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_essential_contacts_contact#notification_category_subscriptions GoogleEssentialContactsContact#notification_category_subscriptions}
	NotificationCategorySubscriptions *[]*string `field:"required" json:"notificationCategorySubscriptions" yaml:"notificationCategorySubscriptions"`
	// The resource to save this contact for. Format: organizations/{organization_id}, folders/{folder_id} or projects/{project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_essential_contacts_contact#parent GoogleEssentialContactsContact#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_essential_contacts_contact#id GoogleEssentialContactsContact#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_essential_contacts_contact#timeouts GoogleEssentialContactsContact#timeouts}
	Timeouts *GoogleEssentialContactsContactTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

