package googledataplextaskiambinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleDataplexTaskIamBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#lake GoogleDataplexTaskIamBinding#lake}.
	Lake *string `field:"required" json:"lake" yaml:"lake"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#members GoogleDataplexTaskIamBinding#members}.
	Members *[]*string `field:"required" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#role GoogleDataplexTaskIamBinding#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#task_id GoogleDataplexTaskIamBinding#task_id}.
	TaskId *string `field:"required" json:"taskId" yaml:"taskId"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#condition GoogleDataplexTaskIamBinding#condition}
	Condition *GoogleDataplexTaskIamBindingCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#id GoogleDataplexTaskIamBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#location GoogleDataplexTaskIamBinding#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.37.0/docs/resources/google_dataplex_task_iam_binding#project GoogleDataplexTaskIamBinding#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
}

