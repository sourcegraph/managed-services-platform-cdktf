package googlepubsublitesubscription

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GooglePubsubLiteSubscriptionConfig struct {
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
	// Name of the subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#name GooglePubsubLiteSubscription#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A reference to a Topic resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#topic GooglePubsubLiteSubscription#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// delivery_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#delivery_config GooglePubsubLiteSubscription#delivery_config}
	DeliveryConfig *GooglePubsubLiteSubscriptionDeliveryConfig `field:"optional" json:"deliveryConfig" yaml:"deliveryConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#id GooglePubsubLiteSubscription#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#project GooglePubsubLiteSubscription#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the pubsub lite topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#region GooglePubsubLiteSubscription#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#timeouts GooglePubsubLiteSubscription#timeouts}
	Timeouts *GooglePubsubLiteSubscriptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The zone of the pubsub lite topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_lite_subscription#zone GooglePubsubLiteSubscription#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

