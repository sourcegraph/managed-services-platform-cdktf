package googleeventarctrigger


type GoogleEventarcTriggerDestinationHttpEndpoint struct {
	// Required.
	//
	// The URI of the HTTP enpdoint. The value must be a RFC2396 URI string. Examples: 'http://10.10.10.8:80/route', 'http://svc.us-central1.p.local:8080/'. Only HTTP and HTTPS protocols are supported. The host can be either a static IP addressable from the VPC specified by the network config, or an internal DNS hostname of the service resolvable via Cloud DNS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_eventarc_trigger#uri GoogleEventarcTrigger#uri}
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

