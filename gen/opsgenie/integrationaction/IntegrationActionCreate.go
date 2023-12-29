package integrationaction


type IntegrationActionCreate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#name IntegrationAction#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#alert_actions IntegrationAction#alert_actions}.
	AlertActions *[]*string `field:"optional" json:"alertActions" yaml:"alertActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#alias IntegrationAction#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#append_attachments IntegrationAction#append_attachments}.
	AppendAttachments interface{} `field:"optional" json:"appendAttachments" yaml:"appendAttachments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#custom_priority IntegrationAction#custom_priority}.
	CustomPriority *string `field:"optional" json:"customPriority" yaml:"customPriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#description IntegrationAction#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#entity IntegrationAction#entity}.
	Entity *string `field:"optional" json:"entity" yaml:"entity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#extra_properties IntegrationAction#extra_properties}.
	ExtraProperties *map[string]*string `field:"optional" json:"extraProperties" yaml:"extraProperties"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#filter IntegrationAction#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#ignore_alert_actions_from_payload IntegrationAction#ignore_alert_actions_from_payload}.
	IgnoreAlertActionsFromPayload interface{} `field:"optional" json:"ignoreAlertActionsFromPayload" yaml:"ignoreAlertActionsFromPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#ignore_extra_properties_from_payload IntegrationAction#ignore_extra_properties_from_payload}.
	IgnoreExtraPropertiesFromPayload interface{} `field:"optional" json:"ignoreExtraPropertiesFromPayload" yaml:"ignoreExtraPropertiesFromPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#ignore_responders_from_payload IntegrationAction#ignore_responders_from_payload}.
	IgnoreRespondersFromPayload interface{} `field:"optional" json:"ignoreRespondersFromPayload" yaml:"ignoreRespondersFromPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#ignore_tags_from_payload IntegrationAction#ignore_tags_from_payload}.
	IgnoreTagsFromPayload interface{} `field:"optional" json:"ignoreTagsFromPayload" yaml:"ignoreTagsFromPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#ignore_teams_from_payload IntegrationAction#ignore_teams_from_payload}.
	IgnoreTeamsFromPayload interface{} `field:"optional" json:"ignoreTeamsFromPayload" yaml:"ignoreTeamsFromPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#message IntegrationAction#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#note IntegrationAction#note}.
	Note *string `field:"optional" json:"note" yaml:"note"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#order IntegrationAction#order}.
	Order *float64 `field:"optional" json:"order" yaml:"order"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#priority IntegrationAction#priority}.
	Priority *string `field:"optional" json:"priority" yaml:"priority"`
	// responders block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#responders IntegrationAction#responders}
	Responders interface{} `field:"optional" json:"responders" yaml:"responders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#source IntegrationAction#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#tags IntegrationAction#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#type IntegrationAction#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.35/docs/resources/integration_action#user IntegrationAction#user}.
	User *string `field:"optional" json:"user" yaml:"user"`
}

