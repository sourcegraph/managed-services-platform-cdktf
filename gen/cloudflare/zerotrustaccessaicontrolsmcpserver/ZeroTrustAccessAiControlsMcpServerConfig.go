package zerotrustaccessaicontrolsmcpserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ZeroTrustAccessAiControlsMcpServerConfig struct {
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
	// Available values: "oauth", "bearer", "unauthenticated".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#auth_type ZeroTrustAccessAiControlsMcpServer#auth_type}
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#hostname ZeroTrustAccessAiControlsMcpServer#hostname}.
	Hostname *string `field:"required" json:"hostname" yaml:"hostname"`
	// server id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#id ZeroTrustAccessAiControlsMcpServer#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#name ZeroTrustAccessAiControlsMcpServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#account_id ZeroTrustAccessAiControlsMcpServer#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#auth_credentials ZeroTrustAccessAiControlsMcpServer#auth_credentials}.
	AuthCredentials *string `field:"optional" json:"authCredentials" yaml:"authCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#description ZeroTrustAccessAiControlsMcpServer#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#updated_prompts ZeroTrustAccessAiControlsMcpServer#updated_prompts}.
	UpdatedPrompts interface{} `field:"optional" json:"updatedPrompts" yaml:"updatedPrompts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.19.0/docs/resources/zero_trust_access_ai_controls_mcp_server#updated_tools ZeroTrustAccessAiControlsMcpServer#updated_tools}.
	UpdatedTools interface{} `field:"optional" json:"updatedTools" yaml:"updatedTools"`
}

