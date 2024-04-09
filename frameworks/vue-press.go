package frameworks

import command "initcli/util"

func InitVuePress(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "create-vuepress", projectName, optionalArg)
}
