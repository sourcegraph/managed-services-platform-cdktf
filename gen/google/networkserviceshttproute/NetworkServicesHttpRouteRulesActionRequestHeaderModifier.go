package networkserviceshttproute


type NetworkServicesHttpRouteRulesActionRequestHeaderModifier struct {
	// Add the headers with given map where key is the name of the header, value is the value of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#add NetworkServicesHttpRoute#add}
	Add *map[string]*string `field:"optional" json:"add" yaml:"add"`
	// Remove headers (matching by header names) specified in the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#remove NetworkServicesHttpRoute#remove}
	Remove *[]*string `field:"optional" json:"remove" yaml:"remove"`
	// Completely overwrite/replace the headers with given map where key is the name of the header, value is the value of the header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/network_services_http_route#set NetworkServicesHttpRoute#set}
	Set *map[string]*string `field:"optional" json:"set" yaml:"set"`
}

