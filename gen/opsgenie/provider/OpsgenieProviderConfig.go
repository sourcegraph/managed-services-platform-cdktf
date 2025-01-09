package provider


type OpsgenieProviderConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs#api_key OpsgenieProvider#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs#alias OpsgenieProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs#api_retry_count OpsgenieProvider#api_retry_count}.
	ApiRetryCount *float64 `field:"optional" json:"apiRetryCount" yaml:"apiRetryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs#api_retry_wait_max OpsgenieProvider#api_retry_wait_max}.
	ApiRetryWaitMax *float64 `field:"optional" json:"apiRetryWaitMax" yaml:"apiRetryWaitMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs#api_retry_wait_min OpsgenieProvider#api_retry_wait_min}.
	ApiRetryWaitMin *float64 `field:"optional" json:"apiRetryWaitMin" yaml:"apiRetryWaitMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opsgenie/opsgenie/0.6.37/docs#api_url OpsgenieProvider#api_url}.
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
}

