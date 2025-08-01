package googleeventarctrigger


type GoogleEventarcTriggerMatchingCriteria struct {
	// Required.
	//
	// The name of a CloudEvents attribute. Currently, only a subset of attributes are supported for filtering. All triggers MUST provide a filter for the 'type' attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_trigger#attribute GoogleEventarcTrigger#attribute}
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// Required. The value for the attribute. See https://cloud.google.com/eventarc/docs/creating-triggers#trigger-gcloud for available values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_trigger#value GoogleEventarcTrigger#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Optional.
	//
	// The operator used for matching the events with the value of the filter. If not specified, only events that have an exact key-value pair specified in the filter are matched. The only allowed value is 'match-path-pattern'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_trigger#operator GoogleEventarcTrigger#operator}
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

