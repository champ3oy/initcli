package frameworks

import command "initcli/util"

func InitSolid(projectName string, optionalArg string) error {
	return command.RunCommand("npm", "init", "solid@latest", projectName, optionalArg)
}
