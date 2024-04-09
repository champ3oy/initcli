package frameworks

import command "initcli/util"

func InitAngular(projectName string, optionalArg string) error {
	return command.RunCommand("ng", "new", projectName, optionalArg)
}
