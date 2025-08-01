package googlechronicledataaccessscope


type GoogleChronicleDataAccessScopeAllowedDataAccessLabelsIngestionLabel struct {
	// Required. The key of the ingestion label. Always required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_chronicle_data_access_scope#ingestion_label_key GoogleChronicleDataAccessScope#ingestion_label_key}
	IngestionLabelKey *string `field:"required" json:"ingestionLabelKey" yaml:"ingestionLabelKey"`
	// Optional.
	//
	// The value of the ingestion label. Optional. An object
	// with no provided value and some key provided would match
	// against the given key and ANY value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_chronicle_data_access_scope#ingestion_label_value GoogleChronicleDataAccessScope#ingestion_label_value}
	IngestionLabelValue *string `field:"optional" json:"ingestionLabelValue" yaml:"ingestionLabelValue"`
}

