package googlecontaineranalysisoccurrence


type GoogleContainerAnalysisOccurrenceAttestation struct {
	// The serialized payload that is verified by one or more signatures. A base64-encoded string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_container_analysis_occurrence#serialized_payload GoogleContainerAnalysisOccurrence#serialized_payload}
	SerializedPayload *string `field:"required" json:"serializedPayload" yaml:"serializedPayload"`
	// signatures block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_container_analysis_occurrence#signatures GoogleContainerAnalysisOccurrence#signatures}
	Signatures interface{} `field:"required" json:"signatures" yaml:"signatures"`
}

