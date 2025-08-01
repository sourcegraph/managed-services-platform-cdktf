package googledatacatalogentry


type GoogleDataCatalogEntryGcsFilesetSpec struct {
	// Patterns to identify a set of files in Google Cloud Storage.
	//
	// See [Cloud Storage documentation](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)
	// for more information. Note that bucket wildcards are currently not supported. Examples of valid filePatterns:
	//
	// * gs://bucket_name/dir/*: matches all files within bucket_name/dir directory.
	// * gs://bucket_name/dir/**: matches all files in bucket_name/dir spanning all subdirectories.
	// * gs://bucket_name/file*: matches files prefixed by file in bucket_name
	// * gs://bucket_name/??.txt: matches files with two characters followed by .txt in bucket_name
	// * gs://bucket_name/[aeiou].txt: matches files that contain a single vowel character followed by .txt in bucket_name
	// * gs://bucket_name/[a-m].txt: matches files that contain a, b, ... or m followed by .txt in bucket_name
	// * gs://bucket_name/a/* /b: matches all files in bucket_name that match a/* /b pattern, such as a/c/b, a/d/b
	// * gs://another_bucket/a.txt: matches gs://another_bucket/a.txt
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_data_catalog_entry#file_patterns GoogleDataCatalogEntry#file_patterns}
	//
	// Note: The above comment contained a comment block ending sequence (* followed by /). We have introduced a space between to prevent syntax errors. Please ignore the space.
	FilePatterns *[]*string `field:"required" json:"filePatterns" yaml:"filePatterns"`
}

