package frameworks

type Framework struct {
	Name        string
	UseOfficial bool
	CliCommand  []string
	TemplateDir string
}

var All = []Framework{
	{Name: "express-js", UseOfficial: false, TemplateDir: "express-js"},
	{Name: "express-ts", UseOfficial: false, TemplateDir: "express-ts"},
	{Name: "fastapi", UseOfficial: false, TemplateDir: "fastapi"},
	{Name: "gin", UseOfficial: false, TemplateDir: "gin"},
	{Name: "fiber", UseOfficial: false, TemplateDir: "fiber"},
	// If we add frameworks with official CLIs:
	{Name: "hono", UseOfficial: true, CliCommand: []string{"npx", "create-hono@latest"}},
	{Name: "nextjs", UseOfficial: true, CliCommand: []string{"npx", "create-next-app@latest"}},
	{Name: "nestjs", UseOfficial: true, CliCommand: []string{"npx", "@nestjs/cli@latest"}},
}
