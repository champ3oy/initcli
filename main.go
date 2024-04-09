package main

import (
	"fmt"
	"initcli/frameworks"
	"os"
	"os/exec"
)

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error running command '%s': %v", command, err)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Framework argument is required.")
		os.Exit(1)
	}

	framework := os.Args[1]
	projectName := "myapp"
	var optionalArg string

	if len(os.Args) > 2 {
		projectName = os.Args[2]
	}

	if len(os.Args) > 3 {
		optionalArg = os.Args[3]
	}

	switch framework {
	case "nextjs":
		if err := frameworks.InitNextJS(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Next.js project:", err)
			os.Exit(1)
		}
	case "expo":
		if err := frameworks.InitExpo(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Expo project:", err)
			os.Exit(1)
		}
	case "react":
		if err := frameworks.InitReact(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing React project:", err)
			os.Exit(1)
		}
	case "vue":
		if err := frameworks.InitVue(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Vue.js project:", err)
			os.Exit(1)
		}
	case "angular":
		if err := frameworks.InitAngular(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Angular project:", err)
			os.Exit(1)
		}
	case "reactnative":
		if err := frameworks.InitReactNative(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing React Native project:", err)
			os.Exit(1)
		}
	case "svelte":
		if err := frameworks.InitSvelte(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Svelte project:", err)
			os.Exit(1)
		}
	case "solid":
		if err := frameworks.InitSolid(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Solid project:", err)
			os.Exit(1)
		}
	case "electron":
		if err := frameworks.InitElectron(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Electron project:", err)
			os.Exit(1)
		}
	case "nestjs":
		if err := frameworks.InitNestJS(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Nest.js project:", err)
			os.Exit(1)
		}
	case "flutter":
		if err := frameworks.InitFlutter(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Flutter project:", err)
			os.Exit(1)
		}
	case "vuepress":
		if err := frameworks.InitVuePress(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing VuePress project:", err)
			os.Exit(1)
		}
	case "gatsby":
		if err := frameworks.InitGatsby(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Gatsby project:", err)
			os.Exit(1)
		}
	case "hugo":
		if err := frameworks.InitHugo(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Hugo project:", err)
			os.Exit(1)
		}
	case "nuxtjs":
		if err := frameworks.InitNuxtJS(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Nuxt.js project:", err)
			os.Exit(1)
		}
	case "flutterweb":
		if err := frameworks.InitFlutterWeb(projectName, optionalArg); err != nil {
			fmt.Println("Error initializing Flutter web project:", err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid framework:", framework)
		os.Exit(1)
	}

}
