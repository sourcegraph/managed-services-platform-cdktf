package bigqueryreservation


type BigqueryReservationAutoscale struct {
	// Number of slots to be scaled when needed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/bigquery_reservation#max_slots BigqueryReservation#max_slots}
	MaxSlots *float64 `field:"optional" json:"maxSlots" yaml:"maxSlots"`
}

