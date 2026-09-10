package zerotrustgatewaysettings


type ZeroTrustGatewaySettingsSettingsBlockPage struct {
	// Specify the block page background color in `#rrggbb` format when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#background_color ZeroTrustGatewaySettings#background_color}
	BackgroundColor *string `field:"optional" json:"backgroundColor" yaml:"backgroundColor"`
	// Specify whether to enable the custom block page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#enabled ZeroTrustGatewaySettings#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Specify the block page footer text when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#footer_text ZeroTrustGatewaySettings#footer_text}
	FooterText *string `field:"optional" json:"footerText" yaml:"footerText"`
	// Specify the block page header text when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#header_text ZeroTrustGatewaySettings#header_text}
	HeaderText *string `field:"optional" json:"headerText" yaml:"headerText"`
	// Specify whether to append context to target_uri as query parameters. This applies only when the mode is redirect_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#include_context ZeroTrustGatewaySettings#include_context}
	IncludeContext interface{} `field:"optional" json:"includeContext" yaml:"includeContext"`
	// Specify the full URL to the logo file when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#logo_path ZeroTrustGatewaySettings#logo_path}
	LogoPath *string `field:"optional" json:"logoPath" yaml:"logoPath"`
	// Specify the admin email for users to contact when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#mailto_address ZeroTrustGatewaySettings#mailto_address}
	MailtoAddress *string `field:"optional" json:"mailtoAddress" yaml:"mailtoAddress"`
	// Specify the subject line for emails created from the block page when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#mailto_subject ZeroTrustGatewaySettings#mailto_subject}
	MailtoSubject *string `field:"optional" json:"mailtoSubject" yaml:"mailtoSubject"`
	// Specify whether to redirect users to a Cloudflare-hosted block page or a customer-provided URI. Available values: "", "customized_block_page", "redirect_uri".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#mode ZeroTrustGatewaySettings#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Specify the block page title when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#name ZeroTrustGatewaySettings#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Indicate that this setting was shared via the Orgs API and read only for the current account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#read_only ZeroTrustGatewaySettings#read_only}
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Indicate the account tag of the account that shared this setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#source_account ZeroTrustGatewaySettings#source_account}
	SourceAccount *string `field:"optional" json:"sourceAccount" yaml:"sourceAccount"`
	// Specify whether to suppress detailed information at the bottom of the block page when the mode is customized_block_page.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#suppress_footer ZeroTrustGatewaySettings#suppress_footer}
	SuppressFooter interface{} `field:"optional" json:"suppressFooter" yaml:"suppressFooter"`
	// Specify the URI to redirect users to when the mode is redirect_uri.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#target_uri ZeroTrustGatewaySettings#target_uri}
	TargetUri *string `field:"optional" json:"targetUri" yaml:"targetUri"`
	// Indicate the version number of the setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_gateway_settings#version ZeroTrustGatewaySettings#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

