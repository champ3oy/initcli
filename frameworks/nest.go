package frameworks

import command "initcli/util"

func InitNestJS(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "create-nestjs-app", projectName, optionalArg)
}
