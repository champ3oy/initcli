package frameworks

import command "initcli/util"

func InitFlutter(projectName string, optionalArg string) error {
	return command.RunCommand("flutter", "create", projectName, optionalArg)
}
