package main

import (
	"fmt"
	"os"
	"reflect"
	"scm/lib/command"
	"strings"
)

var argsMap = map[string]func(options []string){
	"init":     command.Init,
	"i":        command.Init,
	"branch":   command.Branch,
	"b":        command.Branch,
	"checkout": command.Checkout,
	"co":       command.Checkout,
	"commit":   command.Commit,
	"c":        command.Commit,
	"pull":     command.Pull,
	"p":        command.Pull,
	"merge":    command.Merge,
	"m":        command.Merge,
	"fetch":    command.Fetch,
	"f":        command.Fetch,
	"add":      command.Add,
	"a":        command.Add,
}

func printHelp(msg string) {
	fmt.Println(msg)
	fmt.Print("commands: ")
	fmt.Print(reflect.ValueOf(argsMap).MapKeys())
}

func routeCommand(args []string) {
	if len(args) == 0 {
		printHelp("no command supplied.")
	} else {
		cmd := strings.TrimSpace(args[0])
		options := args[1:]
		if cmd == "" {
			printHelp("no command supplied.")
		} else {
			//if arg is in map execute else print help
			if val, pres := argsMap[cmd]; pres {
				val(options)
			} else {
				printHelp("command not found.")
			}
		}
	}

}

func main() {
	args := os.Args[1:]
	routeCommand(args)
}
