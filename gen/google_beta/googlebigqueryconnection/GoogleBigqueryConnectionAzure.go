package googlebigqueryconnection


type GoogleBigqueryConnectionAzure struct {
	// The id of customer's directory that host the data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_connection#customer_tenant_id GoogleBigqueryConnection#customer_tenant_id}
	CustomerTenantId *string `field:"required" json:"customerTenantId" yaml:"customerTenantId"`
	// The Azure Application (client) ID where the federated credentials will be hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_connection#federated_application_client_id GoogleBigqueryConnection#federated_application_client_id}
	FederatedApplicationClientId *string `field:"optional" json:"federatedApplicationClientId" yaml:"federatedApplicationClientId"`
}

