package zerotrustdevicesettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustDeviceSettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#account_id ZeroTrustDeviceSettings#account_id}.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Sets the time limit, in seconds, that a user can use an override code to bypass WARP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#disable_for_time ZeroTrustDeviceSettings#disable_for_time}
	DisableForTime *float64 `field:"optional" json:"disableForTime" yaml:"disableForTime"`
	// Controls whether the external emergency disconnect feature is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#external_emergency_signal_enabled ZeroTrustDeviceSettings#external_emergency_signal_enabled}
	ExternalEmergencySignalEnabled interface{} `field:"optional" json:"externalEmergencySignalEnabled" yaml:"externalEmergencySignalEnabled"`
	// The SHA256 fingerprint (64 hexadecimal characters) of the HTTPS server certificate for the external_emergency_signal_url.
	//
	// If provided, the WARP client will use this value to verify the server's identity. The device will ignore any response if the server's certificate fingerprint does not exactly match this value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#external_emergency_signal_fingerprint ZeroTrustDeviceSettings#external_emergency_signal_fingerprint}
	ExternalEmergencySignalFingerprint *string `field:"optional" json:"externalEmergencySignalFingerprint" yaml:"externalEmergencySignalFingerprint"`
	// The interval at which the WARP client fetches the emergency disconnect signal, formatted as a duration string (e.g., "5m", "2m30s", "1h"). Minimum 30 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#external_emergency_signal_interval ZeroTrustDeviceSettings#external_emergency_signal_interval}
	ExternalEmergencySignalInterval *string `field:"optional" json:"externalEmergencySignalInterval" yaml:"externalEmergencySignalInterval"`
	// The HTTPS URL from which to fetch the emergency disconnect signal.
	//
	// Must use HTTPS and have an IPv4 or IPv6 address as the host.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#external_emergency_signal_url ZeroTrustDeviceSettings#external_emergency_signal_url}
	ExternalEmergencySignalUrl *string `field:"optional" json:"externalEmergencySignalUrl" yaml:"externalEmergencySignalUrl"`
	// Enable gateway proxy filtering on TCP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#gateway_proxy_enabled ZeroTrustDeviceSettings#gateway_proxy_enabled}
	GatewayProxyEnabled interface{} `field:"optional" json:"gatewayProxyEnabled" yaml:"gatewayProxyEnabled"`
	// Enable gateway proxy filtering on UDP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#gateway_udp_proxy_enabled ZeroTrustDeviceSettings#gateway_udp_proxy_enabled}
	GatewayUdpProxyEnabled interface{} `field:"optional" json:"gatewayUdpProxyEnabled" yaml:"gatewayUdpProxyEnabled"`
	// Enable installation of cloudflare managed root certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#root_certificate_installation_enabled ZeroTrustDeviceSettings#root_certificate_installation_enabled}
	RootCertificateInstallationEnabled interface{} `field:"optional" json:"rootCertificateInstallationEnabled" yaml:"rootCertificateInstallationEnabled"`
	// Enable using CGNAT virtual IPv4.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/zero_trust_device_settings#use_zt_virtual_ip ZeroTrustDeviceSettings#use_zt_virtual_ip}
	UseZtVirtualIp interface{} `field:"optional" json:"useZtVirtualIp" yaml:"useZtVirtualIp"`
}

