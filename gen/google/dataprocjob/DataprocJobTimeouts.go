package dataprocjob


type DataprocJobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_job#create DataprocJob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/dataproc_job#delete DataprocJob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

