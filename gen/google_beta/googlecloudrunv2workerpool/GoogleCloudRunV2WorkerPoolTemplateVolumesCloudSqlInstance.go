package googlecloudrunv2workerpool


type GoogleCloudRunV2WorkerPoolTemplateVolumesCloudSqlInstance struct {
	// The Cloud SQL instance connection names, as can be found in https://console.cloud.google.com/sql/instances. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run. Format: {project}:{location}:{instance}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_cloud_run_v2_worker_pool#instances GoogleCloudRunV2WorkerPool#instances}
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
}

