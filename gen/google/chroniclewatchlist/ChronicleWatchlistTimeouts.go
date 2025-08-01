package chroniclewatchlist


type ChronicleWatchlistTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#create ChronicleWatchlist#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#delete ChronicleWatchlist#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.45.0/docs/resources/chronicle_watchlist#update ChronicleWatchlist#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

