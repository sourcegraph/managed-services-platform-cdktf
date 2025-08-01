package googlememcacheinstance


type GoogleMemcacheInstanceNodeConfig struct {
	// Number of CPUs per node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_memcache_instance#cpu_count GoogleMemcacheInstance#cpu_count}
	CpuCount *float64 `field:"required" json:"cpuCount" yaml:"cpuCount"`
	// Memory size in Mebibytes for each memcache node.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.45.0/docs/resources/google_memcache_instance#memory_size_mb GoogleMemcacheInstance#memory_size_mb}
	MemorySizeMb *float64 `field:"required" json:"memorySizeMb" yaml:"memorySizeMb"`
}

