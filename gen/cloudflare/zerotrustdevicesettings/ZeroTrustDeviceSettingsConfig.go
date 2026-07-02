package zerotrustdevicesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDeviceSettingsConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_settings#account_id ZeroTrustDeviceSettings#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Sets the time limit, in seconds, that a user can use an override code to bypass WARP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_settings#disable_for_time ZeroTrustDeviceSettings#disable_for_time}
	DisableForTime *float64 `field:"optional" json:"disableForTime" yaml:"disableForTime"`
	// Enable gateway proxy filtering on TCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_settings#gateway_proxy_enabled ZeroTrustDeviceSettings#gateway_proxy_enabled}
	GatewayProxyEnabled any `field:"optional" json:"gatewayProxyEnabled" yaml:"gatewayProxyEnabled"`
	// Enable gateway proxy filtering on UDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_settings#gateway_udp_proxy_enabled ZeroTrustDeviceSettings#gateway_udp_proxy_enabled}
	GatewayUdpProxyEnabled any `field:"optional" json:"gatewayUdpProxyEnabled" yaml:"gatewayUdpProxyEnabled"`
	// Enable installation of cloudflare managed root certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_settings#root_certificate_installation_enabled ZeroTrustDeviceSettings#root_certificate_installation_enabled}
	RootCertificateInstallationEnabled any `field:"optional" json:"rootCertificateInstallationEnabled" yaml:"rootCertificateInstallationEnabled"`
	// Enable using CGNAT virtual IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.7.1/docs/resources/zero_trust_device_settings#use_zt_virtual_ip ZeroTrustDeviceSettings#use_zt_virtual_ip}
	UseZtVirtualIp any `field:"optional" json:"useZtVirtualIp" yaml:"useZtVirtualIp"`
}
