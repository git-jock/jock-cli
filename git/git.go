package git

import (
	"fmt"
	"github.com/git-jock/jock-cli/config"
	"os"
)

type run string

func gitClone(args []string, folders map[string]config.FolderConfig) {

}

func gitCommon(args []string, folders map[string]config.FolderConfig) {
	for k, v := range folders {
		if gitConfig, ok := v.Plugins["git"]; ok {
			fmt.Printf("Running %s on folder %s in path %s", args[0], k, gitConfig)
			//cmd := exec.Command("git", args[0], "-C", gitConfig["remote"])
			//cmd.Stdout = os.Stdout
			//cmd.Stderr = os.Stderr
			//_ = cmd.Run()
		}
	}
}

func Run(args []string, folders map[string]config.FolderConfig) {
	if len(args) == 0 {
		fmt.Println("No args given")
		os.Exit(1)
	}

	if len(folders) == 0 {
		fmt.Println("No folders given")
		os.Exit(1)
	}

	switch args[0] {
	case "clone":
		gitClone(args, folders)
		break
	case "add", "branch", "checkout", "clean", "commit", "fetch", "gc", "grep", "init", "merge", "pull", "push", "rebase", "reset", "restore", "rm", "stash", "switch", "tag":
		gitCommon(args, folders)
		break
	default:
		fmt.Println("Unsupported command " + args[0])
		os.Exit(1)
	}
}
