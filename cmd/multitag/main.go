package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/mkungla/multidag"
)

var commands map[string]*flag.FlagSet

func main() {

	setupcmd()

	if len(os.Args) < 2 {
		help()
		os.Exit(0)
	}

	fset, ok := commands[os.Args[1]]
	if !ok {
		help()
		os.Exit(0)
	}
	fset.Parse(os.Args[2:])

	switch os.Args[1] {
	case "demo":
		demo(fset)
	default:
		help()
		os.Exit(1)
	}
}

func setupcmd() {
	commands = make(map[string]*flag.FlagSet)
	demoCmd := flag.NewFlagSet("demo", flag.ExitOnError)
	_ = demoCmd.Bool("config", false, "prints json output of demo config.")
	// _ = demoCmd.Bool("serve", false, "serve demo api.")
	commands[demoCmd.Name()] = demoCmd
}

func help() {
	fmt.Println(flag.CommandLine.Name())
	flag.CommandLine.PrintDefaults()

	for cmd, fset := range commands {
		fmt.Println(cmd)
		fset.PrintDefaults()
	}
}

func demo(fset *flag.FlagSet) {
	print := fset.Lookup("config").Value.String() == "true"
	// serve := fset.Lookup("serve").Value.String() == "true"

	// Create new authority by username / namespace
	authority, _ := multidag.New("mkungla")
	authority.Meta.Title = "mkungla's authority"
	authority.Meta.Description = "authority represents something like org|user|group|account etc."

	ws, _ := authority.NewWorkspace("demo")
	ws.Meta.Title = "Demo Workspace"
	ws.Meta.Description = "This is demo workspace."

	dag := ws.NewDAG("Demo DAG")
	dag.Meta.Description = "This is demo DAG."

	issues := dag.NewNode("issues")
	issues.Meta.Description = "Runs your DAG anytime the issues event occurs."
	issues.NewPort(multidag.PortOut)

	noop := dag.NewNode("noop")
	noop.Meta.Description = "Just a dummy node."

	schedule := dag.NewNode("schedule")
	schedule.Meta.Description = "You can schedule a workflow to run at specific UTC times using POSIX cron syntax."
	schedule.NewPort(multidag.PortOut)

	issueComment := dag.NewNode("issue comment")
	issueComment.Meta.Description = "Runs your workflow anytime the issue_comment event occurs."
	issueComment.NewPort(multidag.PortOut)

	manageLabels := dag.NewNode("manage labels")
	manageLabels.NewPort(multidag.PortIn)
	manageLabels.NewPort(multidag.PortOut)

	composeComment := dag.NewNode("compose comment")
	composeComment.NewPort(multidag.PortIn)
	composeComment.NewPort(multidag.PortOut)

	reaction := dag.NewNode("reaction")
	reaction.NewPort(multidag.PortIn)

	edited := dag.NewNode("edited")
	edited.NewPort(multidag.PortIn)

	closed := dag.NewNode("closed")
	closed.NewPort(multidag.PortIn)

	deleted := dag.NewNode("deleted")
	deleted.NewPort(multidag.PortIn)

	daily := dag.NewNode("daily")
	daily.NewPort(multidag.PortIn)
	daily.NewPort(multidag.PortOut)

	addLabels := dag.NewNode("add labels")
	addLabels.NewPort(multidag.PortIn)

	removeLabels := dag.NewNode("remove labels")
	removeLabels.NewPort(multidag.PortIn)

	createComment := dag.NewNode("create comment")
	createComment.NewPort(multidag.PortIn)

	stale := dag.NewNode("stale")
	stale.NewPort(multidag.PortIn)

	// if serve {
	// 	fmt.Println("SERVING API:")
	// }

	if print {
		fmt.Println("config:")
		data, _ := json.MarshalIndent(authority, "", "  ")
		fmt.Println(string(data))
	}
}
