package frameworks

import command "initcli/util"

func InitFlutterWeb(projectName string, optionalArg string) error {
	return command.RunCommand("flutter", "create", "-t", "web", projectName, optionalArg)
}
