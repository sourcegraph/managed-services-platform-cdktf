package zerotrustaccessapplication


type ZeroTrustAccessApplicationPoliciesConnectionRulesRdp struct {
	// Clipboard formats allowed when copying from local machine to remote RDP session.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#allowed_clipboard_local_to_remote_formats ZeroTrustAccessApplication#allowed_clipboard_local_to_remote_formats}
	AllowedClipboardLocalToRemoteFormats *[]*string `field:"optional" json:"allowedClipboardLocalToRemoteFormats" yaml:"allowedClipboardLocalToRemoteFormats"`
	// Clipboard formats allowed when copying from remote RDP session to local machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_application#allowed_clipboard_remote_to_local_formats ZeroTrustAccessApplication#allowed_clipboard_remote_to_local_formats}
	AllowedClipboardRemoteToLocalFormats *[]*string `field:"optional" json:"allowedClipboardRemoteToLocalFormats" yaml:"allowedClipboardRemoteToLocalFormats"`
}

