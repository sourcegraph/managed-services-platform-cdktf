package beyondcorpsecuritygatewayiammember

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BeyondcorpSecurityGatewayIamMemberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_iam_member#member BeyondcorpSecurityGatewayIamMember#member}.
	Member *string `field:"required" json:"member" yaml:"member"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_iam_member#role BeyondcorpSecurityGatewayIamMember#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_iam_member#security_gateway_id BeyondcorpSecurityGatewayIamMember#security_gateway_id}.
	SecurityGatewayId *string `field:"required" json:"securityGatewayId" yaml:"securityGatewayId"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_iam_member#condition BeyondcorpSecurityGatewayIamMember#condition}
	Condition *BeyondcorpSecurityGatewayIamMemberCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_iam_member#id BeyondcorpSecurityGatewayIamMember#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_iam_member#location BeyondcorpSecurityGatewayIamMember#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/beyondcorp_security_gateway_iam_member#project BeyondcorpSecurityGatewayIamMember#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

