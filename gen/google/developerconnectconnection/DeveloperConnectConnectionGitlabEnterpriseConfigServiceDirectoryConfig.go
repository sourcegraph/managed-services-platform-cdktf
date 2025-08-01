package developerconnectconnection


type DeveloperConnectConnectionGitlabEnterpriseConfigServiceDirectoryConfig struct {
	// Required. The Service Directory service name. Format: projects/{project}/locations/{location}/namespaces/{namespace}/services/{service}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/developer_connect_connection#service DeveloperConnectConnection#service}
	Service *string `field:"required" json:"service" yaml:"service"`
}

