package connectivitydirectoryservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectivityDirectoryServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#host ConnectivityDirectoryService#host}.
	Host *ConnectivityDirectoryServiceHost `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#name ConnectivityDirectoryService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Available values: "tcp", "http".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#type ConnectivityDirectoryService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Account identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#account_id ConnectivityDirectoryService#account_id}
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Available values: "postgresql", "mysql".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#app_protocol ConnectivityDirectoryService#app_protocol}
	AppProtocol *string `field:"optional" json:"appProtocol" yaml:"appProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#http_port ConnectivityDirectoryService#http_port}.
	HttpPort *float64 `field:"optional" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#https_port ConnectivityDirectoryService#https_port}.
	HttpsPort *float64 `field:"optional" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#tcp_port ConnectivityDirectoryService#tcp_port}.
	TcpPort *float64 `field:"optional" json:"tcpPort" yaml:"tcpPort"`
	// TLS settings for a connectivity service.
	//
	// If omitted, the default mode (`verify_full`) is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/connectivity_directory_service#tls_settings ConnectivityDirectoryService#tls_settings}
	TlsSettings *ConnectivityDirectoryServiceTlsSettings `field:"optional" json:"tlsSettings" yaml:"tlsSettings"`
}

