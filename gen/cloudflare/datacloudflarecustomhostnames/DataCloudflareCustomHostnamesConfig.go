package datacloudflarecustomhostnames

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataCloudflareCustomHostnamesConfig struct {
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
	// Filter by the certificate authority that issued the SSL certificate. Available values: "google", "lets_encrypt", "ssl_com".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#certificate_authority DataCloudflareCustomHostnames#certificate_authority}
	CertificateAuthority *string `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Filter by custom origin server name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#custom_origin_server DataCloudflareCustomHostnames#custom_origin_server}
	CustomOriginServer *string `field:"optional" json:"customOriginServer" yaml:"customOriginServer"`
	// Direction to order hostnames. Available values: "asc", "desc".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#direction DataCloudflareCustomHostnames#direction}
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#hostname DataCloudflareCustomHostnames#hostname}.
	Hostname *DataCloudflareCustomHostnamesHostname `field:"optional" json:"hostname" yaml:"hostname"`
	// Filter by the hostname's activation status.
	//
	// Available values: "active", "pending", "active_redeploying", "moved", "pending_deletion", "deleted", "pending_blocked", "pending_migration", "pending_provisioned", "test_pending", "test_active", "test_active_apex", "test_blocked", "test_failed", "provisioned", "blocked".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#hostname_status DataCloudflareCustomHostnames#hostname_status}
	HostnameStatus *string `field:"optional" json:"hostnameStatus" yaml:"hostnameStatus"`
	// Hostname ID to match against.
	//
	// This ID was generated and returned during the initial custom_hostname creation. This parameter cannot be used with the 'hostname' parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#id DataCloudflareCustomHostnames#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Max items to fetch, default: 1000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#max_items DataCloudflareCustomHostnames#max_items}
	MaxItems *float64 `field:"optional" json:"maxItems" yaml:"maxItems"`
	// Field to order hostnames by. Available values: "ssl", "ssl_status".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#order DataCloudflareCustomHostnames#order}
	Order *string `field:"optional" json:"order" yaml:"order"`
	// Whether to filter hostnames based on if they have SSL enabled. Available values: 0, 1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#ssl DataCloudflareCustomHostnames#ssl}
	Ssl *float64 `field:"optional" json:"ssl" yaml:"ssl"`
	// Filter by SSL certificate status.
	//
	// Available values: "initializing", "pending_validation", "deleted", "pending_issuance", "pending_deployment", "pending_deletion", "pending_expiration", "expired", "active", "initializing_timed_out", "validation_timed_out", "issuance_timed_out", "deployment_timed_out", "deletion_timed_out", "pending_cleanup", "staging_deployment", "staging_active", "deactivating", "inactive", "backup_issued", "holding_deployment".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#ssl_status DataCloudflareCustomHostnames#ssl_status}
	SslStatus *string `field:"optional" json:"sslStatus" yaml:"sslStatus"`
	// Filter by whether the custom hostname is a wildcard hostname.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#wildcard DataCloudflareCustomHostnames#wildcard}
	Wildcard interface{} `field:"optional" json:"wildcard" yaml:"wildcard"`
	// Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/data-sources/custom_hostnames#zone_id DataCloudflareCustomHostnames#zone_id}
	ZoneId *string `field:"optional" json:"zoneId" yaml:"zoneId"`
}

