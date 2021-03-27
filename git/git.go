package main

import (
	"fmt"
	"os"
	"os/exec"
)

type run string

var commandHandlers = map[string]interface{}{
	"clone":    gitClone,
	"add":      gitCommon,
	"restore":  gitCommon,
	"rm":       gitCommon,
	"branch":   gitCommon,
	"commit":   gitCommon,
	"reset":    gitCommon,
	"switch":   gitCommon,
	"tag":      gitCommon,
	"fetch":    gitCommon,
	"pull":     gitCommon,
	"push":     gitCommon,
	"checkout": gitCommon,
}

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

	handler := commandHandlers[args[0]]

	if handler == nil {
		fmt.Println("Unsupported command " + args[0])
		os.Exit(1)
	}

	handler.(func([]string))(args)
}

var JockPlugin run
