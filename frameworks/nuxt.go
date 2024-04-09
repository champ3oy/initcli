package frameworks

import command "initcli/util"

func InitNuxtJS(projectName string, optionalArg string) error {
	return command.RunCommand("npx", "create-nuxt-app", projectName, optionalArg)
}
