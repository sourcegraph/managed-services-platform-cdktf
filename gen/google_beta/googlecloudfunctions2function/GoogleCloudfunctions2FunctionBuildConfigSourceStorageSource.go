package googlecloudfunctions2function


type GoogleCloudfunctions2FunctionBuildConfigSourceStorageSource struct {
	// Google Cloud Storage bucket containing the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloudfunctions2_function#bucket GoogleCloudfunctions2Function#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Google Cloud Storage generation for the object. If the generation is omitted, the latest generation will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloudfunctions2_function#generation GoogleCloudfunctions2Function#generation}
	Generation *float64 `field:"optional" json:"generation" yaml:"generation"`
	// Google Cloud Storage object containing the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.36.1/docs/resources/google_cloudfunctions2_function#object GoogleCloudfunctions2Function#object}
	Object *string `field:"optional" json:"object" yaml:"object"`
}

