package projectsymbolsource


type ProjectSymbolSourceLayout struct {
	// The casing of the symbol source layout.
	//
	// The layout of the folder structure. The options are: `default` - Default (mixed case), `uppercase` - Uppercase, `lowercase` - Lowercase.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/project_symbol_source#casing ProjectSymbolSource#casing}
	Casing *string `field:"required" json:"casing" yaml:"casing"`
	// The layout of the folder structure.
	//
	// The options are: `native` - Platform-Specific (SymStore / GDB / LLVM), `symstore` - Microsoft SymStore, `symstore_index2` - Microsoft SymStore (with index2.txt), `ssqp` - Microsoft SSQP, `unified` - Unified Symbol Server Layout, `debuginfod` - debuginfod.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/jianyuan/sentry/0.14.3/docs/resources/project_symbol_source#type ProjectSymbolSource#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

