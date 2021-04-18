package git

import (
	"fmt"
	"github.com/git-jock/jock-cli/config"
	"os"
	"os/exec"
	"strings"
)

func getPath(location string) string {
	location = strings.ReplaceAll(location, "~", "$HOME")
	location = os.ExpandEnv(location)
	return location
}

const git = "git"

func gitClone(args []string, folders map[string]config.FolderConfig) {
	beginning := []string{"clone"}
	ending := args[1:]

	buildArgs := func(remote string, location string) []string {
		return append(append(beginning, remote, location), ending...)
	}

	for k, v := range folders {
		if gitConfig, ok := v.Plugins[git]; ok {
			fmt.Printf("Running [%s] on [%s]\n", args[0], k)
			cmd := exec.Command(git, buildArgs(gitConfig["remote"], getPath(v.Location))...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			_ = cmd.Run()
		}
	}
}

func gitCommon(args []string, folders map[string]config.FolderConfig) {
	beginning := []string{"-C"}

	buildArgs := func(location string) []string {
		return append(append(beginning, location), args...)
	}

	for k, v := range folders {
		fmt.Printf("Running [%s] on [%s]\n", args[0], k)
		cmd := exec.Command(git, buildArgs(getPath(v.Location))...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
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
