package googlevertexaiindexendpointdeployedindex


type GoogleVertexAiIndexEndpointDeployedIndexDeployedIndexAuthConfigAuthProvider struct {
	// A list of allowed JWT issuers.
	//
	// Each entry must be a valid Google service account, in the following format: service-account-name@project-id.iam.gserviceaccount.com
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index#allowed_issuers GoogleVertexAiIndexEndpointDeployedIndex#allowed_issuers}
	AllowedIssuers *[]*string `field:"optional" json:"allowedIssuers" yaml:"allowedIssuers"`
	// The list of JWT audiences.
	//
	// that are allowed to access. A JWT containing any of these audiences will be accepted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_vertex_ai_index_endpoint_deployed_index#audiences GoogleVertexAiIndexEndpointDeployedIndex#audiences}
	Audiences *[]*string `field:"optional" json:"audiences" yaml:"audiences"`
}

