package googlepubsubsubscription


type GooglePubsubSubscriptionPushConfigOidcToken struct {
	// Service account email to be used for generating the OIDC token.
	//
	// The caller (for subscriptions.create, subscriptions.patch, and
	// subscriptions.modifyPushConfig RPCs) must have the
	// iam.serviceAccounts.actAs permission for the service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_pubsub_subscription#service_account_email GooglePubsubSubscription#service_account_email}
	ServiceAccountEmail *string `field:"required" json:"serviceAccountEmail" yaml:"serviceAccountEmail"`
	// Audience to be used when generating OIDC token.
	//
	// The audience claim
	// identifies the recipients that the JWT is intended for. The audience
	// value is a single case-sensitive string. Having multiple values (array)
	// for the audience field is not supported. More info about the OIDC JWT
	// token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
	// Note: if not specified, the Push endpoint URL will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.15.0/docs/resources/google_pubsub_subscription#audience GooglePubsubSubscription#audience}
	Audience *string `field:"optional" json:"audience" yaml:"audience"`
}

