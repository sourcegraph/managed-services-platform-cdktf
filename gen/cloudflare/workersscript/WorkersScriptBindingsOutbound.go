package workersscript


type WorkersScriptBindingsOutbound struct {
	// Pass information from the Dispatch Worker to the Outbound Worker through the parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#params WorkersScript#params}
	Params *[]*string `field:"optional" json:"params" yaml:"params"`
	// Outbound worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/5.17.0/docs/resources/workers_script#worker WorkersScript#worker}
	Worker *WorkersScriptBindingsOutboundWorker `field:"optional" json:"worker" yaml:"worker"`
}

