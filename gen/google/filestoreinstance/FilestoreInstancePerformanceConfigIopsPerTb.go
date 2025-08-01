package filestoreinstance


type FilestoreInstancePerformanceConfigIopsPerTb struct {
	// The instance max IOPS will be calculated by multiplying the capacity of the instance (TB) by max_iops_per_tb, and rounding to the nearest 1000.
	//
	// The instance max IOPS
	// will be changed dynamically based on the instance
	// capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/filestore_instance#max_iops_per_tb FilestoreInstance#max_iops_per_tb}
	MaxIopsPerTb *float64 `field:"optional" json:"maxIopsPerTb" yaml:"maxIopsPerTb"`
}

