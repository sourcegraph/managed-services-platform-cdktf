package googledeveloperconnectconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDeveloperConnectConnectionConfig struct {
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
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and connection_id from the method_signature of Create RPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#connection_id GoogleDeveloperConnectConnection#connection_id}
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#location GoogleDeveloperConnectConnection#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// Optional. Allows clients to store small amounts of arbitrary data.
	//
	// **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
	// Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#annotations GoogleDeveloperConnectConnection#annotations}
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// bitbucket_cloud_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#bitbucket_cloud_config GoogleDeveloperConnectConnection#bitbucket_cloud_config}
	BitbucketCloudConfig *GoogleDeveloperConnectConnectionBitbucketCloudConfig `field:"optional" json:"bitbucketCloudConfig" yaml:"bitbucketCloudConfig"`
	// bitbucket_data_center_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#bitbucket_data_center_config GoogleDeveloperConnectConnection#bitbucket_data_center_config}
	BitbucketDataCenterConfig *GoogleDeveloperConnectConnectionBitbucketDataCenterConfig `field:"optional" json:"bitbucketDataCenterConfig" yaml:"bitbucketDataCenterConfig"`
	// crypto_key_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#crypto_key_config GoogleDeveloperConnectConnection#crypto_key_config}
	CryptoKeyConfig *GoogleDeveloperConnectConnectionCryptoKeyConfig `field:"optional" json:"cryptoKeyConfig" yaml:"cryptoKeyConfig"`
	// Optional.
	//
	// If disabled is set to true, functionality is disabled for this connection.
	// Repository based API methods and webhooks processing for repositories in
	// this connection will be disabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#disabled GoogleDeveloperConnectConnection#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Optional.
	//
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#etag GoogleDeveloperConnectConnection#etag}
	Etag *string `field:"optional" json:"etag" yaml:"etag"`
	// github_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#github_config GoogleDeveloperConnectConnection#github_config}
	GithubConfig *GoogleDeveloperConnectConnectionGithubConfig `field:"optional" json:"githubConfig" yaml:"githubConfig"`
	// github_enterprise_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#github_enterprise_config GoogleDeveloperConnectConnection#github_enterprise_config}
	GithubEnterpriseConfig *GoogleDeveloperConnectConnectionGithubEnterpriseConfig `field:"optional" json:"githubEnterpriseConfig" yaml:"githubEnterpriseConfig"`
	// gitlab_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#gitlab_config GoogleDeveloperConnectConnection#gitlab_config}
	GitlabConfig *GoogleDeveloperConnectConnectionGitlabConfig `field:"optional" json:"gitlabConfig" yaml:"gitlabConfig"`
	// gitlab_enterprise_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#gitlab_enterprise_config GoogleDeveloperConnectConnection#gitlab_enterprise_config}
	GitlabEnterpriseConfig *GoogleDeveloperConnectConnectionGitlabEnterpriseConfig `field:"optional" json:"gitlabEnterpriseConfig" yaml:"gitlabEnterpriseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#id GoogleDeveloperConnectConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional. Labels as key value pairs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#labels GoogleDeveloperConnectConnection#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#project GoogleDeveloperConnectConnection#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_developer_connect_connection#timeouts GoogleDeveloperConnectConnection#timeouts}
	Timeouts *GoogleDeveloperConnectConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

