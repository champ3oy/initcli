package frameworks

import command "initcli/util"

func InitGatsby(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "gatsby", "new", projectName, optionalArg)
}
