package frameworks

import command "initcli/util"

func InitReact(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "create-react-app", projectName, optionalArg)
}
