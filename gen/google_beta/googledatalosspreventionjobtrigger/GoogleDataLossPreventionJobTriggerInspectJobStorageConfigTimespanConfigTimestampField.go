package googledatalosspreventionjobtrigger


type GoogleDataLossPreventionJobTriggerInspectJobStorageConfigTimespanConfigTimestampField struct {
	// Specification of the field containing the timestamp of scanned items. Used for data sources like Datastore and BigQuery.
	//
	// For BigQuery: Required to filter out rows based on the given start and end times. If not specified and the table was
	// modified between the given start and end times, the entire table will be scanned. The valid data types of the timestamp
	// field are: INTEGER, DATE, TIMESTAMP, or DATETIME BigQuery column.
	//
	// For Datastore. Valid data types of the timestamp field are: TIMESTAMP. Datastore entity will be scanned if the
	// timestamp property does not exist or its value is empty or invalid.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_loss_prevention_job_trigger#name GoogleDataLossPreventionJobTrigger#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

