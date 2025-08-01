package firebaseapphostingtraffic


type FirebaseAppHostingTrafficTarget struct {
	// splits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/firebase_app_hosting_traffic#splits FirebaseAppHostingTraffic#splits}
	Splits interface{} `field:"required" json:"splits" yaml:"splits"`
}

