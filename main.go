package main

import (
	"fmt"
	"os"
	"plugin"
	"regexp"
	"text/tabwriter"
)

func main() {
	args := getArgs()

	run(args)
}

func run(c Command) {
	if c.help {
		displayHelp()
	}

	if c.version {
		displayVersion()
	}

	if len(c.plugin) > 0 {
		runPlugin(c)
	}
}

/**********************************************
ARGS
 **********************************************/

var rVersion, _ = regexp.Compile("--version|-v")
var rHelp, _ = regexp.Compile("--help|-h")
var rFolder, _ = regexp.Compile("--folder|-f")

type Command struct {
	version    bool
	help       bool
	folders    []string
	plugin     string
	pluginArgs []string
}

func getArgs() Command {
	args := os.Args[1:]

	c := Command{}

	for i := 0; i < len(args); i++ {
		found := false
		if rVersion.MatchString(args[i]) {
			c.version = true
			found = true
		}

		if rHelp.MatchString(args[i]) {
			c.help = true
			found = true
		}

		if rFolder.MatchString(args[i]) {
			c.folders = append(c.folders, args[i+1])
			i++
			found = true
		}

		if !found {
			c = setPluginDetails(c, args[i:])
			break
		}
	}

	return c
}

func setPluginDetails(c Command, pluginArgs []string) Command {
	c.plugin = pluginArgs[0]

	if len(pluginArgs) > 1 {
		c.pluginArgs = pluginArgs[1:]
	}

	return c
}

/**********************************************
JockPlugin
**********************************************/

type JockPlugin interface {
	Run(args []string)
}

func runPlugin(c Command) {
	wd, _ := os.Getwd()

	plug, err := plugin.Open(wd + "/" +c.plugin + "/" + c.plugin + ".so")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	symPlugin, err := plug.Lookup("JockPlugin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var jockPlugin JockPlugin
	jockPlugin, ok := symPlugin.(JockPlugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	jockPlugin.Run(c.pluginArgs)
}

/**********************************************
Print
**********************************************/

func displayVersion() {
	fmt.Println("0.0.0")
	os.Exit(0)
}

func displayHelp() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "Usage:")
	fmt.Fprintln(writer, "	jock [JOCK_ARGS] [PLUGIN] [PLUGIN_COMMAND] [PLUGIN_COMMAND_ARGS]")
	fmt.Fprintln(writer, "	e.g. jock -f user-serv git clone --recurse-submodules")
	fmt.Fprintln(writer, "")
	fmt.Fprintln(writer, "Jock Arguments:")
	fmt.Fprintln(writer, "	--version, -v	Print version and exit")
	fmt.Fprintln(writer, "	--help, -h		Print this help text and exit")
	fmt.Fprintln(writer, "	--folder, -f	Define a folder to run the plugin command on. May specify multiple.")

	writer.Flush()

	os.Exit(0)
}
