package googledataplexasset


type GoogleDataplexAssetDiscoverySpec struct {
	// Required. Whether discovery is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataplex_asset#enabled GoogleDataplexAsset#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// csv_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataplex_asset#csv_options GoogleDataplexAsset#csv_options}
	CsvOptions *GoogleDataplexAssetDiscoverySpecCsvOptions `field:"optional" json:"csvOptions" yaml:"csvOptions"`
	// Optional.
	//
	// The list of patterns to apply for selecting data to exclude during discovery. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataplex_asset#exclude_patterns GoogleDataplexAsset#exclude_patterns}
	ExcludePatterns *[]*string `field:"optional" json:"excludePatterns" yaml:"excludePatterns"`
	// Optional.
	//
	// The list of patterns to apply for selecting data to include during discovery if only a subset of the data should considered. For Cloud Storage bucket assets, these are interpreted as glob patterns used to match object names. For BigQuery dataset assets, these are interpreted as patterns to match table names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataplex_asset#include_patterns GoogleDataplexAsset#include_patterns}
	IncludePatterns *[]*string `field:"optional" json:"includePatterns" yaml:"includePatterns"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataplex_asset#json_options GoogleDataplexAsset#json_options}
	JsonOptions *GoogleDataplexAssetDiscoverySpecJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
	// Optional.
	//
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) for running discovery periodically. Successive discovery runs must be scheduled at least 60 minutes apart. The default value is to run discovery every 60 minutes. To explicitly set a timezone to the cron tab, apply a prefix in the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or TZ=${IANA_TIME_ZONE}". The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone database. For example, "CRON_TZ=America/New_York 1 * * * *", or "TZ=America/New_York 1 * * * *".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.78.0/docs/resources/google_dataplex_asset#schedule GoogleDataplexAsset#schedule}
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
}

