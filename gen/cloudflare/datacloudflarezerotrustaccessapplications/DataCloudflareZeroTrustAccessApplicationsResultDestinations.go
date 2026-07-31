package datacloudflarezerotrustaccessapplications


type DataCloudflareZeroTrustAccessApplicationsResultDestinations struct {
	// A MCP server id configured in ai-controls. Access will secure the MCP server if accessed through a MCP portal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/data-sources/zero_trust_access_applications#mcp_server_id DataCloudflareZeroTrustAccessApplications#mcp_server_id}
	McpServerId *string `field:"optional" json:"mcpServerId" yaml:"mcpServerId"`
}

