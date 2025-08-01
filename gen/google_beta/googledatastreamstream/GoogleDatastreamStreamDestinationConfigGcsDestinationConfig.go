package googledatastreamstream


type GoogleDatastreamStreamDestinationConfigGcsDestinationConfig struct {
	// avro_file_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#avro_file_format GoogleDatastreamStream#avro_file_format}
	AvroFileFormat *GoogleDatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat `field:"optional" json:"avroFileFormat" yaml:"avroFileFormat"`
	// The maximum duration for which new events are added before a file is closed and a new file is created.
	//
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s". Defaults to 900s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#file_rotation_interval GoogleDatastreamStream#file_rotation_interval}
	FileRotationInterval *string `field:"optional" json:"fileRotationInterval" yaml:"fileRotationInterval"`
	// The maximum file size to be saved in the bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#file_rotation_mb GoogleDatastreamStream#file_rotation_mb}
	FileRotationMb *float64 `field:"optional" json:"fileRotationMb" yaml:"fileRotationMb"`
	// json_file_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#json_file_format GoogleDatastreamStream#json_file_format}
	JsonFileFormat *GoogleDatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat `field:"optional" json:"jsonFileFormat" yaml:"jsonFileFormat"`
	// Path inside the Cloud Storage bucket to write data to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_datastream_stream#path GoogleDatastreamStream#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

