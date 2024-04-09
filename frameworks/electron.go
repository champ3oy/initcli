package frameworks

import command "initcli/util"

func InitElectron(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "create-electron-app", projectName, optionalArg)
}
