package memorystoreinstance


type MemorystoreInstancePersistenceConfigAofConfig struct {
	// Optional. The fsync mode.   Possible values:  NEVER EVERY_SEC ALWAYS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/memorystore_instance#append_fsync MemorystoreInstance#append_fsync}
	AppendFsync *string `field:"optional" json:"appendFsync" yaml:"appendFsync"`
}

