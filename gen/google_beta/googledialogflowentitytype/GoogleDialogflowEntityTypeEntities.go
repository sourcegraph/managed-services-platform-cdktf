package googledialogflowentitytype


type GoogleDialogflowEntityTypeEntities struct {
	// A collection of value synonyms.
	//
	// For example, if the entity type is vegetable, and value is scallions, a synonym
	// could be green onions.
	// For KIND_LIST entity types:
	// * This collection must contain exactly one synonym equal to value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_entity_type#synonyms GoogleDialogflowEntityType#synonyms}
	Synonyms *[]*string `field:"required" json:"synonyms" yaml:"synonyms"`
	// The primary value associated with this entity entry.
	//
	// For example, if the entity type is vegetable, the value
	// could be scallions.
	// For KIND_MAP entity types:
	// * A reference value to be used in place of synonyms.
	// For KIND_LIST entity types:
	// * A string that can contain references to other entity types (with or without aliases).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dialogflow_entity_type#value GoogleDialogflowEntityType#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

