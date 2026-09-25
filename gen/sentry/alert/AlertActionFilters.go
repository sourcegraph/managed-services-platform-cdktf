package alert


type AlertActionFilters struct {
	// The actions to perform.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#actions Alert#actions}
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// The logic to apply to the conditions.
	//
	// `any` will evaluate all conditions, and return true if any of those are met. `any-short` will stop evaluating conditions as soon as one is met. `all` will evaluate all conditions, and return true if all of those are met. `none` will return true if none of the conditions are met, will return false immediately if any are met. Valid values are: `any`, `any-short`, `all`, and `none`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#logic_type Alert#logic_type}
	LogicType *string `field:"required" json:"logicType" yaml:"logicType"`
	// The conditions to evaluate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.15.7/docs/resources/alert#conditions Alert#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

