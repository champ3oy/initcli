package frameworks

import command "initcli/util"

func InitExpo(projectName string, optionalArg string) error {
	return command.RunCommand("npx expo", "init", projectName, optionalArg)
}
