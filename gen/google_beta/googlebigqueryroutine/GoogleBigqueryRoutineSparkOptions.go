package googlebigqueryroutine


type GoogleBigqueryRoutineSparkOptions struct {
	// Archive files to be extracted into the working directory of each executor.
	//
	// For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#archive_uris GoogleBigqueryRoutine#archive_uris}
	ArchiveUris *[]*string `field:"optional" json:"archiveUris" yaml:"archiveUris"`
	// Fully qualified name of the user-provided Spark connection object. Format: "projects/{projectId}/locations/{locationId}/connections/{connectionId}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#connection GoogleBigqueryRoutine#connection}
	Connection *string `field:"optional" json:"connection" yaml:"connection"`
	// Custom container image for the runtime environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#container_image GoogleBigqueryRoutine#container_image}
	ContainerImage *string `field:"optional" json:"containerImage" yaml:"containerImage"`
	// Files to be placed in the working directory of each executor.
	//
	// For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#file_uris GoogleBigqueryRoutine#file_uris}
	FileUris *[]*string `field:"optional" json:"fileUris" yaml:"fileUris"`
	// JARs to include on the driver and executor CLASSPATH. For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#jar_uris GoogleBigqueryRoutine#jar_uris}
	JarUris *[]*string `field:"optional" json:"jarUris" yaml:"jarUris"`
	// The fully qualified name of a class in jarUris, for example, com.example.wordcount. Exactly one of mainClass and main_jar_uri field should be set for Java/Scala language type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#main_class GoogleBigqueryRoutine#main_class}
	MainClass *string `field:"optional" json:"mainClass" yaml:"mainClass"`
	// The main file/jar URI of the Spark application.
	//
	// Exactly one of the definitionBody field and the mainFileUri field must be set for Python.
	// Exactly one of mainClass and mainFileUri field should be set for Java/Scala language type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#main_file_uri GoogleBigqueryRoutine#main_file_uri}
	MainFileUri *string `field:"optional" json:"mainFileUri" yaml:"mainFileUri"`
	// Configuration properties as a set of key/value pairs, which will be passed on to the Spark application.
	//
	// For more information, see Apache Spark and the procedure option list.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#properties GoogleBigqueryRoutine#properties}
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Python files to be placed on the PYTHONPATH for PySpark application.
	//
	// Supported file types: .py, .egg, and .zip. For more information about Apache Spark, see Apache Spark.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#py_file_uris GoogleBigqueryRoutine#py_file_uris}
	PyFileUris *[]*string `field:"optional" json:"pyFileUris" yaml:"pyFileUris"`
	// Runtime version. If not specified, the default runtime version is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_bigquery_routine#runtime_version GoogleBigqueryRoutine#runtime_version}
	RuntimeVersion *string `field:"optional" json:"runtimeVersion" yaml:"runtimeVersion"`
}

