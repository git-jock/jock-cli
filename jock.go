package main

import (
	"fmt"
	"github.com/git-jock/jock-cli/config"
	"github.com/git-jock/jock-cli/git"
	"os"
)

func main() {
	config.ReadConfig()

	invocation := getInvocationDetails()

	run(invocation)
}

// Takes in the invocation details and applies run logic to it.
func run(invocation *InvocationDetails) {
	if invocation.version {
		displayVersion()
	}

	if invocation.help || len(invocation.plugin) == 0 {
		displayHelp()
	}

	runPlugin(invocation)
}

/**********************************************
Invocation
 **********************************************/

// InvocationDetails holds the processed details of the jock invocation for logic to be applied later.
type InvocationDetails struct {
	version    bool
	help       bool
	folders    map[string]config.FolderConfig
	plugin     string
	pluginArgs []string
}

// Loops over the arguments supplied to jock and returns a pointer to the populated invocation details.
func getInvocationDetails() *InvocationDetails {
	args := os.Args[1:]

	invocation := &InvocationDetails{
		folders: make(map[string]config.FolderConfig),
	}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--help", "-h":
			invocation.help = true
			break
		case "--version", "-v":
			invocation.version = true
			break
		case "--folder", "-f":
			i++
			invocation.folders[args[i]] = config.GetFolderConfig(args[i])
			break
		default:
			setPluginDetails(invocation, args[i:])
			return invocation
		}
	}

	return invocation
}

// Takes in a pointer to the invocation details and populates the relevant fields from the plugin argument slice.
func setPluginDetails(c *InvocationDetails, pluginArgs []string) {
	c.plugin = pluginArgs[0]

	if len(pluginArgs) > 1 {
		c.pluginArgs = pluginArgs[1:]
	}
}

/**********************************************
JockPlugin
**********************************************/

// Will run the plugin.
func runPlugin(invocation *InvocationDetails) {
	fmt.Printf("Plugin:      %s\n", invocation.plugin)
	fmt.Printf("Plugin Args: %s\n", invocation.pluginArgs)
	fmt.Printf("Folders:     %s\n", invocation.folders)

	for k, v := range invocation.folders {
		fmt.Printf("%s config for %s: %s\n", invocation.plugin, k, v.Plugins[invocation.plugin])
	}

	switch invocation.plugin {
	case "git":
		git.Run(invocation.pluginArgs, invocation.folders)
	}
}

/**********************************************
Print
**********************************************/

// Will display the version of jock-cli and exit.
func displayVersion() {
	fmt.Println("0.0.0")
	os.Exit(0)
}

// Displays help information for jock, including example usage and a description of available options and flags.
func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("    jock [JOCK_ARGS] [PLUGIN] [PLUGIN_COMMAND] [PLUGIN_COMMAND_ARGS]")
	fmt.Println("    e.g. jock -f user-serv git clone --recurse-submodules")
	fmt.Println("")
	fmt.Println("Jock Arguments:")
	fmt.Println("    --version, -v	Print version and exit")
	fmt.Println("    --help, -h     Print this help text and exit")
	fmt.Println("    --folder, -f	Define one or more folders to run the plugin command on")
	os.Exit(0)
}
