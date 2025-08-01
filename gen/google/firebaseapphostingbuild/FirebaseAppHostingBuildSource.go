package firebaseapphostingbuild


type FirebaseAppHostingBuildSource struct {
	// codebase block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firebase_app_hosting_build#codebase FirebaseAppHostingBuild#codebase}
	Codebase *FirebaseAppHostingBuildSourceCodebase `field:"optional" json:"codebase" yaml:"codebase"`
	// container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firebase_app_hosting_build#container FirebaseAppHostingBuild#container}
	Container *FirebaseAppHostingBuildSourceContainer `field:"optional" json:"container" yaml:"container"`
}

