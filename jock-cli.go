package main

import (
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"
)

func main() {
	getArgs()
}

var rVersion, _ = regexp.Compile("--version|-v")
var rHelp, _ = regexp.Compile("--help|-h")
var rFolder, _ = regexp.Compile("--folder|-f")

func getArgs() {
	args := os.Args[1:]

	var folders []string

	for i := 0; i < len(args); i++ {
		found := false
		if rVersion.MatchString(args[i]) {
			displayVersion()
		}

		if rHelp.MatchString(args[i]) {
			displayHelp()
		}

		if rFolder.MatchString(args[i]) {
			folders = append(folders, args[i+1])
			i++
			found = true
		}

		if !found {
			fmt.Println("not found, must be plugin: " + args[i])
			break
		}
	}

	fmt.Println("folders:")
	fmt.Printf("%v\n", folders)
}

func displayVersion() {
	fmt.Println("0.0.0")
	os.Exit(0)
}

func displayHelp() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintln(writer, "Usage:")
	fmt.Fprintln(writer, "	jock [JOCK_ARGS] [PLUGIN] [PLUGIN_COMMAND] [PLUGIN_COMMAND_ARGS]")
	fmt.Fprintln(writer, "	e.g. jock --group services git clone --recurse-submodules")
	fmt.Fprintln(writer, "")
	fmt.Fprintln(writer, "Jock Arguments:")
	fmt.Fprintln(writer, "	--version, -v	Print version and exit")
	fmt.Fprintln(writer, "	--help, -h		Print this help text and exit")
	fmt.Fprintln(writer, "	--folder, -f	Define a folder to run the plugin command on. May specify multiple.")

	writer.Flush()

	os.Exit(0)
}
