package frameworks

import command "initcli/util"

func InitReactNative(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "react-native", "init", projectName, optionalArg)
}
