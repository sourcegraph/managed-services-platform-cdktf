package containeranalysisnote


type ContainerAnalysisNoteAttestationAuthority struct {
	// hint block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.15.0/docs/resources/container_analysis_note#hint ContainerAnalysisNote#hint}
	Hint *ContainerAnalysisNoteAttestationAuthorityHint `field:"required" json:"hint" yaml:"hint"`
}

