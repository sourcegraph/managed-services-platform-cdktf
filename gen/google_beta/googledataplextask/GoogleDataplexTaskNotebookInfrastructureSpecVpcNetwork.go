package googledataplextask


type GoogleDataplexTaskNotebookInfrastructureSpecVpcNetwork struct {
	// The Cloud VPC network in which the job is run.
	//
	// By default, the Cloud VPC network named Default within the project is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_task#network GoogleDataplexTask#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// List of network tags to apply to the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_task#network_tags GoogleDataplexTask#network_tags}
	NetworkTags *[]*string `field:"optional" json:"networkTags" yaml:"networkTags"`
	// The Cloud VPC sub-network in which the job is run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_task#sub_network GoogleDataplexTask#sub_network}
	SubNetwork *string `field:"optional" json:"subNetwork" yaml:"subNetwork"`
}

