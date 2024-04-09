# InitCLI

InitCLI is a command-line tool written in Go that allows you to easily initialize projects for various web and app frameworks.

## Development Usage

```bash
go run main.go <framework> <project-name> [optional-args...]
```

### Frameworks and Arguments

| Framework     | Description                       | Arguments                          |
| ------------- | --------------------------------- | ---------------------------------- |
| `nextjs`      | Initialize a Next.js project      | `project-name`, `optional-args...` |
| `expo`        | Initialize an Expo project        | `project-name`, `optional-args...` |
| `react`       | Initialize a React project        | `project-name`, `optional-args...` |
| `vue`         | Initialize a Vue.js project       | `project-name`, `optional-args...` |
| `angular`     | Initialize an Angular project     | `project-name`, `optional-args...` |
| `reactnative` | Initialize a React Native project | `project-name`, `optional-args...` |
| `svelte`      | Initialize a Svelte project       | `project-name`, `optional-args...` |
| `solid`       | Initialize a Solid project        | `project-name`, `optional-args...` |
| `electron`    | Initialize an Electron project    | `project-name`, `optional-args...` |
| `nestjs`      | Initialize a Nest.js project      | `project-name`, `optional-args...` |
| `flutter`     | Initialize a Flutter project      | `project-name`, `optional-args...` |
| `vuepress`    | Initialize a VuePress project     | `project-name`, `optional-args...` |
| `gatsby`      | Initialize a Gatsby project       | `project-name`, `optional-args...` |
| `hugo`        | Initialize a Hugo project         | `project-name`, `optional-args...` |
| `nuxtjs`      | Initialize a Nuxt.js project      | `project-name`, `optional-args...` |
| `flutterweb`  | Initialize a Flutter web project  | `project-name`, `optional-args...` |

## Example

Initialize a Next.js project named "my-next-app" with TypeScript:

```bash
initcli nextjs my-next-app --typescript
```

Initialize a React Native project named "my-react-native-app" with Expo:

```bash
initcli reactnative my-react-native-app --template blank
```

## Contributions

Contributions are welcome! If you'd like to add support for initializing projects for a new framework, simply follow these steps:

1. Navigate to the `frameworks` folder in the repository.
2. Create a new Go file named after the framework you want to add support for (e.g., `myframework.go`).
3. Add the initialization function for the new framework in the Go file. Here's an example of what it might look like:

```go
package frameworks

import "github.com/your/command-package"

func InitMyFramework(projectName string, optionalArg string) error {
    return command.RunCommand("mycommand", "myarguments", projectName, optionalArg)
}
```

[by cirlormx](https://x.com/cirlormx)
