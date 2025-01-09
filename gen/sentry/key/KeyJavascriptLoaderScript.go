package key


type KeyJavascriptLoaderScript struct {
	// The version of the browser SDK to load.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/key#browser_sdk_version Key#browser_sdk_version}
	BrowserSdkVersion *string `field:"optional" json:"browserSdkVersion" yaml:"browserSdkVersion"`
	// Whether debug bundles & logging are enabled for this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/key#debug_enabled Key#debug_enabled}
	DebugEnabled interface{} `field:"optional" json:"debugEnabled" yaml:"debugEnabled"`
	// Whether performance monitoring is enabled for this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/key#performance_monitoring_enabled Key#performance_monitoring_enabled}
	PerformanceMonitoringEnabled interface{} `field:"optional" json:"performanceMonitoringEnabled" yaml:"performanceMonitoringEnabled"`
	// Whether session replay is enabled for this key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/key#session_replay_enabled Key#session_replay_enabled}
	SessionReplayEnabled interface{} `field:"optional" json:"sessionReplayEnabled" yaml:"sessionReplayEnabled"`
}

