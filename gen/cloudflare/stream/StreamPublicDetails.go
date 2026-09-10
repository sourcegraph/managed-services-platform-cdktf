package stream


type StreamPublicDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/stream#channel_link Stream#channel_link}.
	ChannelLink *string `field:"optional" json:"channelLink" yaml:"channelLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/stream#logo Stream#logo}.
	Logo *string `field:"optional" json:"logo" yaml:"logo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/stream#share_link Stream#share_link}.
	ShareLink *string `field:"optional" json:"shareLink" yaml:"shareLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/stream#title Stream#title}.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

