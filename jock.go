package main

import (
	"fmt"
	"github.com/git-jock/jock-cli/proto"
	"github.com/git-jock/jock-cli/shared"
	"github.com/hashicorp/go-plugin"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

func main() {
	readConfig()

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

type FolderConfig struct {
	Location string                 `mapstructure:"location"`
	Plugins  map[string]interface{} `mapstructure:"plugins"`
}

// Holds the processed details of the jock invocation for logic to be applied later.
type InvocationDetails struct {
	version    bool
	help       bool
	folders    map[string]FolderConfig
	plugin     string
	pluginArgs []string
}

// Loops over the arguments supplied to jock and returns a pointer to the populated invocation details.
func getInvocationDetails() *InvocationDetails {
	args := os.Args[1:]

	invocation := &InvocationDetails{
		folders: make(map[string]FolderConfig),
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
			invocation.folders[args[i]] = getFolderConfig(args[i])
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

	//log.SetOutput(ioutil.Discard)

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.HandShake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", "./example/example"),
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	raw, err := rpcClient.Dispense("grpcPlugin")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	jp := raw.(shared.JP)

	var req proto.RunRequest
	req.Args = invocation.pluginArgs

	result, err := jp.Run(&req)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Println(result)

	os.Exit(0)
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

/**********************************************
Config
**********************************************/

// Reads the config file from ~/.jockrc
func readConfig() {
	viper.SetConfigName(".jockrc")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println(err)
		} else {
			// Config file was found but another error was produced
			fmt.Println(err)
		}
	}
}

const FOLDERS = "folders"

// Checks that the config contains the specified folder, and returns marshalled FolderConfig if so.
func getFolderConfig(name string) FolderConfig {
	folder := FolderConfig{}

	if !viper.InConfig(FOLDERS) {
		fmt.Print("Config does not contain folders")
		os.Exit(1)
	}

	if !viper.Sub(FOLDERS).InConfig(name) {
		fmt.Printf("Config folders do not contain [%s]", name)
		os.Exit(1)
	}

	err := viper.Sub(FOLDERS).UnmarshalKey(name, &folder)
	if err != nil {
		fmt.Printf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}

	return folder
}
