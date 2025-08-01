package endpointsserviceconsumersiambinding


type EndpointsServiceConsumersIamBindingCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/endpoints_service_consumers_iam_binding#expression EndpointsServiceConsumersIamBinding#expression}.
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/endpoints_service_consumers_iam_binding#title EndpointsServiceConsumersIamBinding#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/endpoints_service_consumers_iam_binding#description EndpointsServiceConsumersIamBinding#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

