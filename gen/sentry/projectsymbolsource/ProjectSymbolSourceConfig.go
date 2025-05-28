package projectsymbolsource

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectSymbolSourceConfig struct {
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
	// The human-readable name of the source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#name ProjectSymbolSource#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The organization of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#organization ProjectSymbolSource#organization}
	Organization *string `field:"required" json:"organization" yaml:"organization"`
	// The project of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#project ProjectSymbolSource#project}
	Project *string `field:"required" json:"project" yaml:"project"`
	// The type of symbol source.
	//
	// One of `appStoreConnect` (App Store Connect), `http` (SymbolServer (HTTP)), `gcs` (Google Cloud Storage), `s3` (Amazon S3).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#type ProjectSymbolSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The AWS Access Key.Required for S3 sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#access_key ProjectSymbolSource#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// The App Store Connect Issuer ID. Required for AppStoreConnect sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#app_connect_issuer ProjectSymbolSource#app_connect_issuer}
	AppConnectIssuer *string `field:"optional" json:"appConnectIssuer" yaml:"appConnectIssuer"`
	// The App Store Connect API Private Key. Required for AppStoreConnect sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#app_connect_private_key ProjectSymbolSource#app_connect_private_key}
	AppConnectPrivateKey *string `field:"optional" json:"appConnectPrivateKey" yaml:"appConnectPrivateKey"`
	// The App Store Connect App ID. Required for AppStoreConnect sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#app_id ProjectSymbolSource#app_id}
	AppId *string `field:"optional" json:"appId" yaml:"appId"`
	// The GCS or S3 bucket where the source resides.
	//
	// Required for GCS and S3 sourcse, invalid for HTTP and AppStoreConnect sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#bucket ProjectSymbolSource#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The GCS email address for authentication. Required for GCS sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#client_email ProjectSymbolSource#client_email}
	ClientEmail *string `field:"optional" json:"clientEmail" yaml:"clientEmail"`
	// Layout settings for the source. This is required for HTTP, GCS, and S3 sources and invalid for AppStoreConnect sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#layout ProjectSymbolSource#layout}
	Layout *ProjectSymbolSourceLayout `field:"optional" json:"layout" yaml:"layout"`
	// The password for accessing the source. Optional for HTTP sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#password ProjectSymbolSource#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The GCS or S3 prefix. Optional for GCS and S3 sourcse, invalid for HTTP and AppStoreConnect sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#prefix ProjectSymbolSource#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The GCS private key. Required for GCS sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#private_key ProjectSymbolSource#private_key}
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// The source's S3 region. Required for S3 sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#region ProjectSymbolSource#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The AWS Secret Access Key.Required for S3 sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#secret_key ProjectSymbolSource#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// The source's URL. Optional for HTTP sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#url ProjectSymbolSource#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// The user name for accessing the source. Optional for HTTP sources, invalid for all others.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.5/docs/resources/project_symbol_source#username ProjectSymbolSource#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

