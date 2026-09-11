package hostnametlssetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HostnameTlsSettingConfig struct {
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
	// The hostname for which the tls settings are set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hostname_tls_setting#hostname HostnameTlsSetting#hostname}
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// The TLS Setting name.
	//
	// The value type depends on the setting:
	// - `ciphers`: value is an array of cipher suite strings (e.g., `["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]`)
	// - `min_tls_version`: value is a TLS version string (`"1.0"`, `"1.1"`, `"1.2"`, or `"1.3"`)
	// - `http2`: value is `"on"` or `"off"`
	// Available values: "ciphers", "min_tls_version", "http2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hostname_tls_setting#setting_id HostnameTlsSetting#setting_id}
	SettingId *string `field:"required" json:"settingId" yaml:"settingId"`
	// The TLS setting value.
	//
	// The type depends on the `setting_id` used in the request path:
	// - `ciphers`: an array of allowed cipher suite strings in BoringSSL format (e.g., `["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]`)
	// - `min_tls_version`: a string indicating the minimum TLS version — one of `"1.0"`, `"1.1"`, `"1.2"`, or `"1.3"` (e.g., `"1.2"`)
	// - `http2`: a string indicating whether HTTP/2 is enabled — `"on"` or `"off"` (e.g., `"on"`)
	// Available values: "1.0", "1.1", "1.2", "1.3", "on", "off".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hostname_tls_setting#value HostnameTlsSetting#value}
	Value *map[string]interface{} `field:"required" json:"value" yaml:"value"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/hostname_tls_setting#zone_id HostnameTlsSetting#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

