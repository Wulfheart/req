package requester

const (
	Newline        = "\r\n"
	WindowsNewline = "\r\n"
	UnixNewline    = "\n"

	RequestSeparator = "###"

	EnvironmentVariablesFileName = "env.req.env"
	CollectionVariablesFileName  = "collection.req.env"
	ConfigDirectoryName          = ".req"
	GlobalHeaderFileName         = "global_headers"
	PreRunFileName               = "pre_run.req"

	TemplateVariableDeclarationLeft  = "{{"
	TemplateVariableDeclarationRight = "}}"
)
