package computereservation


type ComputeReservationSpecificReservation struct {
	// The number of resources that are allocated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_reservation#count ComputeReservation#count}
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// instance_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_reservation#instance_properties ComputeReservation#instance_properties}
	InstanceProperties *ComputeReservationSpecificReservationInstanceProperties `field:"optional" json:"instanceProperties" yaml:"instanceProperties"`
	// Specifies the instance template to create the reservation. If you use this field, you must exclude the instanceProperties field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/compute_reservation#source_instance_template ComputeReservation#source_instance_template}
	SourceInstanceTemplate *string `field:"optional" json:"sourceInstanceTemplate" yaml:"sourceInstanceTemplate"`
}

