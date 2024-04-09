package frameworks

import command "initcli/util"

func InitVue(projectName string, optionalArg string) error {
	return command.RunCommand("vue", "create", projectName, optionalArg)
}
