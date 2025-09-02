package googledataplexentry


type GoogleDataplexEntryAspectsAspect struct {
	// The content of the aspect in JSON form, according to its aspect type schema.
	//
	// The maximum size of the field is 120KB (encoded as UTF-8).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_dataplex_entry#data GoogleDataplexEntry#data}
	Data *string `field:"required" json:"data" yaml:"data"`
}

