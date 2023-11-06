package googleaccesscontextmanageraccesslevels


type GoogleAccessContextManagerAccessLevelsAccessLevelsBasicConditionsDevicePolicyOsConstraints struct {
	// The operating system type of the device. Possible values: ["OS_UNSPECIFIED", "DESKTOP_MAC", "DESKTOP_WINDOWS", "DESKTOP_LINUX", "DESKTOP_CHROME_OS", "ANDROID", "IOS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_access_context_manager_access_levels#os_type GoogleAccessContextManagerAccessLevels#os_type}
	OsType *string `field:"required" json:"osType" yaml:"osType"`
	// The minimum allowed OS version.
	//
	// If not set, any version
	// of this OS satisfies the constraint.
	// Format: "major.minor.patch" such as "10.5.301", "9.2.1".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_access_context_manager_access_levels#minimum_version GoogleAccessContextManagerAccessLevels#minimum_version}
	MinimumVersion *string `field:"optional" json:"minimumVersion" yaml:"minimumVersion"`
}

