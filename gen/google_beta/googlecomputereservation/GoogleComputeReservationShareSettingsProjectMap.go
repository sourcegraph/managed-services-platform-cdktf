package googlecomputereservation


type GoogleComputeReservationShareSettingsProjectMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_reservation#id GoogleComputeReservation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The project id/number, should be same as the key of this project config in the project map.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_compute_reservation#project_id GoogleComputeReservation#project_id}
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
}

