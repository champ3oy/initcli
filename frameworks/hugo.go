package frameworks

import command "initcli/util"

func InitHugo(projectName string, optionalArg string) error {
	return command.RunCommand("hugo", "new", "site", projectName, optionalArg)
}
