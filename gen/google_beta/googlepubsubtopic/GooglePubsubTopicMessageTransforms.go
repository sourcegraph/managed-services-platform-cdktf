package googlepubsubtopic


type GooglePubsubTopicMessageTransforms struct {
	// Controls whether or not to use this transform.
	//
	// If not set or 'false',
	// the transform will be applied to messages. Default: 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#disabled GooglePubsubTopic#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// javascript_udf block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_pubsub_topic#javascript_udf GooglePubsubTopic#javascript_udf}
	JavascriptUdf *GooglePubsubTopicMessageTransformsJavascriptUdf `field:"optional" json:"javascriptUdf" yaml:"javascriptUdf"`
}

