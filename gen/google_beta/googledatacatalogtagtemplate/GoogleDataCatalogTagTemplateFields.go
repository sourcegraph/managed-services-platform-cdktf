package googledatacatalogtagtemplate


type GoogleDataCatalogTagTemplateFields struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_catalog_tag_template#field_id GoogleDataCatalogTagTemplate#field_id}.
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_catalog_tag_template#type GoogleDataCatalogTagTemplate#type}
	Type *GoogleDataCatalogTagTemplateFieldsType `field:"required" json:"type" yaml:"type"`
	// A description for this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_catalog_tag_template#description GoogleDataCatalogTagTemplate#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name for this field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_catalog_tag_template#display_name GoogleDataCatalogTagTemplate#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Whether this is a required field. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_catalog_tag_template#is_required GoogleDataCatalogTagTemplate#is_required}
	IsRequired interface{} `field:"optional" json:"isRequired" yaml:"isRequired"`
	// The order of this field with respect to other fields in this tag template.
	//
	// A higher value indicates a more important field. The value can be negative.
	// Multiple fields can have the same order, and field orders within a tag do not have to be sequential.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.34.0/docs/resources/google_data_catalog_tag_template#order GoogleDataCatalogTagTemplate#order}
	Order *float64 `field:"optional" json:"order" yaml:"order"`
}

