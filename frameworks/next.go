package frameworks

import command "initcli/util"

func InitNextJS(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "create-next-app", projectName, optionalArg)
}
