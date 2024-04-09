package frameworks

import command "initcli/util"

func InitSvelte(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "degit", "sveltejs/template", projectName, optionalArg)
}
