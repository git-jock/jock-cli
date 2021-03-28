package main

import (
	"fmt"
	"os"
	"os/exec"
)

type run string

func gitClone() {
}

func gitCommon(args []string) {
	cmd := exec.Command("git", args[0], "--help")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func (r run) Run(args []string) {
	if len(args) == 0 {
		fmt.Println("No args given")
		os.Exit(1)
	}

	switch args[0] {
	case "clone":
		gitClone()
		break
	case "add":
	case "restore":
	case "rm":
	case "branch":
	case "commit":
	case "reset":
	case "switch":
	case "tag":
	case "fetch":
	case "pull":
	case "push":
	case "checkout":
		gitCommon(args)
		break
	default:
		fmt.Println("Unsupported command " + args[0])
		os.Exit(1)
	}
}

var JockPlugin run
