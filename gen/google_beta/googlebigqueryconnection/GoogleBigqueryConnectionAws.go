package googlebigqueryconnection


type GoogleBigqueryConnectionAws struct {
	// access_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_connection#access_role GoogleBigqueryConnection#access_role}
	AccessRole *GoogleBigqueryConnectionAwsAccessRole `field:"required" json:"accessRole" yaml:"accessRole"`
}

