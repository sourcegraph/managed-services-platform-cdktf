package datacloudflarehostnametlssetting

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareHostnameTlsSettingConfig struct {
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
	// The TLS Setting name.
	//
	// The value type depends on the setting:
	// - `ciphers`: value is an array of cipher suite strings (e.g., `["ECDHE-RSA-AES128-GCM-SHA256", "AES128-GCM-SHA256"]`)
	// - `min_tls_version`: value is a TLS version string (`"1.0"`, `"1.1"`, `"1.2"`, or `"1.3"`)
	// - `http2`: value is `"on"` or `"off"`
	// Available values: "ciphers", "min_tls_version", "http2".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/hostname_tls_setting#setting_id DataCloudflareHostnameTlsSetting#setting_id}
	SettingId *string `field:"required" json:"settingId" yaml:"settingId"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/hostname_tls_setting#zone_id DataCloudflareHostnameTlsSetting#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

